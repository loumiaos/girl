package login

import (
	"game/msg"

	"github.com/snowyyj001/loumiao/gorpc"

	"github.com/snowyyj001/loumiao"
)

func handlerLogin(igo gorpc.IGoRoutine, clientid int, data interface{}) interface{} {
	loginData := data.(*msg.C_A_Login)

	userId := igo.Call("DB_Player", "loginAccount", loginData).(int)
	var resp = &msg.A_C_Login{}
	if userId == 0 {
		resp.ErrorStr = "密码错误"
	} else {
		resp.UserID = int32(userId)
	}
	loumiao.SendClient(igo, clientid, resp)
	return nil
}
