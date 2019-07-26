package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
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

	Cfg.SetConfigFile("G:/go_workspace/GOPATH/src/pjx/helloworld/config/local.yaml")
	if e = Cfg.ReadInConfig(); e != nil {
		panic(e)
	}
	fmt.Println(Cfg.GetString("version"))
	Cfg.Set("name","ft")
	Cfg.WriteConfig()
}
