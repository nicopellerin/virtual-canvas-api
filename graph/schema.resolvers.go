package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/nicopellerin/virtual-canvas-api/graph/database"
	"github.com/nicopellerin/virtual-canvas-api/graph/generated"
	"github.com/nicopellerin/virtual-canvas-api/graph/model"
	"github.com/nicopellerin/virtual-canvas-api/graph/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *imageResolver) Ratio(ctx context.Context, obj *models.Image) (*float64, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *imageResolver) Lighting(ctx context.Context, obj *models.Image) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UsernameInput) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) LoginUser(ctx context.Context, input model.LoginUserInput) (*models.User, error) {
	var user *models.User
	database.Collection.FindOne(ctx, bson.D{{"username", input.Username}}).Decode(&user)

	if user.Username != input.Username {
		fmt.Println("User does not exist")
		return nil, errors.New("User does not exist")
	}

	// match := user.CheckPasswordHash(input.Password)
	// if !match {
	// 	fmt.Println("Password is not valid")
	// 	return nil, errors.New("Password is not valid")
	// }

	// authToken, err := user.GenerateJWT()
	// if err != nil {
	// 	fmt.Println("Failed to generate token")
	// }
	return user, nil
	// return &model.AuthResponse{
	// 	AuthToken: authToken,
	// 	User:      user,
	// }, nil
}

func (r *mutationResolver) SignupUser(ctx context.Context, input model.SignupUserInput) (*model.AuthResponse, error) {
	// var userDB *models.User

	// database.Collection.FindOne(ctx, bson.D{{"username", input.Username}}).Decode(&userDB)

	// if userDB.Username != "" {
	// 	fmt.Println("User already exists")
	// 	return nil, errors.New("User already exists")
	// }

	// hash, err := userDB.HashPassword(input.Password)
	// if err != nil {
	// 	fmt.Println("Error hashing password")
	// 	return nil, err
	// }

	// userDB.Password = hash

	// if input.Username != "" {
	// 	insertResult, err := database.Collection.InsertOne(ctx, userDB)
	// 	if err != nil {
	// 		log.Fatal(err, insertResult)
	// 	}
	// }

	// validToken, err := userDB.GenerateJWT()
	// if err != nil {
	// 	fmt.Println("Failed to generate token")
	// }

	// return &model.AuthResponse{
	// 	AuthToken: validToken,
	// 	User:      userDB,
	// }, nil
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddArtwork(ctx context.Context, input model.AddArtworkInput) (*models.Image, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUser(ctx context.Context, input *model.UsernameInput) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) ID(ctx context.Context, obj *models.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Image returns generated.ImageResolver implementation.
func (r *Resolver) Image() generated.ImageResolver { return &imageResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type imageResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
