package gate

import (
	"msg"

	"github.com/snowyyj001/loumiao"
	lconfig "github.com/snowyyj001/loumiao/config"
	lgate "github.com/snowyyj001/loumiao/gate"
	"github.com/snowyyj001/loumiao/network"
)

func StartGate() {
	watchdog := &lgate.GateServer{ServerType: network.CLIENT_CONNECT}
	loumiao.Prepare(watchdog, "GateServer", false)

	watchdog.OnServerConnected = OnServerConnected
}

func OnServerConnected(uid int) {
	if uid != 1 {
		ReportGateInfo(uid, 0)
	}
}

func ReportGateInfo(uid int, num int) {
	req := &msg.G_L_GateInfo{Ip: lconfig.NET_GATE_IP, Port: lconfig.NET_GATE_PORT, OnLineNumber: num}
	loumiao.SendRpc(uid, req)
}
