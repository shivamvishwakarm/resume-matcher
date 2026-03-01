package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Resume struct {
	ID         bson.ObjectID `bson:"_id,omitempty" json:"id"`
	UserEmail  string        `bson:"user_email" json:"user_email"`
	Filename   string        `bson:"filename" json:"filename"`
	Filepath   string        `bson:"filepath" json:"filepath"`
	QdrantID   string        `bson:"qdrant_id" json:"qdrant_id"`
	UploadedAt time.Time     `bson:"uploaded_at" json:"uploaded_at"`
}
