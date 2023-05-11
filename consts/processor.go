package consts

type Processor int

const (
	PULLER Processor = iota
	PUSHER
	PROCESSOR1
)

var Processors = []Processor{
	PULLER,
	PUSHER,
	PROCESSOR1,
}
