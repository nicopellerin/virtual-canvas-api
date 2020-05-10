package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/nicopellerin/virtual-canvas-api/graph/generated"
	"github.com/nicopellerin/virtual-canvas-api/graph/models"
)

func (r *mutationResolver) UpdateUser(ctx context.Context, input models.UpdateUserInput) (*models.User, error) {
	return r.UsersCollection.UpdateUser(ctx, input)
}

func (r *mutationResolver) LoginUser(ctx context.Context, input models.LoginUserInput) (*models.AuthResponse, error) {
	return r.UsersCollection.LoginUser(ctx, input)
}

func (r *mutationResolver) SignupUser(ctx context.Context, input models.SignupUserInput) (*models.AuthResponse, error) {
	return r.UsersCollection.SignupUser(ctx, input)
}

func (r *mutationResolver) AddArtwork(ctx context.Context, input models.AddArtworkInput) (*models.Image, error) {
	return nil, nil
}

func (r *mutationResolver) UpdateArtwork(ctx context.Context, input *models.UpdateArtworkInput) (*models.Image, error) {
	return r.UsersCollection.UpdateArtwork(ctx, input)
}

func (r *mutationResolver) DeleteArtwork(ctx context.Context, input *models.DeleteArtworkInput) (*models.Image, error) {
	return r.UsersCollection.DeleteArtwork(ctx, input)
}

func (r *publicProfileResolver) ID(ctx context.Context, obj *models.PublicProfile) (string, error) {
	return obj.ID.Hex(), nil
}

func (r *queryResolver) GetUser(ctx context.Context, input *models.UsernameInput) (*models.User, error) {
	return r.UsersCollection.GetUser(ctx, input)
}

func (r *queryResolver) GetPublicProfile(ctx context.Context, input *models.UsernameInput) (*models.PublicProfile, error) {
	return r.UsersCollection.GetPublicProfile(ctx, input)
}

func (r *userResolver) ID(ctx context.Context, obj *models.User) (string, error) {
	return obj.ID.Hex(), nil
}

func (r *userResolver) Images(ctx context.Context, obj *models.User) ([]*models.Image, error) {
	return *obj.Images, nil
}

func (r *userResolver) Social(ctx context.Context, obj *models.User) (*models.Social, error) {
	return &obj.Social, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// PublicProfile returns generated.PublicProfileResolver implementation.
func (r *Resolver) PublicProfile() generated.PublicProfileResolver { return &publicProfileResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type publicProfileResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
