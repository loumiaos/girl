package main

import (
	"log"
	"robot/logic"

	"sync"
)

func main() {
	log.Println("robot run!")
	var wg sync.WaitGroup
	wg.Add(1)

	logic.Start()

	wg.Wait()
}
