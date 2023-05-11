package comsumer

import "github.com/tiyee/palmon/consts"

type IProcessor interface {
	Run(worker *worker, job Job)
	Name() string
}

var processors map[consts.Processor]IProcessor

func init() {
	processors = map[consts.Processor]IProcessor{
		consts.PULLER:     Puller{},
		consts.PROCESSOR1: Processor1{},
		consts.PUSHER:     Pusher{},
	}
}
