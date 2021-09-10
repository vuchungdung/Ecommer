package utilities

import "github.com/gin-gonic/gin"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Error(c *gin.Context, status int, data interface{}) {
	c.JSON(status, gin.H{"error": data})
}

func Success(c *gin.Context, status int, data interface{}) {
	c.JSON(status, gin.H{"success": data})
}
