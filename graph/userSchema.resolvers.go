package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/nicopellerin/virtual-canvas-api/graph/generated"
	"github.com/nicopellerin/virtual-canvas-api/graph/models"
)

func (r *userResolver) ID(ctx context.Context, obj *models.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Images(ctx context.Context, obj *models.User) ([]*models.Image, error) {
	panic(fmt.Errorf("not implemented"))
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
