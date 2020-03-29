// 机器人服务
package logic

import (
	"fmt"
	"robot/config"
	"robot/msg"

	"github.com/snowyyj001/loumiao/log"
	"github.com/snowyyj001/loumiao/message"
	"github.com/snowyyj001/loumiao/network"
)

var (
	robots   map[int]*Robot
	hanlders map[string]func(robot *Robot, data interface{})
)

type Robot struct {
	client     *network.WebClient //如果使用socket，请使用ClientSocket
	m_ClientId int
}

func PacketFunc(socketid int, buff []byte, nlen int) bool {
	defer func() {
		if err := recover(); err != nil {
			log.Debugf("MsgProcess PacketFunc %v", err)
		}
	}()
	err, name, pm := message.Decode(buff, nlen)
	if err != nil {
		return false
	}

	//log.Debugf("PacketFunc: %d, %s, %v", socketid, name, pm)

	hd, has := hanlders[name]
	if has {
		robot, has := robots[socketid]
		if has == false {
			log.Infof("client has been deleted %d", socketid)
		} else {
			hd(robot, pm)
		}
	}
	return true
}

func (self *Robot) Login() {
	self.client = new(network.WebClient)
	self.client.SetClientId(self.m_ClientId)
	self.client.Init("127.0.0.1", 4567)
	self.client.SetConnectType(network.CLIENT_CONNECT)
	self.client.BindPacketFunc(PacketFunc)
	self.client.Start()
}
func Start() {
	for i := 0; i < config.ROBOT_NUMBER; i++ {
		robots[i] = &Robot{m_ClientId: i}
		go robots[i].Login()
	}
}

func register() {
	hanlders["CONNECT"] = onConnect
	hanlders["A_C_Login"] = onLoginAccount
	hanlders["S_C_Login"] = onLoginPlayer
	hanlders["S_C_JoinRoom"] = onJoinRoom

	hanlders["R_C_JoinRoom"] = onRCJoinRoom
	hanlders["R_C_Sync_Players"] = onRCSync_Players
}

func init() {
	robots = make(map[int]*Robot)
	hanlders = make(map[string]func(robot *Robot, data interface{}))
	register()
}

func onConnect(robot *Robot, data interface{}) {
	log.Debugf("onConnect: %d", robot.m_ClientId)

	login := new(msg.C_A_Login)
	login.AccountName = fmt.Sprintf("%d:账户", robot.m_ClientId)
	login.Password = "123456"
	login.Channel = 1
	login.LoginType = 1
	login.Sex = 0
	login.NickName = fmt.Sprintf("%d:昵称", robot.m_ClientId)
	login.HeadIcon = "www.baidu.com"
	log.Debugf("发送登陆消息: %v", login)
	robot.client.SendMsg("C_A_Login", login)
}

func onLoginAccount(robot *Robot, data interface{}) {
	resp := data.(*msg.A_C_Login)
	log.Debugf("onLoginAccount: %v", resp)

	login := new(msg.C_S_Login)
	login.UserID = resp.UserID
	robot.client.SendMsg("C_S_Login", login)
}

func onLoginPlayer(robot *Robot, data interface{}) {
	resp := data.(*msg.S_C_Login)
	log.Debugf("onLoginPlayer: %v", resp)

	log.Debugf("player登陆成功 %d", robot.m_ClientId)

	req := new(msg.C_S_JoinRoom)
	req.RoomId = 101001
	req.Service = "niuniu"
	robot.client.SendMsg("C_S_JoinRoom", req)
}

func onJoinRoom(robot *Robot, data interface{}) {
	resp := data.(*msg.S_C_JoinRoom)
	log.Debugf("onJoinRoom: %v", resp)

}

func onRCJoinRoom(robot *Robot, data interface{}) {
	resp := data.(*msg.R_C_JoinRoom)
	log.Debugf("onRCJoinRoom: %v", resp)
}

func onRCSync_Players(robot *Robot, data interface{}) {
	resp := data.(*msg.R_C_Sync_Players)
	fmt.Println("onRCSync_Players: %v", resp.Players[0])
}
