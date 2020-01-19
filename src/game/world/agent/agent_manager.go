package agent

import (
	"log"
	"sync"
)

type AgentMgr struct {
	agents     map[int]*Agent
	acc_id_Map map[string]int
}

var (
	inst *AgentMgr
	once sync.Once
)

func GetAgentMgr() *AgentMgr {
	once.Do(func() {
		inst = &AgentMgr{}
		inst.agents = make(map[int]*Agent)
		inst.acc_id_Map = make(map[string]int)

	})
	return inst
}

func (self *AgentMgr) AddAgent(agent *Agent) {
	if self.agents[agent.ID] != nil {
		log.Fatalf("AgentMgr.AddAgent error,[%d] has already added", agent.ID)
		return
	}
	self.agents[agent.ID] = agent
}

func (self *AgentMgr) GetAgent(playerId int) *Agent {
	return self.agents[playerId]
}
