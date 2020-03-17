//牛牛服务
package niuniu

import (
	"github.com/snowyyj001/loumiao"
	"github.com/snowyyj001/loumiao/gorpc"
	"github.com/snowyyj001/loumiao/log"
)

func GetRoomId(gameId int, idx int) int {
	return gameId*1000 + idx
}

var (
	This *GameServer
)

type GameServer struct {
	gorpc.GoRoutineLogic

	GameId   int
	GameRule string
	RoomNum  int
	Rooms    map[int]*Room
	Players  map[int]int
	Agents   map[int]int
}

func (self *GameServer) DoInit() {
	log.Infof("%s DoInit", self.Name)
	This = self

	loumiao.RegisterNetHandler(self, "C_R_SitDown", handlerSitDown)

}

func (self *GameServer) DoRegsiter() {
	self.Register("joinRoom", self.handlerJoinRoom)
	self.Register("disconnect", self.handlerLogOut)
}

func (self *GameServer) DoStart() {
	log.Infof("%s DoStart", self.Name)

	self.Agents = make(map[int]int)
	self.Players = make(map[int]int)
	self.Rooms = make(map[int]*Room)
	for i := 0; i < self.RoomNum; i++ {
		roomid := GetRoomId(self.GameId, i+1)
		self.Rooms[roomid] = new(Room)
		self.Rooms[roomid].doStart(roomid)
	}

	self.RunTicker(1000, self.Update)
}

func (self *GameServer) DoDestory() {
	log.Info("niuniu GameServer destory")
}

func (self *GameServer) GetRoomByClientId(clientId int) (int, *Room) {
	userId := self.Agents[clientId]
	roomId := self.Players[userId]
	return userId, This.Rooms[roomId]
}

func (self *GameServer) GetRoomByUserId(userId int) *Room {
	roomId := self.Players[userId]
	return This.Rooms[roomId]
}

func (self *GameServer) GetRoomByRoomId(roomId int) *Room {
	return This.Rooms[roomId]
}

func (self *GameServer) Update(igo gorpc.IGoRoutine, data interface{}) interface{} {
	dt := data.(int64)
	for _, room := range self.Rooms {
		room.update(dt)
	}
	return nil
}
