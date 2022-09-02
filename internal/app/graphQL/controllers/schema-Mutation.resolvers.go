package controllers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/xiaoJack/GraphQL-Golang/internal/app/graphQL/generated"
	"github.com/xiaoJack/GraphQL-Golang/internal/pkg/models"
)

// CreateDetail is the resolver for the createDetail field.
func (r *mutationResolver) CreateDetail(ctx context.Context, input models.Detail) (*models.Detail, error) {

	data, err := r.service.DetailAdd(ctx, input.Name, float32(input.Price))
	if err != nil {
		return nil, err
	}

	return &models.Detail{ID: data.Id, Name: data.Name, Price: float64(data.Price), CreatedTime: data.CreatedTime.AsTime()}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
