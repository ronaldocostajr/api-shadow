package logSystem

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	Log *log.Logger
)

func WriteLogFile(texto string) {
	// set location of log file
	var logpath = os.Getenv("logfile_path")
	//"logSystem/systemInfo.log"
	fmt.Println(logpath)
	flag.Parse()
	var file, err1 = os.OpenFile(logpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err1 != nil {
		panic(err1)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	Log.Println("LogFile " + texto)
}
