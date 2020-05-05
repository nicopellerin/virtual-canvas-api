package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/nicopellerin/virtual-canvas-api/graph/generated"
	"github.com/nicopellerin/virtual-canvas-api/graph/models"
)

func (r *imageResolver) Ratio(ctx context.Context, obj *models.Image) (*float64, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *imageResolver) Lighting(ctx context.Context, obj *models.Image) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

// Image returns generated.ImageResolver implementation.
func (r *Resolver) Image() generated.ImageResolver { return &imageResolver{r} }

type imageResolver struct{ *Resolver }
