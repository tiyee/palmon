package schema

import (
	"errors"
	"github.com/tiyee/palmon/consts"
)

type Job struct {
	Processor int  `json:"processor"`
	Payload   any  `json:"payload"`
	Prior     int8 `json:"prior"`
}

func (j *Job) Valid() error {
	if j.Processor == 0 || j.Processor >= len(consts.Processors) {
		return errors.New("invalid processor")
	}
	if j.Prior < 1 || j.Prior > 3 {
		return errors.New("invalid prior")
	}
	return nil
}
func (j *Job) Hook() {

}
