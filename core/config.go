package core

import (
	"fmt"
	"grpc-demo/global"
	"log"
)
import (
	"github.com/spf13/viper"
)

const DefaultConfigFile = "./core/config.yaml"

func LoadConfig() {
	config := viper.New()
	config.SetConfigFile(DefaultConfigFile)
	if err := config.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	global.Configs = config.AllSettings()
	for key, value := range global.Configs {
		log.Printf("key=%s value=%s", key, value)
	}
}
