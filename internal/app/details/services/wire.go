//go:build wireinject
// +build wireinject

package services

import (
	"github.com/google/wire"
	"github.com/xiaoJack/GraphQL-Golang/internal/app/details/repositories"
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

func CreateDetailsService(cf string, sto repositories.DetailsRepository) (DetailsService, error) {
	panic(wire.Build(testProviderSet))
}
