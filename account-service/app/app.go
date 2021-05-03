package app

import (
	"log"
	cache "micro/app/database/cache"
	lgorm "micro/app/database/gorm"
	"micro/config"
	utilCache "micro/util/cache"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/krishnarajvr/micro-common/locale"
	"github.com/krishnarajvr/micro-common/middleware"
	newrelic "github.com/newrelic/go-agent"
)

//AppConfig - Application config
type AppConfig struct {
	Dbs     *Dbs
	Lang    *locale.Locale
	Router  *gin.Engine
	BaseURL string
	Cfg     *config.Conf
}

type Dbs struct {
	DB    *gorm.DB
	Cache *utilCache.RedisClient
}

func NewrelicMiddleware(appName string, key string) gin.HandlerFunc {
	if appName == "" || key == "" {
		return func(c *gin.Context) {}
	}

	config := newrelic.NewConfig(appName, key)
	app, err := newrelic.NewApplication(config)

	if err != nil {
		panic(err)
	}

	return func(c *gin.Context) {
		txn := app.StartTransaction(c.Request.URL.Path, c.Writer, c.Request)
		defer txn.End()
		c.Next()
	}
}

// InitRouter - Create gin router
func InitRouter(cfg *config.Conf, excludeList map[string]interface{}) (*gin.Engine, error) {
	router := gin.Default()
	router.Use(middleware.LoggerToFile(cfg.Log.LogFilePath, cfg.Log.LogFileName))

	if len(cfg.App.NewrelicKey) != 0 {
		router.Use(NewrelicMiddleware(cfg.App.Name, cfg.App.NewrelicKey))
	}

	router.Use(middleware.TenantValidator(excludeList))
	router.Use(gin.Recovery())

	return router, nil
}

// InitLocale - Create locale object
func InitLocale(cfg *config.Conf) (*locale.Locale, error) {
	langLocale := locale.Locale{}
	dir, err := os.Getwd()

	if err != nil {
		log.Print("Not able to get current working director")
		panic(err)
	}

	lang := langLocale.New(cfg.App.Lang, dir+"/locale/*/*", "en-GR", "en-US", "zh-CN")

	return lang, nil
}

// InitDS establishes connections to fields in dataSources
func InitDS(config *config.Conf) (*Dbs, error) {
	db, err := lgorm.New(config)
	redisCache, err := cache.New(config)

	if err != nil {
		log.Println("Connection failed")
		panic(err)
	}

	return &Dbs{
		DB:    db,
		Cache: redisCache,
	}, nil
}

//Close to be used in graceful server shutdown
func Close(d *Dbs) error {
	//Todo Check the error
	//return d.DB.Close()
	return nil
}
