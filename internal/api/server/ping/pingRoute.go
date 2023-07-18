package ping

import (
	"github.com/gin-gonic/gin"
)

func SetServerHealthRoute(router *gin.RouterGroup) {

	router.GET("/ping", getHealth)
}
