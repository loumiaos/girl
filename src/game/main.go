package main

import (
	"game/config"
	"game/room"

	"dbmodel"
	"game/db"
	"game/gate"
	"game/world"

	"github.com/snowyyj001/loumiao"
	"github.com/snowyyj001/loumiao/dbbase/mysqldb"

	//"github.com/snowyyj001/loumiao/dbbase/redisdb"
	"github.com/snowyyj001/loumiao/log"

	lconf "github.com/snowyyj001/loumiao/config"
)

func init() {
	lconf.MYSQL_URI = config.MYSQL_URI
	lconf.MYSQL_DBNAME = config.MYSQL_DBNAME
	lconf.MYSQL_ACCOUNT = config.MYSQL_ACCOUNT
	lconf.MYSQL_PASS = config.MYSQL_PASS
}

func main() {
	log.Info("server run!")
	//db connect
	mysqldb.Dial(dbmodel.Models)
	//redisdb.DialDefault()

	//start watch dog
	gate.StartGate()

	//service start
	loumiao.Prepare(new(db.DBServer), "DBServer", true)
	loumiao.Prepare(new(world.WorldServer), "WorldServer", false)
	loumiao.Prepare(new(room.RoomServer), "RoomServer", false)

	loumiao.Run()
}
