package gate

import (
	"gate/config"
	"msg"

	"github.com/snowyyj001/loumiao"
	lconfig "github.com/snowyyj001/loumiao/config"
	lgate "github.com/snowyyj001/loumiao/gate"
	"github.com/snowyyj001/loumiao/network"
)

func RegisterRpc(watchdog *lgate.GateServer) {
	watchdog.RegisterRpc("C_S_Login", config.WORLD_NODE)
	watchdog.RegisterRpc("S_C_Login", config.WORLD_NODE)
	watchdog.RegisterRpc("C_S_JoinRoom", config.WORLD_NODE)
	watchdog.RegisterRpc("S_C_JoinRoom", config.WORLD_NODE)
	watchdog.RegisterRpc("R_C_JoinRoom", config.WORLD_NODE)
	watchdog.RegisterRpc("R_C_SyncPlayers", config.WORLD_NODE)
}

func StartGate() {
	watchdog := &lgate.GateServer{ServerType: network.CLIENT_CONNECT}
	RegisterRpc(watchdog)
	loumiao.Prepare(watchdog, "GateServer", false)

	watchdog.OnServerConnected = OnServerConnected
}

func OnServerConnected(uid int) {
	if uid == config.LOGIN_NODE {
		ReportGateInfo(uid, 0)
	}
}

func ReportGateInfo(uid int, num int) {
	req := &msg.G_L_GateInfo{Ip: lconfig.NET_GATE_IP, Port: lconfig.NET_GATE_PORT, OnLineNumber: num}
	loumiao.SendRpc(uid, req)
}
