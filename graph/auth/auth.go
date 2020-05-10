package auth

import (
	"context"
	"net/http"

	"github.com/nicopellerin/virtual-canvas-api/graph/database"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/nicopellerin/virtual-canvas-api/graph/models"
)

type contextKey struct {
	name string
}

var userCtxKey = &contextKey{"user"}

func Middleware(db *database.Collection) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Cookie("vc-auth")

			if err != nil || c == nil {
				next.ServeHTTP(w, r)
			}

			userID, err := validateAndGetUserID(c)
			if err != nil {
				http.Error(w, "Invalid cookie", http.StatusForbidden)
				return
			}

			user := getUserByID(db, userID)

			ctx := context.WithValue(r.Context(), userCtxKey, user)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func ForContext(ctx context.Context) *models.User {
	raw, _ := ctx.Value(userCtxKey).(*models.User)
	return raw
}

func validateAndGetUserID(c *http.Cookie) (string, error) {
	return c.String(), nil
}

func getUserByID(db *database.DB, userID string) *models.User {
	var user *models.User
	ctx := context.Background()
	db.Collection.FindOne(ctx, bson.D{{"username", userID}}).Decode(&user)
	return user
}
