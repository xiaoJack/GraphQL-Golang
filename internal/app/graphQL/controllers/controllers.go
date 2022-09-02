package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/xiaoJack/GraphQL-Golang/internal/pkg/transports/http"
)

func CreateInitControllersFn(
	resolver *Resolver,
) http.InitControllers {
	return func(r *gin.Engine) {
		//r.Use(resolver.AuthMiddleware) //Auth use must frist
		r.POST("/query", resolver.Handler())
		r.GET("/", resolver.GroundHandler())

	}
}

var ProviderSet = wire.NewSet(NewResolver, CreateInitControllersFn)
