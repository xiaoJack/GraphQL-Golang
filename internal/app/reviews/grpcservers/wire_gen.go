// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package grpcservers

import (
	"github.com/google/wire"
	"github.com/xiaoJack/GraphQL-Golang/internal/app/reviews/services"
	"github.com/xiaoJack/GraphQL-Golang/internal/pkg/config"
	"github.com/xiaoJack/GraphQL-Golang/internal/pkg/database"
	"github.com/xiaoJack/GraphQL-Golang/internal/pkg/log"
)

// Injectors from wire.go:

func CreateReviewsServer(cf string, service services.ReviewsService) (*ReviewsServer, error) {
	viper, err := config.New(cf)
	if err != nil {
		return nil, err
	}
	options, err := log.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	logger, err := log.New(options)
	if err != nil {
		return nil, err
	}
	reviewsServer, err := NewReviewsServer(logger, service)
	if err != nil {
		return nil, err
	}
	return reviewsServer, nil
}

// wire.go:

var testProviderSet = wire.NewSet(log.ProviderSet, config.ProviderSet, database.ProviderSet, ProviderSet)
