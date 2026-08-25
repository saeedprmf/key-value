package configs

import (
	"encoding/json"
	"key-value/core"
	"key-value/widget"
	"log"
)


func Init(){
	core.Log("start initializing configs of app ... ")
	d , err := widget.ReadFile("configs/configs.json")
	
	err = json.Unmarshal(d , &config)
	if err != nil{
		panic(err)
	}
	log.Println(config)
	core.Log("inishialzing of configs ended")
	
}