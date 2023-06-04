package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @success	200	{object}	jsonresult.JSONResult{data=proto.Order}	"desc"
type JSONResult struct {
	Status  int    `json:"status" `
	Message string `json:"message"`
}

func FlictResponse(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"msg": msg,
	})
}
