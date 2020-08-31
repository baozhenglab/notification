package gochannel

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Notification mapping table notifications
type Notification struct {
	ID             primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Type           string             `json:"type" bson:"type"`
	NotifiableType string             `json:"notifiable_type" bson:"notifiable_type"`
	NotifiableID   uint               `json:"notifiable_id" bson:"notifiable_id"`
	Data           interface{}        `json:"data" bson:"data"`
	ReadAt         *time.Time         `json:"read_at" bson:"read_at"`
	CreatedAt      time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at" bson:"updated_at"`
}
