package models

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/nicopellerin/virtual-canvas-api/graph/model"
	"github.com/nicopellerin/virtual-canvas-api/graph/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// Users
type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email    string             `json:"email, omitempty"`
	Password string             `json:"password, omitempty"`
	Username string             `json:"username, omitempty"`
	Images   []Image            `json:"images, omitempty" bson:"images"`
	Social   Social             `json:"social, omitempty" bson:"social"`
}

// HashPassword - Hashes password before saving to DB
func (u *User) HashPassword(password string) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	u.Password = string(passwordHash)

	return nil
}

// CheckPasswordHash - Checks to see if password entered matches password hash
func (u *User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

var secret = utils.GetEnvVars("JWT_SECRET")

func (u *User) GenerateJWT() (*model.AuthToken, error) {
	expiredAt := time.Now().Add(time.Hour * 24 * 7)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: expiredAt.Unix(),
		Id:        u.Username,
		IssuedAt:  time.Now().Unix(),
		Issuer:    "virtual-canvas",
	})

	accessToken, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return nil, nil
	// return &model.AuthToken{
	// 	AccessToken: accessToken,
	// 	ExpiredAt:   expiredAt,
	// }, nil
}
