package routes

import (
	"rs-bandung-api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/rumah-sakit", handlers.GetRumahSakit)
		api.GET("/rumah-sakit/near", handlers.GetRumahSakitNear)
	}
}
