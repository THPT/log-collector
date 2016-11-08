package routes

import (
	"log-collector/app/model"

	"github.com/gin-gonic/gin"
)

func (r *Router) ReciveEventLog(c *gin.Context) {
	event := model.ParseEvent(c)
	err, status := eventEntity.SendEvent(&event)
	if err != nil {
		c.AbortWithStatus(status)
	}
	c.JSON(200, event)
}
