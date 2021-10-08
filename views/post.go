package views

import (
	"time"
)

type Post struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	Content   string    `json:"content"`
	Title     string    `json:"title"`
	CreatedAt time.Time `bson:"created_At" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at,omitempty"`
	Author    string    `json:"author"`
}
