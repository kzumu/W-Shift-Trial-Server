package main

import (
	"log"

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
	err := s.SetupDB()
	if err != nil {
		log.Println(err)
		panic(err)
	}

	s.SetupRouter()

	//err := wsct.DbTest()
	//log.Println(err)

}
