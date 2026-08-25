package main

import (
	"context"
	"key-value/configs"
	"key-value/core"
	"os/signal"
	"syscall"
)


func main(){
    core.Log("starting app ...")
    core.Init()
    configs.Init()
    ctx , stop := signal.NotifyContext(context.Background(),syscall.SIGINT,syscall.SIGTERM)
    defer stop()
    

    <-ctx.Done()
    core.Log("shoting down app ...")
    core.Log("app shoted down by :)")

}


