package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
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
	Images   *[]*Image          `json:"images, omitempty" bson:"images"`
	Social   Social             `json:"social, omitempty" bson:"social"`
}

type UpdateUserInput struct {
	Username  string  `json:"username"`
	Instagram *string `json:"instagram"`
	Facebook  *string `json:"facebook"`
	Website   *string `json:"website"`
}

type UsernameInput struct {
	Username string `json:"username"`
}

type LoginUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignupUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type AuthResponse struct {
	AuthToken *AuthToken `json:"authToken"`
	User      *User      `json:"user"`
}

type AuthToken struct {
	AccessToken string    `json:"accessToken"`
	ExpiredAt   time.Time `json:"expiredAt"`
}

// HashPassword - Hashes password before saving to DB
func (u *User) HashPassword(password string) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return errors.New("YOOOOOO")
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

func (u *User) GenerateJWT() (*AuthToken, error) {
	expiredAt := time.Now().Add(time.Hour * 24 * 7)

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = u.Username
	claims["exp"] = expiredAt

	accessToken, err := token.SignedString([]byte(secret))

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &AuthToken{
		AccessToken: accessToken,
		ExpiredAt:   expiredAt,
	}, nil
}
