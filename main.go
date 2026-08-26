package main

import (
	"fmt"
	"key-value/configs"
	"key-value/core"
	"key-value/widget"
	"os"
	"os/signal"
	"syscall"
)



func exit(sigchan <- chan os.Signal){
    sig := <- sigchan
    fmt.Println(sig)

   

}



func main(){
    
    widget.Log("starting app ...")
    for{
        err := configs.Init()
        if err != nil{
            fmt.Println(err)
            fmt.Scanln()
        } else {
            break
        }
    }
    core.Init()
    

    sigchan := make(chan os.Signal)
    signal.Notify(sigchan , syscall.SIGINT , syscall.SIGTERM)

    exit(sigchan)

}


