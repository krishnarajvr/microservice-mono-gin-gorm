package config

import (
	"log"
	"path"
	"path/filepath"
	"runtime"
	"time"

	"github.com/joeshaw/envdecode"
	common "github.com/krishnarajvr/micro-common"
)

func init() {
	common.LoadEnv()
}

type Conf struct {
	Debug   bool `env:"DEBUG,required"`
	Server  serverConf
	Db      dbConf
	Log     logConf
	App     appConf
	Gateway gatewayConf
	Cache   cacheConf
	Token   tokenConf
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

type cacheConf struct {
	Host     string `env:"REDIS_HOST"`
	Username string `env:"REDIS_USER"`
	Password string `env:"REDIS_PASS"`
}

type tokenConf struct {
	AccessSecret        string `env:"ACCESS_SECRET"`
	RefreshSecret       string `env:"REFRESH_SECRET"`
	AdminAccessExpiry   int    `env:"ADMIN_ACCESS_EXPIRY"`   //In Minutes
	AdminRefreshExpiry  int    `env:"ADMIN_REFRESH_EXPIRY"`  //In Minutes
	ClientAccessExpiry  int    `env:"CLIENT_ACCESS_EXPIRY"`  //In Minutes
	ClientRefreshExpiry int    `env:"CLIENT_REFRESH_EXPIRY"` //In Minutes
}

type appConf struct {
	BaseURL     string `env:"APP_BASE_URL"`
	Lang        string `env:"APP_LANG"`
	Name        string `env:"SERVICE_NAME"`
	RootDir     string `env:"ROOT_DIR"`
	NewrelicKey string `env:"NEWRELIC_KEY"`
}

type gatewayConf struct {
	URL    string `env:"API_GATEWAY_URL"`
	Prefix string `env:"API_GATEWAY_PREFIX"`
}

func GetRootDir() string {
	_, b, _, _ := runtime.Caller(0)

	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func AppConfig() *Conf {
	var c Conf

	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}

	if len(c.App.RootDir) <= 0 {
		c.App.RootDir = GetRootDir()
	}

	if len(c.App.Lang) <= 0 {
		c.App.Lang = "en-US"
	}

	if len(c.App.BaseURL) <= 0 {
		c.App.BaseURL = "api/v1"
	}

	if len(c.Log.LogFilePath) <= 0 {
		c.Log.LogFilePath = c.App.RootDir + "/log"
	}

	if len(c.Log.LogFileName) <= 0 {
		c.Log.LogFileName = "micro.log"
	}

	if len(c.App.Name) <= 0 {
		c.App.Name = "MicroService"
	}

	//Access And Refresh Token Expiry in minutes
	if c.Token.AdminAccessExpiry == 0 {
		c.Token.AdminAccessExpiry = 60 * 2
	}

	if c.Token.AdminRefreshExpiry == 0 {
		c.Token.AdminRefreshExpiry = 60 * 24 * 2
	}

	if c.Token.ClientAccessExpiry == 0 {
		c.Token.ClientAccessExpiry = 60 * 3
	}

	if c.Token.ClientRefreshExpiry == 0 {
		c.Token.ClientRefreshExpiry = 60 * 24 * 3
	}

	return &c
}
