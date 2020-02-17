package niuniu

import (
	"game/world/agent"
)

type FSMState int
type FSMFunc func(dt int64)

const (
	FSM_Idle FSMState = iota
	FSM_Fapai
	FSM_Bipai
	FSM_Result
)

type Room struct {
	id int

	curState   FSMState
	nextState  FSMState
	curFsmFunc [3]FSMFunc
}

func (self *Room) doStart(roomid int) {
	self.id = roomid
	self.curState = FSM_Idle
	self.nextState = FSM_Idle
	self.changeState()
	self.curFsmFunc[0](0)
}

func (self *Room) SetState(state FSMState) {
	self.nextState = state
}

func (self *Room) changeState() {
	switch self.nextState {
	case FSM_Idle:
		self.curFsmFunc[0] = self.onEnterIdle
		self.curFsmFunc[1] = self.onExecIdle
		self.curFsmFunc[2] = self.onExitIdle
	case FSM_Fapai:
		self.curFsmFunc[0] = self.onEnterFaPai
		self.curFsmFunc[1] = self.onExecFaPai
		self.curFsmFunc[2] = self.onExitFaPai
	case FSM_Bipai:
		self.curFsmFunc[0] = self.onEnterBiPai
		self.curFsmFunc[1] = self.onExecBiPai
		self.curFsmFunc[2] = self.onExitBiPai
	case FSM_Result:
		self.curFsmFunc[0] = self.onEnterResult
		self.curFsmFunc[1] = self.onExecResult
		self.curFsmFunc[2] = self.onExitResult
	default:

	}
}

func (self *Room) canJoinRoom(player *agent.Agent) int {
	return 0
}

func (self *Room) joinRoom(player *agent.Agent) {

}
