package main

import "fmt"

var tmplMainGo = []byte(`package main

import (
    "fmt"
)

func main() {
    fmt.Println("hello world")
}
`)

func GetBasicTmplMainGo() []byte {
	return tmplMainGo
}

var tmplConfigYAML = []byte(`# auto-generated version field
version: v1.0
`)

func GetBasicTmplConfigYAML() []byte {
	return tmplConfigYAML
}

var tmplConfigYAMLGo = []byte(`package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"runtime/debug"
)

var Cfg *viper.Viper

func init() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
			os.Exit(-1)
		}
	}()
	var e error

	Cfg = viper.New()

	Cfg.SetConfigFile("%s")
	if e = Cfg.ReadInConfig(); e != nil {
		panic(e)
	}
	fmt.Println(Cfg.GetString("version"))
}`)

func GetBasicTmplConfigYAMLGO(filePath string) []byte {
	return []byte(fmt.Sprintf(string(tmplConfigYAMLGo), filePath))
}
