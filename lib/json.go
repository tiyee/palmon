package lib

import (
	"encoding/json"
	"github.com/tiyee/palmon/schema"
	"io"
)

func JSONArgs(r io.Reader, v schema.ISchema) error {
	if err := json.NewDecoder(r).Decode(v); err != nil {
		return err
	}
	return v.Valid()

}
