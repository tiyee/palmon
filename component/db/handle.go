package db

import "github.com/tiyee/palmon/external/storage"

var DB IDb

func Init() {
	DB = &storage.Tasks{}

}
