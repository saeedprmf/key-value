package core

import (
	"key-value/widget"
)

func ShotDown(){
	widget.WriteFileApend(applog.buf.String() , "log.txt")
	applog.buf.Reset()
	applog.counter = 0
}