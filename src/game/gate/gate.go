package gate

import (
	"github.com/snowyyj001/loumiao"
	lgate "github.com/snowyyj001/loumiao/gate"
	"github.com/snowyyj001/loumiao/network"
)

func StartGate() {
	watchdog := &lgate.GateServer{ServerType: network.SERVER_CONNECT}
	loumiao.Prepare(watchdog, "GateServer", false)

	watchdog.OnClientConnected = OnClientConnected
}

func OnClientConnected(uid int) {
	loumiao.RegisterRpcHandler(uid, "C_S_Login")
	loumiao.RegisterRpcHandler(uid, "S_C_Login")
	loumiao.RegisterRpcHandler(uid, "C_S_JoinRoom")
	loumiao.RegisterRpcHandler(uid, "S_C_JoinRoom")
	loumiao.RegisterRpcHandler(uid, "R_C_JoinRoom")
	loumiao.RegisterRpcHandler(uid, "R_C_SyncPlayers")

}
