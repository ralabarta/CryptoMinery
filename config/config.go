package config

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
)

var (
	env string
)

func init() {
	flag.StringVar(&env, "e", "TEST", "env variable")
}

type ViperHandle struct {
	V      *viper.Viper
	Config *AppConfig
}

func New() ViperHandle {
	return ViperHandle{
		V: viper.New(),
	}
}

var Config *AppConfig

// InitConfig 初始化配置文件
func (v *ViperHandle) InitConfig(path string) (err error) {
	flag.Parse()
	v.V.SetConfigName(env)            // name of config file (without extension)
	v.V.SetConfigType("yaml")         // REQUIRED if the config file does not have the extension in the name
	v.V.AddConfigPath(path + "/conf") // call multiple times to add many search paths
	err = v.V.ReadInConfig()          // Find and read the config file
	if err != nil {                   // Handle errors reading the config file
		panic(fmt.Errorf("Viper read config error: %s \n", err))
	}
	var ret AppConfig
	if err := v.V.Unmarshal(&ret); err != nil {
		panic(fmt.Errorf("Viper unmarsha1 config error: %s \n", err))
	}
	v.Config = &ret
	Config = &ret
	return
}
