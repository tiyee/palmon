package main

import (
	"github.com/tiyee/palmon/component/log"
	"github.com/tiyee/palmon/controller/comsumer"
)

func main() {
	log.SetupLogger("./logs/worker.log")
	dispatcher := comsumer.NewDispatcher(10, 100)
	dispatcher.Run()

}
