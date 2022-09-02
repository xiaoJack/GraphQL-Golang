package controllers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/xiaoJack/GraphQL-Golang/internal/app/graphQL/generated"
	"github.com/xiaoJack/GraphQL-Golang/internal/app/graphQL/services"
	"go.uber.org/zap"
	"net/http"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	logger  *zap.Logger
	service services.GraphService
}

func NewResolver(logger *zap.Logger, s services.GraphService) *Resolver {
	return &Resolver{
		logger:  logger,
		service: s,
	}
}

func (s *Resolver) Handler() gin.HandlerFunc {

	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: s}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func (s *Resolver) GroundHandler() gin.HandlerFunc {

	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func (s *Resolver) AuthMiddleware(c *gin.Context) {

	auth := c.Request.Header.Get("Authorization")
	if auth != "testAuth" {
		resp := map[string]string{"error": "bad auth"}

		c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
	}

	c.Next()
}
