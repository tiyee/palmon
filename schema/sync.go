package schema

import (
	"encoding/json"
)

type SyncJob struct {
	Error   int             `json:"error"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

func (j *SyncJob) Valid() error {

	return nil
}
func (j *SyncJob) Hook() {

}
