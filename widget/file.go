package widget

import (
	"fmt"
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
		if os.IsNotExist(err){
			fmt.Println(fname , "is not exist pleas put it in corect directory")
			Log("file "+fname+" is not exist")
		} else {
			fmt.Println("there is some errors with opening ",fname," file")
			Log("unexpected error " + err.Error())
		}
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

