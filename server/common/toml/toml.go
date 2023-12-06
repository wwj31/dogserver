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
		Config = map[string]interface{}{}
	} else {
		Config = v.(map[string]interface{})
	}

}

func GetBool(k string, defaultValue ...bool) bool {
	v, ok := Config[k]
	if !ok {
		if v, ok = BaseConfig[k]; !ok {
			if len(defaultValue) > 0 {
				return defaultValue[0]
			}

			panic(fmt.Errorf("common Config not find k %v", k))
		}
	}
	return v.(bool)
}

func Get(k string, defaultValue ...string) string {
	v, ok := Config[k]
	if !ok {
		if v, ok = BaseConfig[k]; !ok {
			if len(defaultValue) > 0 {
				return defaultValue[0]
			}

			panic(fmt.Errorf("common Config not find k %v", k))
		}
	}
	return v.(string)
}

func GetArray(k string, defaultValue ...string) []string {
	v, ok := Config[k]
	if !ok {
		if v, ok = BaseConfig[k]; !ok {
			if len(defaultValue) > 0 {
				return defaultValue
			}

			panic(fmt.Errorf("common Config not find k %v", k))
		}
	}
	arr := v.([]interface{})
	array := make([]string, 0, len(arr))
	for _, v := range arr {
		array = append(array, v.(string))
	}
	return array
}

func GetB(k string) (string, bool) {
	v, ok := Config[k]
	if !ok {
		if v, ok = BaseConfig[k]; !ok {
			return "", false
		}
	}
	return v.(string), true
}
