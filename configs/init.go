package configs

import (
	"encoding/json"
	"key-value/widget"
	"time"
)


type InitError struct{
	msg string
	data string
}
func (ir *InitError)Error()string{
	return ir.msg + "\n" + ir.data
}







func Init() error {
	widget.Log("start initializing configs of app ... " , "[INFO] : " , widget.Green)
	d , err := widget.ReadFile("configs/configs.json")
	if err != nil{
		switch err.Error() {
			case "file is not exist":
				widget.Log("config.json not found try to find ..." , "[ERROR] : " , widget.Red)
				for i := 0 ; i < 12 ; i ++ {
					time.Sleep(time.Second*5)
					d , err = widget.ReadFile("configs/configs.json")
					if err != nil {
						widget.Log("config.json not found try to find ..." , "[ERROR] : " , widget.Red)
						continue
					} else {
						widget.Log("config.jdon is found" , "[OK] : " , widget.Green)
						break
					}
				}
				if err != nil{
					panic("config.json not found")
				}

		}
	}


	if err != nil {
		widget.Log(err.Error() , "[ERROR] : " , widget.Red)
	}



	err = json.Unmarshal(d , &config)
	if err != nil{
		return &InitError{"there are some errors with load json datas from configs.json pleas check it and press enter ..." , err.Error()}
	}

	widget.Log("inishialzing of configs ended" , "[OK] : " , widget.Green)
	return nil
}