package views

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Post struct {
	ID 			primitive.ObjectID		`bson:"_id,omitempty" json:"id,omitempty"`
	Content		string					`json:"content"`
	Title		string					`json:"title"`
	CreatedAt	time.Time				`bson:"created_At" json:"created_at"`
	UpdatedAt	time.Time				`bson:"updated_at" json:"updated_at,omitempty"`
}

//Posts lista de posts
type Posts []*Post