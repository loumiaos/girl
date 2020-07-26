package gate

import (
	"github.com/snowyyj001/loumiao"
	lgate "github.com/snowyyj001/loumiao/gate"
	"github.com/snowyyj001/loumiao/network"
)

func StartGate() {
	watchdog := &lgate.GateServer{ServerType: network.CLIENT_CONNECT}
	loumiao.Prepare(watchdog, "GateServer", false)

	watchdog.OnServerConnected = OnServerConnected
}

func OnServerConnected(uid int) {
}
