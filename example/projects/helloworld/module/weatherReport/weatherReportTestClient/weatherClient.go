package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	rsp,e:=http.Get("http://localhost:8080/weather/")
	if e!=nil {
		fmt.Println(e.Error())
		return
	}
	fmt.Println(rsp.Status)
	if rsp!=nil && rsp.Body!=nil{
		buf,_:=ioutil.ReadAll(rsp.Body)
		fmt.Println(string(buf))
	}
}

