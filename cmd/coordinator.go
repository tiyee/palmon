package main

import (
	server "github.com/tiyee/palmon"
	"github.com/tiyee/palmon/component/db"
	"github.com/tiyee/palmon/component/log"
	"github.com/tiyee/palmon/component/queue"
)

func main() {
	log.SetupLogger("./logs/coordinator.log")
	db.Init()
	queue.Init()
	server.Run()
}
