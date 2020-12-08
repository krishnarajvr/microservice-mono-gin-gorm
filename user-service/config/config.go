package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joeshaw/envdecode"
)

type Conf struct {
	Debug  bool `env:"DEBUG,required"`
	Server serverConf
	Db     dbConf
	Log    logConf
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

func AppConfig() *Conf {
	var c Conf

	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Print("Not able to get current working director")
	}

	fmt.Println(c)

	if len(c.Log.LogFilePath) <= 0 {
		c.Log.LogFilePath = dir + "/log"
	}
	if len(c.Log.LogFileName) <= 0 {
		c.Log.LogFileName = "micro.log"
	}

	fmt.Println(c)

	return &c
}
