package db

import (
	"dbmodel"

	"github.com/snowyyj001/loumiao/dbbase/mysqldb"
	"github.com/snowyyj001/loumiao/gorpc"
)

func getPlayer(igo gorpc.IGoRoutine, data interface{}) interface{} {
	var useid = data

	user := dbmodel.User{}
	mysqldb.DB.First(&user, useid)

	return user
}

func getGameCfg(igo gorpc.IGoRoutine, data interface{}) interface{} {
	cfgs := make([]dbmodel.GameCfg, 1)
	mysqldb.DB.Find(&cfgs)
	return cfgs
}
