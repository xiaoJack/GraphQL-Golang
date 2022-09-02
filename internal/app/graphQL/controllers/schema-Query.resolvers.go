package controllers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/xiaoJack/GraphQL-Golang/internal/app/graphQL/generated"
	"github.com/xiaoJack/GraphQL-Golang/internal/pkg/models"
)

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context, id int) (*models.Detail, error) {
	data, err := r.service.DetailGet(ctx, int64(id))
	if err != nil {
		return nil, err
	}

	return &models.Detail{ID: data.Id, Name: data.Name, Price: float64(data.Price), CreatedTime: data.CreatedTime.AsTime()}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
