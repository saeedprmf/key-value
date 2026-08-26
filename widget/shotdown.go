package widget



func ShotDown(){
	WriteFileApend(applog.buf.String() , "log.txt")
	applog.buf.Reset()
	applog.counter = 0
}