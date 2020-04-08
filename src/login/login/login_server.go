// 登陆服务
package login

import (
	"sync/atomic"

	"github.com/snowyyj001/loumiao"
	"github.com/snowyyj001/loumiao/gorpc"
	"github.com/snowyyj001/loumiao/log"
)

var (
	This *LoginServer
)

type LoginServer struct {
	gorpc.GoRoutineLogic
	m_nIdSeed int32
}

func (self *LoginServer) DoInit() {
	log.Info("LoginServer DoInit")
	This = self

	loumiao.RegisterNetHandler(self, "C_A_Login", handlerLogin)
}

func (self *LoginServer) DoRegsiter() {

}

func (self *LoginServer) DoDestory() {
	log.Info("LoginServer destory")
}

func (self *LoginServer) AssignClientId() int {
	return int(atomic.AddInt32(&self.m_nIdSeed, 1))
}
