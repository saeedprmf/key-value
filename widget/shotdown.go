package widget

import (
	"os"
	"os/exec"
	"time"
)


func savelog(){
	data , err := ReadFile("log.txt")
	if err != nil {
		switch err.Error(){
			case "file is not exist":
				Log("log file is not exist" , "[ERROR] :" , Red)
				return
		}
	}
	name := "./logs/log_" + time.Now().Format("2006-01-02_15-04-05") + ".txt"
	err = os.WriteFile(name , data , 0664)
	if err != nil {
		Log("there is some errors with saving log file" , "[ERROR] :" , Red)
		return
	}
	Log("log file is saved as " + name , "[INFO] :" , Green)
	cmd := exec.Command("rm" , "log.txt")
	err = cmd.Run()
	if err != nil {
		Log("there is some errors with deleting log file" , "[ERROR] :" , Red)
		return
	}
	Log("log file is deleted" , "[INFO] :" , Green)
}



func ShutDown(){
	WriteFileApend(applog.buf.String() , "log.txt")
	applog.buf.Reset()
	applog.counter = 0
	savelog()
}