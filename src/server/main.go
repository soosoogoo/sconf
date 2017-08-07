package main

import (
	"server/conf"
	"server/game"
	"server/gate"
	"server/login"

	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"

	"server/db/mysql"

	"fmt"
)

func main() {

	var md = new(mysql.MysqlDriver)
	var userInfo = md.Test()
	fmt.Println(userInfo.Username)
	//fmt.Println(conf.Server.LogLevel)

	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)

}
