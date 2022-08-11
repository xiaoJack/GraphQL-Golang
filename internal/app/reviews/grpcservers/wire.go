//go:build wireinject
// +build wireinject

package grpcservers

import (
	"github.com/google/wire"
	"github.com/xiaoJack/GraphQL-Golang/internal/app/reviews/services"
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

func CreateReviewsServer(cf string, service services.ReviewsService) (*ReviewsServer, error) {
	panic(wire.Build(testProviderSet))
}
