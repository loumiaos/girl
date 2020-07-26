package main

import (
	"gate/game"
	"gate/gate"

	"github.com/snowyyj001/loumiao"
	"github.com/snowyyj001/loumiao/log"
)

func init() {

}

func main() {
	log.Info("gate server run!")

	gate.StartGate()

	//service start
	loumiao.Prepare(new(game.GameServer), "GameServer", false)

	loumiao.Run()
}
