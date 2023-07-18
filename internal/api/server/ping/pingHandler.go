package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getHealth(context *gin.Context) {
	response := healthResponse{
		Response: "Working like a charm!",
	}

	context.JSON(http.StatusOK, response)
}
