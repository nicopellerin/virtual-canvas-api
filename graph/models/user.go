package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Users
type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email    string             `json:"email, omitempty"`
	Password string             `json:"password, omitempty"`
	Username string             `json:"username, omitempty"`
	Images   *[]*Image          `json:"images, omitempty"`
	Social   Social             `json:"social_links, omitempty" bson:"social_links"`
}
