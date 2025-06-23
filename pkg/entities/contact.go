package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Contact struct {
	ID        primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Nama      string             `json:"nama" bson:"nama"`
	NoHp      string             `json:"nohp" bson:"nohp,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type DeleteRequest struct {
	ID string `json:"id"`
}
