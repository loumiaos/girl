package db

import (
	"dbmodel"
	"login/config"
	"login/msg"

	"github.com/snowyyj001/loumiao/dbbase/mysqldb"
	"github.com/snowyyj001/loumiao/gorpc"
	"github.com/snowyyj001/loumiao/util"
)

func loginAccount(igo gorpc.IGoRoutine, data interface{}) interface{} {
	var m = data.(*msg.C_A_Login)

	query := struct {
		ID        int
		Passwd    string
		Channel   int
		LoginType int
	}{}

	mysqldb.DB.Raw("select id,passwd,channel,login_type from user where account = ?", m.AccountName).Scan(&query)
	if query.ID == 0 {
		user := dbmodel.User{}
		user.Account = m.AccountName
		user.HeadIconUrl = m.HeadIcon
		user.Passwd = m.Password
		user.NickName = m.NickName
		user.Channel = int(m.Channel)
		user.LoginType = int(m.LoginType)
		user.Sex = int(m.Sex)
		user.Channel = int(m.Channel)
		user.ActiveFlag = 0
		user.Gold = config.PLAYER_GOLD
		user.Coin = config.PLAYER_COIN
		user.Money = config.PLAYER_MONEY
		user.RegisterTime = util.TimeStamp()
		user.LoginTime = user.RegisterTime
		user.LogoutTime = 0
		user.UnderWrite = ""
		user.IpAddr = ""
		mysqldb.DB.Create(&user)

		query.ID = user.ID
	} else {
		if query.Passwd != m.Password || query.Channel != int(m.Channel) || query.LoginType != int(m.LoginType) {
			return 0
		}
		if len(m.NickName) > 0 || len(m.HeadIcon) > 0 {
			mysqldb.DB.Exec("update user set nick_name = ?, head_icon_url = ? where id = ?", m.NickName, m.HeadIcon, query.ID)
		}
	}

	return query.ID
}
