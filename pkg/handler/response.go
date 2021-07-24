package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

type errorResponse struct {
	Message string `json:"message"`
}

type resultResponse struct {
	Result string `json:"result"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Fatal(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}