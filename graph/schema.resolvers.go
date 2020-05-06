package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/nicopellerin/virtual-canvas-api/graph/database"
	"github.com/nicopellerin/virtual-canvas-api/graph/generated"
	"github.com/nicopellerin/virtual-canvas-api/graph/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *mutationResolver) UpdateUser(ctx context.Context, input models.UsernameInput) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) LoginUser(ctx context.Context, input models.LoginUserInput) (*models.AuthResponse, error) {
	var user *models.User
	database.Collection.FindOne(ctx, bson.D{{"username", input.Username}}).Decode(&user)

	if user.Username != input.Username {
		fmt.Println("User does not exist")
		return nil, errors.New("User does not exist")
	}

	match := user.CheckPasswordHash(input.Password)
	if !match {
		fmt.Println("Password is not valid")
		return nil, errors.New("Password is not valid")
	}

	authToken, err := user.GenerateJWT()
	if err != nil {
		fmt.Println("Failed to generate token")
	}

	return &models.AuthResponse{
		AuthToken: authToken,
		User:      user,
	}, nil
}

func (r *mutationResolver) SignupUser(ctx context.Context, input models.SignupUserInput) (*models.AuthResponse, error) {
	user := new(models.User)

	err := database.Collection.FindOne(ctx, bson.D{{"username", input.Username}}).Decode(&user)
	if err == nil {
		return nil, errors.New("User already exists")
	}

	if len(input.Password) < 8 {
		return nil, errors.New("Please choose a password of minimum 8 characters long")
	}

	err = user.HashPassword(input.Password)
	if err != nil {
		fmt.Println("Error hashing password")
		return nil, err
	}

	user.Email = input.Email
	user.Username = input.Username
	user.Images = &[]*models.Image{}

	if input.Username != "" {
		insertResult, err := database.Collection.InsertOne(ctx, user)
		if err != nil {
			log.Fatal(err, insertResult)
		}
	}

	authToken, err := user.GenerateJWT()
	if err != nil {
		fmt.Println("Failed to generate token")
	}

	return &models.AuthResponse{
		AuthToken: authToken,
		User:      user,
	}, nil
}

func (r *mutationResolver) AddArtwork(ctx context.Context, input models.AddArtworkInput) (*models.Image, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *publicProfileResolver) ID(ctx context.Context, obj *models.PublicProfile) (string, error) {
	return obj.ID.Hex(), nil
}

func (r *publicProfileResolver) Images(ctx context.Context, obj *models.PublicProfile) ([]*models.Image, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUser(ctx context.Context, input *models.UsernameInput) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
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
