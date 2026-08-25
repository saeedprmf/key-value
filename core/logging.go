package core

import (
	"bytes"
	"key-value/widget"
	"log"
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



func Log(data string){
	applog.logger.Println(data)
	applog.counter += 1
	if applog.counter >= 3{
		widget.WriteFileApend(applog.buf.String() , "log.txt")
		applog.buf.Reset()
		applog.counter = 0
	}
	

}


