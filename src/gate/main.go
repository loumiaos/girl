package main

import (
	"gate/gate"

	"github.com/snowyyj001/loumiao"
	"github.com/snowyyj001/loumiao/log"
)

func init() {

}

func main() {
	log.Info("gate server run!")

	gate.StartGate()

	loumiao.Run()
}
