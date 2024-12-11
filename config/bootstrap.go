package config

import (
	"bytes"
	"embed"
	"github.com/spf13/viper"
	"log"
	"os"
)

//go:embed *.yaml
var configs embed.FS

const CONF_DIR = "config/"

func init() {
	env := os.Getenv("ENV")
	vp := viper.New()
	// 根据环境变量 ENV 决定要读取的应用启动配置
	configFileStream, err := configs.ReadFile("application." + env + ".yaml")
	if err != nil {
		// 加载不到应用配置，阻挡应用的继续启动
		panic(err)
	}

	vp.SetConfigType("yaml")
	err = vp.ReadConfig(bytes.NewReader(configFileStream))
	if err != nil {
		// 加载不到应用配置，阻挡应用的继续启动
		panic(err)
	}

	err = vp.UnmarshalKey("app", &App)
	if err != nil {
		log.Panic("UnmarshalKey app failed:", err)
	}
	err = vp.UnmarshalKey("database", &Database)
	if err != nil {
		log.Panic("UnmarshalKey database failed:", err)
	}
	err = vp.UnmarshalKey("redis", &Redis)
	if err != nil {
		log.Panic("UnmarshalKey redis failed:", err)
	}
}
