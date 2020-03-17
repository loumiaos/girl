package niuniu

import (
	"game/world/agent"
)

type Player struct {
	agent *agent.Agent
	seat  int
	state RoomState
}

func (self *Player) disConnect() {

}
