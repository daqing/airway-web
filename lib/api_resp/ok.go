package api_resp

import (
	"github.com/gin-gonic/gin"
)

func OK(c *gin.Context, data any) {
	c.JSON(200, gin.H{"code": 200, "data": data, "message": ""})
}
