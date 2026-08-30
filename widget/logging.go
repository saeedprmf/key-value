package widget

import (
	"bytes"
	"log"
)


const (
	Red = "\033[31m"
	Reset = "\033[0m"
	Green = "\033[32m"
)



type LogBase struct{
	buf bytes.Buffer
	logger *log.Logger
	counter int
}


func NewAppLog() *LogBase{
	var applog LogBase = LogBase{}
	applog.logger = log.New(&applog.buf , "" , log.LstdFlags|log.Lmicroseconds)
	return &applog
}

var applog = NewAppLog()



func Log(data any , tag string, color string){
	applog.logger.Println(data)
	
	log.Println(color , tag , Reset , data)

	applog.counter += 1
	if applog.counter >= 3{
		WriteFileApend(applog.buf.String() , "log.txt")
		applog.buf.Reset()
		applog.counter = 0
	}
	

}


