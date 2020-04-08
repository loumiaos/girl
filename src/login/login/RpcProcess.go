package login

import (
	"login/define"
	"login/gate"
	"login/msg"

	"github.com/snowyyj001/loumiao"
	"github.com/snowyyj001/loumiao/gorpc"
)

func handlerLogin(igo gorpc.IGoRoutine, clientid int, data interface{}) interface{} {
	req := data.(*msg.C_A_Login)

	userId := igo.Call("DBServer", "loginAccount", req).(int)
	var resp = &msg.A_C_Login{}
	if userId == 0 {
		resp.ErrCode = define.Err_Login_Pass
	} else {
		resp.TokenID = This.AssignClientId()
		resp.UserID = userId
		resp.GateAddr = gate.GetPerfactGate()
	}
	loumiao.SendClient(clientid, resp)
	return nil
}
