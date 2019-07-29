package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"testing"
)

func TestModuleTmpl(t *testing.T) {
	//var v = struct{
	//	Package string `yaml:"package"`
	//	DirList []string `yaml:"dirList"`
	//} {
	//	"module",
	//	[]string{"{$object}Service", "{$object}Router"},
	//}
	//buf,e:= yaml.Marshal(v)
	//if e!=nil {
	//	fmt.Println(e.Error())
	//	return
	//}
	//
	//fmt.Println(string(buf))
	var m map[string]interface{}
	e :=yaml.Unmarshal([]byte(moduleTmpl), &m)
	if e!=nil {
		panic(e)
	}
	fmt.Println(m)
}
