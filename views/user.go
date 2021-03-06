package views

import (
	"time"
)

type User struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `bson:"created_At" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at,omitempty"`
}
