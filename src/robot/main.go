package main

import (
	"log"
	"robot/config"
	"robot/logic"
	"sync"

	"github.com/snowyyj001/loumiao"
	lconf "github.com/snowyyj001/loumiao/config"
)

func init() {
	lconf.NET_PROTOCOL = config.NET_PROTOCOL
	lconf.NET_WEBSOCKET = config.NET_WEBSOCKET
}
func main() {
	log.Println("robot run!")

	var wg sync.WaitGroup
	wg.Add(1)

	logic.Start()

	loumiao.Run()
	wg.Wait()
}
