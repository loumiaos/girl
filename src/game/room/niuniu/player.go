package niuniu

import (
	"game/world/agent"
)

type Player struct {
	agent *agent.Agent
	seat  int
	state RoomState

	handCards []int
	sortCards []int
}

func (self *Player) disConnect() {

}

func (self *Player) setRoomState(state RoomState) {
	self.state = state
}
