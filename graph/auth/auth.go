package auth

import (
	"context"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/dgrijalva/jwt-go"
	"github.com/nicopellerin/virtual-canvas-api/graph/models"
	"github.com/nicopellerin/virtual-canvas-api/graph/utils"
)

type contextKey struct {
	name string
}

var userCtxKey = &contextKey{"user"}

func Middleware(db *mongo.Database) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var secret = utils.GetEnvVars("JWT_SECRET")

			token, err := jwt.Parse(r.Header.Get("Token"), func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				return []byte(secret), nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				username := claims["client"]

				user := getUserByID(db, username)

				ctx := context.WithValue(r.Context(), userCtxKey, user)

				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
			} else {
				fmt.Println(err)
				next.ServeHTTP(w, r)
			}
		})
	}
}

func ForContext(ctx context.Context) *models.User {
	raw, _ := ctx.Value(userCtxKey).(*models.User)
	return raw
}

func getUserByID(db *mongo.Database, userID interface{}) *models.User {
	var user *models.User
	ctx := context.Background()
	db.Collection("users").FindOne(ctx, bson.D{{"username", userID}}).Decode(&user)
	return user
}
