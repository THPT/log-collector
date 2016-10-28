package routes

import (
	"log-collector/app/entity"
	"log-collector/config"
	"log-collector/middleware"

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
	group := app.Group("/logs")
	group.POST("", router.ReciveEventLog)
	return app
}
