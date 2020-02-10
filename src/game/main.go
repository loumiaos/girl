package main

import (
	"game/config"

	"game/dblogic"
	"game/dbmodel"
	"game/login"
	"game/world"

	"github.com/snowyyj001/loumiao"
	"github.com/snowyyj001/loumiao/dbbase/mysqldb"

	//"github.com/snowyyj001/loumiao/dbbase/redisdb"
	"github.com/snowyyj001/loumiao/gate"
	"github.com/snowyyj001/loumiao/log"

	lconf "github.com/snowyyj001/loumiao/config"
)

func init() {
	lconf.NET_PROTOCOL = config.NET_PROTOCOL
	lconf.NET_WEBSOCKET = config.NET_WEBSOCKET

	lconf.MYSQL_URI = config.MYSQL_URI
	lconf.MYSQL_DBNAME = config.MYSQL_DBNAME
	lconf.MYSQL_ACCOUNT = config.MYSQL_ACCOUNT
	lconf.MYSQL_PASS = config.MYSQL_PASS
}

func main() {
	log.Info("game run!")

	mysqldb.Dial(dbmodel.Models)
	//redisdb.DialDefault()

	loumiao.Prepare(new(gate.GateServer), "GateServer", false)
	loumiao.Prepare(new(login.LoginServer), "LoginServer", false)
	loumiao.Prepare(new(dblogic.DB_Player), "DB_Player", true)
	loumiao.Prepare(new(world.GameServer), "GameServer", false)

	loumiao.Run()
}
