package consts

type TaskState int8

const (
	PENDING TaskState = iota
	PROCESSING
	FULFILL
	CANCEL
	ERROR
)
