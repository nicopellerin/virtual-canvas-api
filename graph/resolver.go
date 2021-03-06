//go:generate go run github.com/99designs/gqlgen

package graph

import (
	"github.com/nicopellerin/virtual-canvas-api/graph/database"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UsersCollection database.UsersCollection
}
