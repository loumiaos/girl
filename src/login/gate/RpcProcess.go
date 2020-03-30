package gate

import (
	"fmt"
	"msg"

	"github.com/snowyyj001/loumiao/gorpc"
)

func handlerGateInfo(igo gorpc.IGoRoutine, clientid int, data interface{}) interface{} {
	fmt.Printf("handlerGateInfo", data)
	req := data.(*msg.G_L_GateInfo)
	var str string = fmt.Sprintf("%s:%d", req.Ip, req.Port)
	GateInfos[str] = req.OnLineNumber
	return nil
}
