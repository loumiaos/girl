package gate

import (
	"fmt"
	"msg"

	"github.com/snowyyj001/loumiao/gorpc"
)

func handlerHeartBeat(igo gorpc.IGoRoutine, clientid int, data interface{}) interface{} {
	resp := &msg.S_C_HeartBeat{}
	loumiao.SendClient(clientid, resp)

	return nil
}
