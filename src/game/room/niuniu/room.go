package niuniu

import (
	"game/msg"
	"game/world/agent"

	"github.com/snowyyj001/loumiao"
	"github.com/snowyyj001/loumiao/log"
)

type FSMFunc func(dt int64)

type Room struct {
	id int

	curState   FSMState
	nextState  FSMState
	curFsmFunc [3]FSMFunc
	fsmTime    int64

	agents   map[int]*Player
	agentLen int

	players   map[int]*Player
	playerLen int

	visiters   map[int]*Player
	visiterLen int
}

func (self *Room) doStart(roomid int) {
	self.id = roomid
	self.agents = make(map[int]*Player)
	self.players = make(map[int]*Player)
	self.visiters = make(map[int]*Player)
	self.curState = FSM_Idle
	self.nextState = FSM_Idle
	self.changeState()
	self.curFsmFunc[0](0)
}

func (self *Room) SendClients(msg interface{}) {
	ids := make([]int, self.agentLen)
	i := 0
	for id, _ := range self.agents {
		ids[i] = id
		i++
	}
	loumiao.SendMulClient(ids, msg)
}

func (self *Room) getAllPlayers() []int {
	ids := make([]int, self.agentLen)
	i := 0
	for id, _ := range self.agents {
		ids[i] = id
		i++
	}
	return ids
}

func (self *Room) canJoinRoom(agent *agent.Agent) int {
	if self.playerLen+self.visiterLen >= MAX_ROOM_RENSHU {
		return Err_RoomFull
	}
	return 0
}

func (self *Room) joinRoom(agent *agent.Agent) {
	player := new(Player)
	player.agent = agent
	player.state = State_Idle
	player.seat = self.allocSeat()
	self.addPlayer(player)

	self.syncGame(player)
}

func (self *Room) disConnect(userId int) {
	player := self.getPlayer(userId)
	if player == nil {
		return
	}
	player.disConnect()

	if player.state == State_Idle {
		self.delPlayer(userId)
	} else {
		//req := &msg.R_C_PlayerOffline{}
		//self.SendClients(req)
	}
}

func (self *Room) allocSeat() int {
	for i := 0; i < MAX_SEAT; i++ {
		if self.players[i] == nil {
			return i
		}
	}
	return -1
}

func (self *Room) getPlayer(userId int) *Player {
	return self.agents[userId]
}

func (self *Room) addPlayer(player *Player) {
	if player.seat < 0 || player.seat >= MAX_ROOM_RENSHU {
		log.Warningf("niuniu room addPlayer error %d", player.seat)
		return
	}

	self.agents[player.agent.ID] = player
	self.visiters[player.seat] = player

	self.agentLen += 1
	self.visiterLen += 1

	req := &msg.R_C_JoinRoom{}
	req.RoomId = self.id
	req.Seat = player.seat
	req.State = int(player.state)

	loumiao.SendClient(player.agent.ClientId, req)
}

func (self *Room) delPlayer(id int) {
	if id < 0 || id >= MAX_SEAT {
		log.Warningf("niuniu room delPlayer error %d", id)
		return
	}
	delete(self.agents, id)
	self.agentLen -= 1
	if self.visiters[id] != nil {
		delete(self.visiters, id)
		self.visiterLen -= 1
	} else {
		delete(self.players, id)
		self.playerLen -= 1
	}
}

func (self *Room) syncGame(target *Player) {
	self.syncTable(target)
	self.syncPlayer(target)
}

func (self *Room) syncTable(target *Player) {
	req := &msg.NN_RC_TableInfo{}
	req.State = int(self.curState)
	req.LeftTime = int(self.fsmTime)
	req.BaseScore = 10

	loumiao.SendClient(target.agent.ClientId, req)
}

func (self *Room) syncPlayer(target *Player) {
	req := &msg.R_C_SyncPlayers{}
	req.Players = []msg.SyncPlayers{}

	for _, player := range self.agents {
		p := msg.SyncPlayers{}
		p.Gold = player.agent.Gold
		p.HeadIconUrl = player.agent.HeadIconUrl
		p.ID = player.agent.ID
		p.IpAddr = player.agent.IpAddr
		p.NickName = player.agent.NickName
		p.Seat = player.seat
		p.Sex = player.agent.Sex
		p.State = int(player.state)

		req.Players = append(req.Players, p)
	}

	loumiao.SendClient(target.agent.ClientId, req)
}

func (self *Room) sitDown(userId int) {
	agent := self.agents[userId]

	delete(self.visiters, agent.seat)
	self.visiterLen--
	for i := 0; i < MAX_SEAT; i++ {
		if self.players[i] == nil {
			agent.seat = i
			self.players[i] = agent
			self.playerLen++
		}

	}
}
