package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PublicProfile struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email    string             `json:"email, omitempty"`
	Username string             `json:"username, omitempty"`
	Images   []*Image           `json:"images, omitempty"`
	Social   Social             `json:"social, omitempty" bson:"social"`
}
