package agent

import (
	"game/config"

	"github.com/snowyyj001/loumiao/gorpc"
	"github.com/snowyyj001/loumiao/log"
	"github.com/snowyyj001/loumiao/util"
)

type AgentMgr struct {
	agents     map[int]*Agent
	acc_id_Map map[string]int
	sid_id_Map map[int]int
	offAgents  map[int]*Agent
}

var (
	This gorpc.IGoRoutine //go的包目录设置和循环引用问题
	inst *AgentMgr
)

func GetAgentMgr() *AgentMgr {
	if inst == nil {
		inst = &AgentMgr{}
		inst.agents = make(map[int]*Agent)
		inst.acc_id_Map = make(map[string]int)
		inst.sid_id_Map = make(map[int]int)
		inst.offAgents = make(map[int]*Agent)
		This = gorpc.GetGoRoutineMgr().GetRoutine("WorldServer")
	}

	return inst
}

func (self *AgentMgr) AddAgent(agent *Agent) {
	if self.agents[agent.ID] != nil {
		log.Errorf("AgentMgr.AddAgent error,[%d] has already added", agent.ID)
		return
	}
	if agent.ID <= 0 || agent.ClientId <= 0 || agent.Account == "" {
		log.Errorf("AgentMgr.AddAgent error,[%d][%s][%s] not leagl", agent.ID, agent.ClientId, agent.Account)
		return
	}
	self.agents[agent.ID] = agent
	self.acc_id_Map[agent.Account] = agent.ID
	self.sid_id_Map[agent.ClientId] = agent.ID
}

func (self AgentMgr) RemoveAgent(userid int) {
	agent := self.GetAgent(userid)
	if agent == nil {
		return
	}
	delete(self.acc_id_Map, agent.Account)
	delete(self.sid_id_Map, agent.ClientId)
	delete(self.agents, userid)

	//放到离线列表里
	self.offAgents[userid] = agent
}

func (self *AgentMgr) GetAgent(userid int) *Agent {
	return self.agents[userid]
}

func (self *AgentMgr) GetAgentByAccount(accName string) *Agent {
	userid, has := self.acc_id_Map[accName]
	if has {
		return self.agents[userid]
	}
	return nil
}

func (self *AgentMgr) GetAgentByServerId(clientId int) *Agent {
	userid, has := self.sid_id_Map[clientId]
	if has {
		return self.agents[userid]
	}
	return nil
}

func (self *AgentMgr) Update(dt int64) {
	curts := util.TimeStamp()
	for id, agent := range self.offAgents {
		if agent.LogoutTime+config.OFFLINE_TIME > curts {
			delete(self.offAgents, id)
		}
	}
}
