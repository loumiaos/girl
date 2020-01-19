// db_player
package dblogic

import (
	"github.com/snowyyj001/loumiao/log"

	"github.com/snowyyj001/loumiao/gorpc"
)

type DB_Player struct {
	gorpc.GoRoutineLogic
}

func (self *DB_Player) DoInit() {
	log.Info("DB_Player DoInit")

}

func (self *DB_Player) DoRegsiter() {
	self.Register("loginAccount", loginAccount)
	self.Register("getPlayer", getPlayer)
}

func (self *DB_Player) DoDestory() {
	log.Info("DB_Player destory")
}
