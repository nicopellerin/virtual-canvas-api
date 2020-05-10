package database

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/nicopellerin/virtual-canvas-api/graph/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersCollection struct {
	DB *mongo.Database
}

func (u *UsersCollection) LoginUser(ctx context.Context, input models.LoginUserInput) (*models.AuthResponse, error) {
	var user *models.User
	u.DB.Collection("users").FindOne(ctx, bson.D{{"username", input.Username}}).Decode(&user)

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

func (u *UsersCollection) SignupUser(ctx context.Context, input models.SignupUserInput) (*models.AuthResponse, error) {
	user := new(models.User)

	err := u.DB.Collection("users").FindOne(ctx, bson.D{{"username", input.Username}}).Decode(&user)
	if err == nil {
		return nil, errors.New("User already exists")
	}

	if len(input.Password) < 8 {
		return nil, errors.New("Please choose a password with a minimum of 8 characters")
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
		insertResult, err := u.DB.Collection("users").InsertOne(ctx, user)
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

func (u *UsersCollection) GetUser(ctx context.Context, input *models.UsernameInput) (*models.User, error) {
	var user *models.User

	filter := bson.M{"username": input.Username}

	err := u.DB.Collection("users").FindOne(ctx, filter).Decode(&user)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return user, nil
}

func (u *UsersCollection) UpdateUser(ctx context.Context, input models.UpdateUserInput) (*models.User, error) {
	user := &models.User{Username: "nicopellerin"}

	filter := bson.M{"username": input.Username}

	update := bson.M{"$set": bson.M{"social.instagram": input.Instagram, "social.facebook": input.Facebook, "social.website": input.Website}}

	res, err := u.DB.Collection("users").UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(res, err)
		return nil, err
	}
	return user, nil
}

func (u *UsersCollection) GetPublicProfile(ctx context.Context, input *models.UsernameInput) (*models.PublicProfile, error) {
	filter := bson.M{"username": input.Username}

	var user models.PublicProfile

	err := u.DB.Collection("users").FindOne(ctx, filter).Decode(&user)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}

func (u *UsersCollection) AddArtwork(ctx context.Context, input models.AddArtworkInput) (*models.Image, error) {
	filter := bson.M{"username": input.Username}
	update := bson.M{"$push": bson.M{"images": &models.Image{Background: input.Background, Border: input.Border, ID: input.ImageID, Name: input.Name, Ratio: input.Ratio, Rotate: input.Rotate, Src: input.Src, Texture: input.Texture}}}

	res, err := u.DB.Collection("users").UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err, res)
		return nil, err
	}

	return &models.Image{}, nil
}

func (u *UsersCollection) UpdateArtwork(ctx context.Context, input models.UpdateArtworkInput) (*models.Image, error) {
	filter := bson.M{"username": input.Username, "images.id": input.ImageID}

	update := bson.M{"$set": bson.M{"images.$.name": input.Name, "images.$.border": input.Border, "images.$.texture": input.Texture, "images.$.background": input.Background, "images.$.rotate": input.Rotate, "images.$.lighting": input.Lighting}}

	res, err := u.DB.Collection("users").UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println(res, err)
		return nil, err
	}
	return &models.Image{}, nil
}

func (u *UsersCollection) DeleteArtwork(ctx context.Context, input *models.DeleteArtworkInput) (*models.Image, error) {
	filter := bson.M{"username": input.Username, "images.id": input.ID}
	update := bson.M{"$pull": bson.M{"images": bson.M{"id": input.ID}}}

	res, err := u.DB.Collection("users").UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(res, err)
		return nil, err
	}

	return &models.Image{}, nil
}
