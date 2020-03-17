package agent

import (
	"game/dbmodel"

	"github.com/snowyyj001/loumiao/util"
)

type Agent struct {
	dbmodel.User
	ClientId int
	GameArea string
}

func (self *Agent) OnLogin() {
	self.LoginTime = util.TimeStamp()
}

func (self *Agent) OnLoginOut() {
	self.LogoutTime = util.TimeStamp()

	if len(self.GameArea) > 0 {
		This.Send(self.GameArea, "disconnect", self.ID)
	}
}

func (self *Agent) SendClient() {

}
