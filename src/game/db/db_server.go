//数据库服务
package db

import (
	"github.com/snowyyj001/loumiao/log"

	"github.com/snowyyj001/loumiao/gorpc"
)

type DBServer struct {
	gorpc.GoRoutineLogic
}

func (self *DBServer) DoInit() {
	log.Info("DBServer DoInit")

}

func (self *DBServer) DoRegsiter() {
	self.Register("loginAccount", loginAccount)
	self.Register("getPlayer", getPlayer)
	self.Register("getGameCfg", getGameCfg)
}

func (self *DBServer) DoDestory() {
	log.Info("DBServer destory")
}
