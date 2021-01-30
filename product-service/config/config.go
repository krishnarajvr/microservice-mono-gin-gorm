package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joeshaw/envdecode"
	common "github.com/krishnarajvr/micro-common"
)

func init() {
	common.LoadEnv()
}

type Conf struct {
	Debug  bool `env:"DEBUG,required"`
	Server serverConf
	Db     dbConf
	Log    logConf
	App    appConf
}

type serverConf struct {
	Port         string        `env:"SERVER_PORT,required"`
	TimeoutRead  time.Duration `env:"SERVER_TIMEOUT_READ,required"`
	TimeoutWrite time.Duration `env:"SERVER_TIMEOUT_WRITE,required"`
	TimeoutIdle  time.Duration `env:"SERVER_TIMEOUT_IDLE,required"`
}

type logConf struct {
	LogFilePath string `env:"Log_FILE_PATH"`
	LogFileName string `env:"LOG_FILE_NAME"`
}

type dbConf struct {
	Host     string `env:"DB_HOST,required"`
	Port     int    `env:"DB_PORT,required"`
	Username string `env:"DB_USER,required"`
	Password string `env:"DB_PASS,required"`
	DbName   string `env:"DB_NAME,required"`
}

type appConf struct {
	BaseURL string `env:"APP_BASE_URL"`
	Lang    string `env:"APP_LANG"`
}

func AppConfig() *Conf {
	var c Conf

	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Print("Not able to get current working director")
	}

	if len(c.App.Lang) <= 0 {
		c.App.Lang = "el-GR"
	}

	if len(c.App.BaseURL) <= 0 {
		c.App.BaseURL = "api/v1"
	}

	if len(c.Log.LogFilePath) <= 0 {
		c.Log.LogFilePath = dir + "/log"
	}

	if len(c.Log.LogFileName) <= 0 {
		c.Log.LogFileName = "micro.log"
	}

	fmt.Println(c)

	return &c
}
