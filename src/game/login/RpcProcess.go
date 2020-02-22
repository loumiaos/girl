package login

import (
	"game/define"
	"game/msg"

	"github.com/snowyyj001/loumiao/gorpc"

	"github.com/snowyyj001/loumiao"
)

func handlerLogin(igo gorpc.IGoRoutine, clientid int, data interface{}) interface{} {
	req := data.(*msg.C_A_Login)

	userId := igo.Call("DBServer", "loginAccount", req).(int)
	var resp = &msg.A_C_Login{}
	if userId == 0 {
		resp.ErrCode = define.Err_Login_Pass
	} else {
		resp.UserID = userId
	}
	loumiao.SendClient(clientid, resp)
	return nil
}
