package rest

import (
	"fio-service/internal/app"
	"github.com/gin-gonic/gin"
)

// appRouter adds handlers to server
func appRouter(r *gin.RouterGroup, a app.App) {
	r.POST("/fio", addFio(a))
	r.GET("/fio/:id", getFioById(a))
	r.GET("/fio", getFioByFilter(a))
	r.PUT("/fio/:id", updateFio(a))
	r.DELETE("/fio/:id", deleteFio(a))
}
