// 登陆服务
package login

import (
	"github.com/snowyyj001/loumiao"
	"github.com/snowyyj001/loumiao/gorpc"
	"github.com/snowyyj001/loumiao/log"
)

var (
	This *LoginServer
)

type LoginServer struct {
	gorpc.GoRoutineLogic
}

func (self *LoginServer) DoInit() {
	log.Info("LoginServer DoInit")
	loumiao.RegisterNetHandler(self, "C_A_Login", handlerLogin)
}

func (self *LoginServer) DoRegsiter() {

}

func (self *LoginServer) DoDestory() {
	log.Info("LoginServer destory")
}
