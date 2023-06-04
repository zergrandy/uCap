package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary		健康檢查
// @Description
// @Tags			System
// @version		1.0
// @produce		application/json
// @Security Bearer
// @Success		200		{object}	controllers.JSONResult
// @Router		/api/v1/healthcheck [get]
func Healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "TOKEN OK",
	})
}

func PING(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"PONG": time.Now(),
	})
}
