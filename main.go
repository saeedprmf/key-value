package main

import (
	"fmt"
	"key-value/configs"
	"key-value/core"
	"os"
	"os/signal"
	"syscall"
)


//TODO: handle errors and dont use panic()


func exit(sigchan <- chan os.Signal){
    sig := <- sigchan
    fmt.Println(sig)

    core.ShotDown()

}



func main(){
    core.Log("starting app ...")
    core.Init()
    configs.Init()

    sigchan := make(chan os.Signal)
    signal.Notify(sigchan , syscall.SIGINT , syscall.SIGTERM)

    exit(sigchan)

}


