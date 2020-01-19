// 游戏服务
package world

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

	loumiao.RegisterNetHandler(self, "C_S_Login", handlerLogin)
}

//在这里注册rpc消息
func (self *GameServer) DoRegsiter() {
	//self.Register("handlerLogin", handlerLogin)
}

func (self *GameServer) DoDestory() {
	log.Info("GameServer destory")
}

func (self *GameServer) Update() {

}
