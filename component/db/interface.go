package db

import "github.com/tiyee/palmon/external/storage"

type IDb interface {
	Peek(id int64) (storage.Task, int)
	Save(row storage.Task)
	Update(row storage.Task) error
	Set(idx int, row storage.Task)
}
