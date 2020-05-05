package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/nicopellerin/virtual-canvas-api/graph/generated"
	"github.com/nicopellerin/virtual-canvas-api/graph/models"
)

func (r *publicProfileResolver) ID(ctx context.Context, obj *models.PublicProfile) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// PublicProfile returns generated.PublicProfileResolver implementation.
func (r *Resolver) PublicProfile() generated.PublicProfileResolver { return &publicProfileResolver{r} }

type publicProfileResolver struct{ *Resolver }
