// 机器人服务
package logic

import (
	"log"
	"strconv"
	"strings"

	"github.com/snowyyj001/loumiao/gorpc"
	"github.com/snowyyj001/loumiao/network"
)

type RoomInfo struct {
	roomId int
}

type Robot struct {
	gorpc.GoRoutineLogic
	client     *network.WebClient //如果使用socket，请使用ClientSocket
	m_ClientId int
	logintype  int
	userId     int
	tokenId    int
	roomInfo   RoomInfo
	roomState  int
	leftTime   int
	baseScore  int
}

func (self *Robot) Login() {
	self.logintype = 1
	self.client = new(network.WebClient)
	self.client.SetClientId(self.m_ClientId)
	self.client.Init("127.0.0.1", 4567)
	self.client.SetConnectType(network.CLIENT_CONNECT)
	self.client.BindPacketFunc(PacketFunc)
	self.client.Start()
}

func (self *Robot) LoginServer(addr string) {
	self.logintype = 2
	self.client = new(network.WebClient)
	self.client.SetClientId(self.m_ClientId)
	str := strings.Split(addr, ":")
	port, _ := strconv.Atoi(str[1])
	self.client.Init(str[0], port)
	self.client.SetConnectType(network.CLIENT_CONNECT)
	self.client.BindPacketFunc(PacketFunc)
	self.client.Start()
}

func (self *Robot) DoStart() {
	log.Println("%s DoStart", self.Name)

	self.RunTicker(1000, self.Update)
}

func (self *Robot) Update(igo gorpc.IGoRoutine, data interface{}) interface{} {
	//dt := data.(int64)

	return nil
}
