package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserEntity struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	Email              string             `bson:"email,omitempty"`
	Password           string             `bson:"password,omitempty"`
	Name               string             `bson:"name,omitempty"`
	Age                int8               `bson:"age,omitempty"`
	PasswordExpiration time.Time          `bson:"password_expiration,omitempty"`
}
