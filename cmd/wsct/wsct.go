package main

import (
	"github.com/shimokp/W-Shift-Trial-Server"
)

func main() {
	//var (
	//	addr   = flag.String("addr", ":8080", "addr to bind")
	//	dbconf = flag.String("dbconf", "dbconfig.yml", "database configuration file.")
	//	env    = flag.String("env", "development", "application envirionment (production, development etc.)")
	//	debug  = flag.Bool("debug", false, "debug mode. default is false.")
	//)
	//flag.Parse()
	//b := wsct.New()
	//b.Init(*dbconf, *env, *debug)
	//b.Run(*addr)

	s := wsct.New()
	s.SetupDB()

	s.SetupRouter()

	//err := wsct.DbTest()
	//log.Println(err)

}
