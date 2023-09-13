package rest

import (
	"fio-service/internal/app"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRESTServer(host string, port int, a app.App) *http.Server {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	api := router.Group("api/v1")
	appRouter(api, a)
	return &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: router,
	}
}
