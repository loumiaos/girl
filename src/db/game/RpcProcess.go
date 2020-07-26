package game

import (
	"gate/msg"

	"github.com/snowyyj001/loumiao"
	"github.com/snowyyj001/loumiao/gorpc"
)

func handlerHeartBeat(igo gorpc.IGoRoutine, clientid int64, data interface{}) interface{} {
	resp := &msg.S_C_HeartBeat{}
	loumiao.SendClient(clientid, resp)

	return nil
}
