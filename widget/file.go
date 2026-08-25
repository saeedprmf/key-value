package widget

import (
	"os"
)


type ReadError struct {
	msg string
}
func (re *ReadError)Error()string{
	return re.msg
}


func WriteFileApend(data string , fname string){
	f , err := os.OpenFile(fname , os.O_APPEND|os.O_WRONLY|os.O_CREATE,0664)
	if err != nil {
		panic(err)
	}
	_ , err = f.WriteString(data)
}

func ReadFile(fname string)([]byte , error){
	data , err := os.ReadFile(fname)
	if err != nil{
		if os.IsNotExist(err){
			return []byte("") , &ReadError{"file is not exist"}
		} else {
			return []byte("") , &ReadError{"there is an unexpected error"}
		}
	}
	return data , nil
}

