package niuniu

import (
	"fmt"
	"game/define"
	"game/msg"
	"game/world/agent"

	//"unsafe"
	"github.com/snowyyj001/loumiao/gorpc"
)

func (self *GameServer) handlerLogOut(igo gorpc.IGoRoutine, data interface{}) interface{} {
	userId := data.(int)
	room := self.GetRoomByUserId(userId)
	room.disConnect(userId)

	return nil
}

func (self *GameServer) handlerJoinRoom(igo gorpc.IGoRoutine, data interface{}) interface{} {
	m := data.(gorpc.M)

	player := m.Data.(agent.Agent)
	roomid := m.Id

	err := 0
	if self.Rooms[roomid] != nil {
		err = self.Rooms[roomid].canJoinRoom(&player)
	} else {
		err = define.Err_Room_NoExist
	}

	if err == 0 {
		self.Rooms[roomid].joinRoom(&player)
		//log.Debugf("玩家%d加入房间%d", player.ID, roomid)
		self.Players[player.ID] = roomid
		self.Agents[player.ClientId] = player.ID
	}
	return err
}

func handlerSitDown(igo gorpc.IGoRoutine, clientid int, data interface{}) interface{} {
	userId, room := This.GetRoomByClientId(clientid)
	if room != nil {
		room.sitDown(userId)
	}
	return nil
}

func handlerPlayerReady(igo gorpc.IGoRoutine, clientid int, data interface{}) interface{} {
	fmt.Println("handlerPlayerReady", data)
	userId, room := This.GetRoomByClientId(clientid)
	if room == nil {
		return nil
	}
	room.ready(userId)
	return nil
}

func handlerQZhuang(igo gorpc.IGoRoutine, clientid int, data interface{}) interface{} {
	userId, room := This.GetRoomByClientId(clientid)
	if room == nil {
		return nil
	}
	m := data.(*msg.NN_RC_QZhuang)
	room.qiangZhuang(userId, m.Flag)
	return nil
}
