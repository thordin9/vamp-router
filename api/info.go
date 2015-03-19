package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetInfo(c *gin.Context) {

	status, err := Runtime(c).GetInfo()
	if err != nil {
		HandleError(c, err)
	} else {
		c.JSON(http.StatusOK, status)
	}
}
