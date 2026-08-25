package main

import (
	"key-value/configs"
	"key-value/core"
)

//TODO: make an expection for shoting down safety 
//TODO: handle errors and dont use panic()

func main(){
    core.Log("starting app ...")
    core.Init()
    configs.Init()

    core.ShotDown()

}


