package main

import (
	"dbmodel"
	"login/config"
	"login/db"
	"login/gate"
	"login/login"

	"github.com/snowyyj001/loumiao"
	"github.com/snowyyj001/loumiao/dbbase/mysqldb"
	"github.com/snowyyj001/loumiao/log"

	//"github.com/snowyyj001/loumiao/dbbase/redisdb"

	lconf "github.com/snowyyj001/loumiao/config"
)

func init() {
	lconf.MYSQL_URI = config.MYSQL_URI
	lconf.MYSQL_DBNAME = config.MYSQL_DBNAME
	lconf.MYSQL_ACCOUNT = config.MYSQL_ACCOUNT
	lconf.MYSQL_PASS = config.MYSQL_PASS
}

func main() {
	log.Info("login server run!")
	//db connect
	mysqldb.Dial(dbmodel.Models)
	//redisdb.DialDefault()

	gate.StartGate()

	//service start
	loumiao.Prepare(new(db.DBServer), "DBServer", true)
	loumiao.Prepare(new(login.LoginServer), "LoginServer", false)

	loumiao.Run()
}
