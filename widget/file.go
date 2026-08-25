package widget

import (
	"os"
)


func WriteFile(data string , fname string){
	f , err := os.OpenFile(fname , os.O_APPEND|os.O_WRONLY|os.O_CREATE,0664)
	if err != nil {
		panic(err)
	}
	_ , err = f.WriteString(data)
}

