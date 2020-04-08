// 机器人服务
package logic

import (
	"fmt"
	"robot/config"
	"robot/msg"
	"strconv"

	"github.com/snowyyj001/loumiao"
	"github.com/snowyyj001/loumiao/gate"
	"github.com/snowyyj001/loumiao/log"
	"github.com/snowyyj001/loumiao/message"
)

var (
	robots   map[int]*Robot
	hanlders map[string]func(robot *Robot, data interface{})
)

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

	hanlders["NN_RC_TableInfo"] = onRCTableInfo
	hanlders["R_C_JoinRoom"] = onRCJoinRoom
	hanlders["R_C_SyncPlayers"] = onRCSync_Players
	hanlders["R_C_Ready"] = onRC_Ready
	hanlders["NN_RC_FaPai"] = onRC_FaPai
}

func init() {
	robots = make(map[int]*Robot)
	hanlders = make(map[string]func(robot *Robot, data interface{}))
	register()
}

func onConnect(robot *Robot, data interface{}) {
	log.Debugf("onConnect: %d", robot.m_ClientId)
	if robot.logintype == 1 {
		login := new(msg.C_A_Login)
		login.AccountName = fmt.Sprintf("%d:账户", robot.m_ClientId)
		login.Password = "123456"
		login.Channel = 1
		login.LoginType = 1
		login.Sex = 0
		login.NickName = fmt.Sprintf("%d:昵称", robot.m_ClientId)
		login.HeadIcon = "http://thirdwx.qlogo.cn/mmopen/vi_32/bhe854X0b6hLdxAHIlsqh3EhlWGO2HaDmgYadP5oZmzhpqjks089dwJjj1doJjKOAZYOxbHSP3bSXRqTUf06Pg/132"
		log.Debugf("发送登陆消息: %v", login)
		robot.client.SendMsg("C_A_Login", login)
	} else {
		var req = gate.LouMiaoLoginGate{UserId: robot.userId, TokenId: robot.tokenId}
		robot.client.SendMsg("LouMiaoLoginGate", req)

		login := new(msg.C_S_Login)
		login.UserID = robot.userId
		robot.client.SendMsg("C_S_Login", login)
		log.Debugf("1.发送登陆消息: %v", login)
	}

}

func onLoginAccount(robot *Robot, data interface{}) {
	resp := data.(*msg.A_C_Login)
	log.Debugf("onLoginAccount: %v", resp)
	robot.client.Stop()
	robot.userId = resp.UserID
	robot.tokenId = resp.TokenID
	robot.LoginServer(resp.GateAddr)
}

func onLoginPlayer(robot *Robot, data interface{}) {
	resp := data.(*msg.S_C_Login)
	log.Debugf("onLoginPlayer: %v", resp)

	log.Debugf("player登陆成功 %d", robot.m_ClientId)

	req := new(msg.C_S_JoinRoom)
	req.RoomId = 101001
	req.Service = "niuniu"
	robot.client.SendMsg("C_S_JoinRoom", req)

	loumiao.Start(robot, strconv.Itoa(robot.userId), false)
}

func onJoinRoom(robot *Robot, data interface{}) {
	resp := data.(*msg.S_C_JoinRoom)
	log.Debugf("onJoinRoom: %v", resp)
}

func onRCTableInfo(robot *Robot, data interface{}) {
	resp := data.(*msg.NN_RC_TableInfo)
	log.Debugf("onRCTableInfo: %v", resp)
	robot.roomState = resp.State
	robot.leftTime = resp.LeftTime
	robot.baseScore = resp.BaseScore
}

func onRCJoinRoom(robot *Robot, data interface{}) {
	resp := data.(*msg.R_C_JoinRoom)
	log.Debugf("onRCJoinRoom: %v", resp)
}

func onRCSync_Players(robot *Robot, data interface{}) {
	resp := data.(*msg.R_C_SyncPlayers)
	fmt.Println("onRCSync_Players: %v", resp.Players[0])

	for _, player := range resp.Players {
		if robot.roomState == 0 {
			if player.State == 0 {
				req := new(msg.C_R_Ready)
				robot.client.SendMsg("C_R_Ready", req)
				break
			}
		}
	}
}

func onRC_Ready(robot *Robot, data interface{}) {
	resp := data.(*msg.R_C_Ready)
	fmt.Println("onRC_Ready: %v", resp)
}

func onRC_FaPai(robot *Robot, data interface{}) {
	resp := data.(*msg.NN_RC_FaPai)
	fmt.Println("onRC_FaPai: %v", resp)
}
