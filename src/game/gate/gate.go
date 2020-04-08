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
}
