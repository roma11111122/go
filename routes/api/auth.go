package api

import (
	"cash/pakage/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAuth(c *gin.Context) {

	resp := response.Gin{C: c}
	resp.Response(http.StatusOK, response.SUCCESS, map[string]string{
		"token": "ssssdss",
	})
}
