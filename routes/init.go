package routes

import (
	"log-collector/app/entity"
	"log-collector/config"
	"log-collector/middleware"
	"time"

	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
)

var (
	eventEntity entity.EventInterface
)

type Router struct {
}

func GetEngine() *gin.Engine {
	eventEntity = entity.NewEventEntity()

	// Set up gin
	gin.SetMode(config.AppMode)
	app := gin.New()
	app.Use(gzip.Gzip(gzip.DefaultCompression))
	app.Use(middleware.CORS())

	// Setup router
	router := &Router{}
	groupPing := app.Group("")
	groupPing.GET("/ping", router.Ping)

	groupLog := app.Group("/logs")
	groupLog.POST("", router.ReciveEventLog)
	return app
}

func (r *Router) Ping(c *gin.Context) {
	c.JSON(200, gin.H{"res": "pong", "time": time.Now().String()})
}
