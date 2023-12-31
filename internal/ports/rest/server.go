package rest

import (
	"fio-service/internal/app"
	"github.com/gin-gonic/gin"
	"net/http"
)

// NewRESTServer creates rest server with handlers
func NewRESTServer(addr string, a app.App) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	api := router.Group("api/v1")
	appRouter(api, a)
	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}
