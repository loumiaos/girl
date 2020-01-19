package world

import (
	"game/dbmodel"
	"game/msg"
	"game/world/agent"

	"github.com/snowyyj001/loumiao"
	"github.com/snowyyj001/loumiao/gorpc"
)

func handlerLogin(igo gorpc.IGoRoutine, clientid int, data interface{}) interface{} {
	loginData := data.(*msg.C_S_Login)

	user := agent.GetAgentMgr().GetAgent(int(loginData.UserID))
	if user == nil {
		userData := igo.Call("DB_Player", "getPlayer", loginData.UserID).(*dbmodel.User)
		user = &agent.Agent{
			User: *userData,
		}
		agent.GetAgentMgr().AddAgent(user)
	}
	user.ClientId = clientid

	user.OnLogin()

	resp := &msg.S_C_Login{}
	resp.UserID = int32(user.ID)
	resp.Gold = user.Gold
	resp.Coin = user.Coin
	resp.Money = user.Money
	resp.HeadIconUrl = user.HeadIconUrl
	resp.UnderWrite = user.UnderWrite
	resp.Sex = int32(user.Sex)
	resp.ActiveFlag = user.ActiveFlag
	resp.NickName = user.NickName

	loumiao.SendClient(igo, clientid, resp)

	return nil
}
