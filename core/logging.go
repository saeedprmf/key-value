package core

import (
	"bytes"
	"key-value/widget"
	"log"
)

var buflog bytes.Buffer
var applog = log.New(&buflog , "" , log.LstdFlags|log.Lmicroseconds)

func Log(data string){
	buflog.Reset()
	applog.Println(data)
	widget.WriteFile(buflog.String() , "log.txt")

}


