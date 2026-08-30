package main

import (
	"key-value/configs"
	"key-value/core"
	"key-value/widget"
	"os"
	"os/signal"
	"syscall"
)

func exit(sigchan <-chan os.Signal) {
	sig := <-sigchan
	widget.ShutDown()
	widget.Log(sig , "[INFO] : " , widget.Green)

}

func main() {

	widget.Log("starting app ..." , "[INFO] : " , widget.Green)
	configs.Init()
	core.Init()

	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	exit(sigchan)

}
