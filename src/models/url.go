package models

import (
	"time"
)

type URL struct {
	Original  string    `bson:"original" json:"original"`
	Short     string    `bson:"short" json:"short"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}
