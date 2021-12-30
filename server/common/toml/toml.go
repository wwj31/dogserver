package toml

import (
	"fmt"

	"github.com/wwj31/dogactor/expect"

	"github.com/spf13/cast"

	"github.com/spf13/viper"
)

var (
	BaseConfig map[string]interface{}
	Config     map[string]interface{}
)

func Init(path string, appType string, appId int) {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	expect.Nil(err)

	BaseConfig = viper.GetStringMap("base")
	server := viper.GetStringMap("server")
	name := appType + "_" + cast.ToString(appId)
	v, ok := server[name]
	if !ok {
		panic(fmt.Errorf("not find %v_%v", appType, appId))
	}
	Config = v.(map[string]interface{})
}

func Get(k string) string {
	v, ok := Config[k]
	if !ok {
		if v, ok = BaseConfig[k]; !ok {
			panic(fmt.Errorf("common Config not find k %v", k))
		}
	}
	return v.(string)
}
