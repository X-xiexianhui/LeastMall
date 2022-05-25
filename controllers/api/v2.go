package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetV2(c *gin.Context) {
	c.String(http.StatusOK, "api v2")
}
