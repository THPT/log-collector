package routes

import (
	"fmt"
	"log-collector/app/model"

	"github.com/gin-gonic/gin"
)

func (r *Router) ReciveEventLog(c *gin.Context) {
	event := model.ParseEvent(c)
	err, status := eventEntity.SendEvent(&event)
	if err != nil {
		c.AbortWithStatus(status)
	}
	fmt.Printf("%+v\n", event)
	c.JSON(200, event)
}
