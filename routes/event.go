package routes

import (
	"log-collector/app/model"

	"github.com/gin-gonic/gin"
)

func (r *Router) ReciveEventLog(c *gin.Context) {
	event := model.Event{}
	c.BindJSON(&event)
	event.CityId = 50
	err, status := eventEntity.SendEvent(&event)
	if err != nil {
		c.AbortWithStatus(status)
	}
	c.AbortWithStatus(200)
}
