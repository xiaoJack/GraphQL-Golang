//go:build wireinject
// +build wireinject

package grpcservers

import (
	"github.com/google/wire"
	"github.com/xiaoJack/GraphQL-Golang/internal/app/details/services"
	"github.com/xiaoJack/GraphQL-Golang/internal/pkg/config"
	"github.com/xiaoJack/GraphQL-Golang/internal/pkg/database"
	"github.com/xiaoJack/GraphQL-Golang/internal/pkg/log"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	ProviderSet,
)

func CreateDetailsServer(cf string, service services.DetailsService) (*DetailsServer, error) {
	panic(wire.Build(testProviderSet))
}
