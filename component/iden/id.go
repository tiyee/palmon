package iden

import (
	"sync/atomic"
	"time"
)

const (
	workerIDBits = uint64(10) // 10bit 工作机器ID中的 5bit workerID

	sequenceBits = uint64(12)

	maxWorkerID = int64(-1) ^ (int64(-1) << workerIDBits) //节点ID的最大值 用于防止溢出
	maxSequence = int64(-1) ^ (int64(-1) << sequenceBits)

	timeLeft = uint8(22) // timeLeft = workerIDBits + sequenceBits // 时间戳向左偏移量
	workLeft = uint8(12) // workLeft = sequenceBits // 节点IDx向左偏移量
	// 2020-05-20 08:00:00 +0800 CST
	twepoch = int64(1589923200000) // 常量时间戳(毫秒)
)

var workId int64 = 0
var sequenceId int64 = 0

func init() {
	workId = 1 // show get from etcd or OS_ENV
}
func getMilliSeconds() int64 {
	return time.Now().UnixNano() / 1e6
}

func Gen() int64 {
	// 1+41+10+12
	workId %= maxWorkerID
	sn := atomic.LoadInt64(&sequenceId)
	sn %= maxSequence
	curTimeStamp := getMilliSeconds() - twepoch
	rightBinValue := curTimeStamp & 0x1FFFFFFFFFF
	rightBinValue <<= timeLeft
	atomic.AddInt64(&sequenceId, 1)
	return rightBinValue | workId<<workLeft | sn
}
