# palmon

![巴鲁兽](https://img.tiyee.cn/u/dil2t3.jpg)

**palmon**是一个纯go实现的任务调度系统，没有任何第三方依赖，但只是**demo**级别,勿用于生产环境!!!


## Demo

目前本人在二级目录下部署了该项目，仅做示例

### 创建任务

> curl -XPOST https://uploader.tiyee.com.cn/palmon/task -d '{"processor":2,"payload":"12344","Prior":1}'

结果

```json
{
	"error": 0,
	"message": "ok",
	"data":
	{
		"task_id": 393778421588561921
	}
}
```

### 获取任务状态和详情

> curl 'https://uploader.tiyee.com.cn/palmon/task?task_id=393778421588561921'

```json
{
	"error": 0,
	"message": "ok",
	"data":{
		"task_id": 393778421588561921,
		"state": 2,
		"result": "1234412344"
	}
}
```


## API

统一返回格式

```json
{
    "error":0,
    "message":"ok",
    "data": {
        "task_id":393708785169862658,
        "state":2,
        "payload":"1234412344"
    }
}
```

| 名称       | 类型     | 说明          |
| ---------- | -------- | ------------- |
| error     | `integer` | 错误编号,0表示无错误，大于0表示有误 |
| message | `string` | 返回的消息备注      |
| data | `any` | 返回的数据 |


### `POST` /task 

添加任务，数据格式如下

```json
{
    "processor":2,
    "payload":"12344",
    "prior":1
}
```

#### 参数说明

| 名称       | 类型     | 说明          |
| ---------- | -------- | ------------- |
| processor     | `enum<integer>` | 任务类型编号(也可以称之为执行者编号)，用来区分不同的任务，以便用不同的执行器去执行.目前只做了2这一个执行器。 |
| payload | `any` | 任务携带的信息      |
| prior | `enum<integer>` | 优先级，1-高,2-中,3-低 |


### `GET` /task

获取任务详情

#### 参数

| 名称       | 类型     | 说明          |
| ---------- | -------- | ------------- |
| task_id     | `integer` | 任务id,创建时返回 |

#### 返回数据

```json
{
    "error":0,
    "message":"ok",
    "data":{
        "task_id":393708785169862658,
        "state":2,
        "result":"1234412344"
    }
}
```


#### 参数说明

| 名称       | 类型     | 说明          |
| ---------- | -------- | ------------- |
| task_id     | `integer` | 任务id |
| state | `enum<integer>` | 任务执行状态,0-待执行,1-执行中,2-已完成,3-已取消,4-发送错误      |
| result | `any` | 执行结果,json化数据 |

## 部署

> 注意，如果是mac系统，编译的时候需要修改`export GOOS=linux`为`export GOOS=darwin`

### docker部署

最简单的是`docker-compose`，直接执行下面命令

```shell
sudo mkdir -p /data/logs/palmon
git clone github.com/tiyee/palmon your_name
cd your_name
sudo docker-compose -f docker-compose.yml build
sudo docker-compose -f docker-compose.yml up -d

```
### kubernetes部署

本项目涉及的`coordinator`和`worker`是互相独立的，需要分别打包，可以直接利用github或其他自带的ci/cd工具打包后将镜像上传到镜像仓库，
但是目前不是高可用架构，`coordinator`只能部署一个pod，`worker`不限制pod数量。默认的`Puller`和`Pusher`里的`coordinator`的地址需要更改成k8s点的内部svc地址

## 高可用配置

![](https://img.tiyee.cn/u/le52t3.jpg!w400)

如图所示，存储和队列系统目前是一个纯go的简单实现，跟`coordinator`耦合在一起，`worker`会不断从`coordinator`拉取任务并执行，然后将任务的执行结果推送到`coordinator`,
如果`coordinator`部署多份，就会任务分布在不同的pod，数据不一致。

如果要高可用，存储和队列需要独立部署，它们本身需要是高可用的。然后`coordinator`就可以多点部署，通过vip系统或k8s的svc来统一访问`coordinator`。

## 优先队列

每个任务都有优先级`prior`属性,它是是一个数字，越小表示优先级越高。如果只是单纯按优先级来执行，有可能会导致低优先级的任务一直得不到执行，俗称`饿死`。所以还得考虑时间因素。
于是，根据一定算法，计算了一个`score`值，公式如下

`score= ( 1 << 62 ) - ( prior << 41 ) - factor * epochTimestamp`

我们可以通过设置`factor`大小来控制时间因素的影响，这样可以保证低优先级的任务也能被执行。

目前队列是一个纯go实现的`最大堆`，任务是按score来决定执行顺序的。但是如果是正式环境，显然不行。得用成熟的支持优先队列的队列来处理，比如`ActiveMQ`,
`RabbitMQ`等。

但新的问题来了，目前市面成熟的优先队列支持的优先级区间都很小，最多0-255，显然不支持我们score这么大的数字。因此需要设计一个新的方案。

所以，可以根据任务优先级投递到不同优先级的队列，然后由一个晒鞋者分别循环读取这些队列，再决定是否投入到活动队列。这种方案用任何队列都可实现，该过程如下图所示。

![优先队列。](https://img.tiyee.cn/u/2b9at3.jpg)

## 执行器
| 名称       | ID     | 说明          |
| ---------- | -------- | ------------- |
| puller    | - | 自带，执行从`coordinator`拉取任务的工作 |
| pusher| - | 自带，任务执行完后，由pusher将结果推送给`coordinator`     |
| proc1 | 2| 自定义执行器，当做一个简单demo，只是将payload repeat 一次然后返回 |

执行器是用来执行任务点，`worker`从`coordinator`拉取任务后，根据任务的类型，由不同的执行器去执行。

但有两个个例外，分别是puller和pusher，他们的任务不是从coordinator获取的，而是节点内部生成的。

`dispatcher`会按照指定的策略定期往节点`jobChannel`写入拉取任务，然后puller获取这个任务后就去coorddinator拉取任务，随后将任务投入`jobChannel`。

同理，pusher的任务也不是`coordinator`获取的， 其他执行器执行完任务后，会将执行结果写入pusher任务的payload里，然后由pusher将payload的内容包装成json数据推送给`coordinator`,
