package world

import (
	"game/dbmodel"
	"game/define"
	"game/msg"
	"game/world/agent"

	"github.com/snowyyj001/loumiao/log"

	"github.com/snowyyj001/loumiao"
	"github.com/snowyyj001/loumiao/gorpc"
)

func handlerDisConnect(igo gorpc.IGoRoutine, clientid int, data interface{}) interface{} {
	user := agent.GetAgentMgr().GetAgentByServerId(clientid)
	if user != nil {
		user.OnLoginOut()
		agent.GetAgentMgr().RemoveAgent(user.ID)
	}
	return nil
}

func handlerHeartBeat(igo gorpc.IGoRoutine, clientid int, data interface{}) interface{} {
	resp := &msg.S_C_HeartBeat{}
	loumiao.SendClient(clientid, resp)

	return nil
}

func handlerLogin(igo gorpc.IGoRoutine, clientid int, data interface{}) interface{} {
	req := data.(*msg.C_S_Login)

	user := agent.GetAgentMgr().GetAgent(int(req.UserID))
	if user == nil {
		userData := igo.Call("DBServer", "getPlayer", req.UserID).(dbmodel.User)
		user = &agent.Agent{
			User:     userData,
			ClientId: clientid,
		}
		agent.GetAgentMgr().AddAgent(user)
	}
	user.ClientId = clientid

	user.OnLogin()

	resp := &msg.S_C_Login{}
	resp.UserID = user.ID
	resp.Gold = user.Gold
	resp.Coin = user.Coin
	resp.Money = user.Money
	resp.HeadIconUrl = user.HeadIconUrl
	resp.UnderWrite = user.UnderWrite
	resp.Sex = user.Sex
	resp.ActiveFlag = user.ActiveFlag
	resp.NickName = user.NickName

	loumiao.SendClient(clientid, resp)

	return nil
}

func handlerJoinRoom(igo gorpc.IGoRoutine, clientid int, data interface{}) interface{} {
	player := agent.GetAgentMgr().GetAgentByServerId(clientid)
	if player == nil {
		log.Warningf("handlerJoinRoom: player[%d] is not exist", clientid)
		return nil
	}
	req := data.(*msg.C_S_JoinRoom)
	server := gorpc.GetGoRoutineMgr().GetRoutine(req.Service)

	err := 0
	for {
		if server == nil {
			err = define.Err_Room_NoExist
			break
		}
		break //这个break容易忘记写，会导致死循环，但goto使用不方便，跪求官方给出do while语法糖
	}
	if err != 0 {
		resp := &msg.S_C_JoinRoom{}
		resp.RoomId = req.RoomId
		resp.ErrCode = err
		loumiao.SendClient(clientid, resp)
		return nil
	}

	m := gorpc.M{Data: *player, Id: req.RoomId}
	err = igo.Call(req.Service, "joinRoom", m).(int)
	if err == 0 {
		player.GameArea = req.Service
	}

	return nil
}
