package routes

import (
	"log-collector/app/entity"
	"log-collector/config"
	"log-collector/middleware"
	"time"

	lumberjack "gopkg.in/natefinch/lumberjack.v2"

	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
)

var (
	eventEntity   entity.EventInterface
	UnexpectedLog = "log/access.log"
)

type Router struct {
}

func GetEngine() *gin.Engine {
	eventEntity = entity.NewEventEntity()

	// Set up gin
	gin.SetMode(config.AppMode)
	app := gin.New()
	app.Use(gin.Recovery())
	app.Use(gin.LoggerWithWriter(NewLogWriter(UnexpectedLog + "_" + config.AppPort)))
	app.Use(gzip.Gzip(gzip.DefaultCompression))
	app.Use(middleware.CORS())

	// Setup router
	router := &Router{}
	groupPing := app.Group("")
	groupPing.GET("/ping", router.Ping)

	groupLog := app.Group("/track")
	groupLog.GET("", router.ReciveEventLog)
	return app
}

func (r *Router) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"service": config.AppPort,
		"time":    time.Now().String(),
	})
}

func NewLogWriter(log string) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   log,
		MaxSize:    50, // megabytes
		MaxBackups: 1000,
		MaxAge:     30, // days
	}
}
