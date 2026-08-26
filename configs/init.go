package configs

import (
	"encoding/json"
	"key-value/widget"
	"log"
)


type InitError struct{
	msg string
	data string
}
func (ir *InitError)Error()string{
	return ir.msg + "\n" + ir.data
}





func Init() error{
	widget.Log("start initializing configs of app ... ")
	d , err := widget.ReadFile("configs/configs.json")

	if err != nil && err.Error() == "file is not exist"{
		return &InitError{"configs.json file dont found please check it and press enter ..." , ""}
	}

	err = json.Unmarshal(d , &config)
	if err != nil{
		return &InitError{"there are some errors with load json datas from configs.json pleas check it and press enter ..." , err.Error()}
	}

	log.Println(config)
	widget.Log("inishialzing of configs ended")
	return nil
}