//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/xiaoJack/GraphQL-Golang/internal/app/graphQL"
	"github.com/xiaoJack/GraphQL-Golang/internal/app/graphQL/controllers"
	"github.com/xiaoJack/GraphQL-Golang/internal/app/graphQL/grpcclients"
	"github.com/xiaoJack/GraphQL-Golang/internal/app/graphQL/services"
	"github.com/xiaoJack/GraphQL-Golang/internal/pkg/app"
	"github.com/xiaoJack/GraphQL-Golang/internal/pkg/config"
	"github.com/xiaoJack/GraphQL-Golang/internal/pkg/consul"
	"github.com/xiaoJack/GraphQL-Golang/internal/pkg/database"
	"github.com/xiaoJack/GraphQL-Golang/internal/pkg/jaeger"
	"github.com/xiaoJack/GraphQL-Golang/internal/pkg/log"
	"github.com/xiaoJack/GraphQL-Golang/internal/pkg/transports/grpc"
	"github.com/xiaoJack/GraphQL-Golang/internal/pkg/transports/http"
)

var providerSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	services.ProviderSet,
	consul.ProviderSet,
	jaeger.ProviderSet,
	http.ProviderSet,
	grpc.ProviderSet,
	details.ProviderSet,
	controllers.ProviderSet,
	grpcclients.ProviderSet,
)

func CreateApp(cf string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
