package config

import (
	"flag"
	"log"
	"os"
)

var (
	Log *log.Logger
)

func init() {
	dir, _ := os.Getwd()
	// set location of log file
	var logpath = dir + "/storage/logs/info.log"

	flag.Parse()
	var file, err1 = os.Create(logpath)

	if err1 != nil {
		panic(err1)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	Log.Println("LogFile : " + logpath)
}
