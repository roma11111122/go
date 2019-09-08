package routers

import (
	"cash/routes/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/auth", api.GetAuth)
	r.GET("/", api.HomePage)
	return r
}
