package configs

import (
	"encoding/json"
	"key-value/core"
	"log"
	"os"
)


func Init(){
	core.Log("start initializing configs of app ... ")
	f , err := os.ReadFile("configs/configs.json")
	if err != nil{
		panic(err)
	}
	
	err = json.Unmarshal(f , &config)
	if err != nil{
		panic(err)
	}
	log.Println(config)
	core.Log("inishialzing of configs ended")
	
}