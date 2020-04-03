package msg

import (
	"github.com/snowyyj001/loumiao/message"
)

func init() {

	message.RegisterPacket(&C_S_HeartBeat{})
	message.RegisterPacket(&S_C_HeartBeat{})
	message.RegisterPacket(&C_S_Login{})
	message.RegisterPacket(&S_C_Login{})
	message.RegisterPacket(&C_S_JoinRoom{})
	message.RegisterPacket(&S_C_JoinRoom{})

	message.RegisterPacket(&R_C_JoinRoom{})
	message.RegisterPacket(&R_C_SyncPlayers{})
}
