//go:build wireinject
// +build wireinject

package repositories

import (
	"github.com/google/wire"
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

func CreateRatingRepository(f string) (RatingsRepository, error) {
	panic(wire.Build(testProviderSet))
}
