package msg

import (
	"github.com/snowyyj001/loumiao/message"
)

func init() {

	message.RegisterPacket(&C_S_HeartBeat{})
	message.RegisterPacket(&S_C_HeartBeat{})
}
