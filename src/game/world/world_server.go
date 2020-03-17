// 游戏服务
package world

import (
	"game/world/agent"

	"github.com/snowyyj001/loumiao"
	"github.com/snowyyj001/loumiao/gorpc"
	"github.com/snowyyj001/loumiao/log"
)

var (
	This *WorldServer
)

type WorldServer struct {
	gorpc.GoRoutineLogic
}

//在这里注册网络消息
func (self *WorldServer) DoInit() {
	log.Info("WorldServer DoInit")
	This = self

	loumiao.RegisterNetHandler(self, "DISCONNECT", handlerOnDisConnect)
	loumiao.RegisterNetHandler(self, "C_S_Login", handlerLogin)
	loumiao.RegisterNetHandler(self, "C_S_HeartBeat", handlerHeartBeat)
	loumiao.RegisterNetHandler(self, "C_S_Login", handlerLogin)
	loumiao.RegisterNetHandler(self, "C_S_JoinRoom", handlerJoinRoom)
}

//在这里注册rpc消息
func (self *WorldServer) DoRegsiter() {
	//self.Register("handlerLogin", handlerLogin)
}

func (self *WorldServer) DoStart() {
	log.Infof("%s DoStart", self.Name)

	self.RunTicker(1000, self.Update)
}

func (self *WorldServer) DoDestory() {
	log.Info("WorldServer destory")
}

func (self *WorldServer) Update(igo gorpc.IGoRoutine, data interface{}) interface{} {
	dt := data.(int64)

	agent.GetAgentMgr().Update(dt)

	return nil
}
