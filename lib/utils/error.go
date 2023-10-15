package utils

import (
	"fmt"

	"github.com/daqing/airway/lib/resp"
	"github.com/gin-gonic/gin"
)

const DefaultMessage = "[Error for gin.Context]"

func LogError(c *gin.Context, err error) {
	LogErrorMsg(c, err, DefaultMessage)
}

func LogErrorMsg(c *gin.Context, err error, message string) {
	fmt.Println(message, "Got error: ", err)

	resp.Error(c, err)

	panic(message)
}
