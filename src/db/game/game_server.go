// 游戏服务
package game

import (
	"github.com/snowyyj001/loumiao"
	"github.com/snowyyj001/loumiao/gorpc"
	"github.com/snowyyj001/loumiao/log"
)

var (
	This *GameServer
)

type GameServer struct {
	gorpc.GoRoutineLogic
}

//在这里注册网络消息
func (self *GameServer) DoInit() {
	log.Info("GameServer DoInit")
	This = self
	loumiao.RegisterNetHandler(self, "C_S_HeartBeat", handlerHeartBeat)
}

//在这里注册rpc消息
func (self *GameServer) DoRegsiter() {
}

func (self *GameServer) DoStart() {
	log.Infof("%s DoStart", self.Name)

}

func (self *GameServer) DoDestory() {
	log.Info("GameServer destory")
}
