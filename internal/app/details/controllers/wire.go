//go:build wireinject
// +build wireinject

package controllers

import (
	"github.com/google/wire"
	"github.com/xiaoJack/GraphQL-Golang/internal/app/details/repositories"
	"github.com/xiaoJack/GraphQL-Golang/internal/app/details/services"
	"github.com/xiaoJack/GraphQL-Golang/internal/pkg/config"
	"github.com/xiaoJack/GraphQL-Golang/internal/pkg/database"
	"github.com/xiaoJack/GraphQL-Golang/internal/pkg/log"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	services.ProviderSet,
	//repositories.ProviderSet,
	ProviderSet,
)

func CreateDetailsController(cf string, sto repositories.DetailsRepository) (*DetailsController, error) {
	panic(wire.Build(testProviderSet))
}
