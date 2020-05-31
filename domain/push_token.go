package domain

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
)

// PushTokenRecord is push token data
type PushTokenRecord struct {
	ID        string    `json:"id" firestore:"id" binding:"required"`
	UID       string    `json:"uid" firestore:"uid"`
	Token     string    `json:"token" firestore:"token" binding:"required"`
	CreatedAt time.Time `json:"-" firestore:"createdAt"`
	UpdatedAt time.Time `json:"-" firestore:"updatedAt"`
}

// PushTokenRepository is repository interface
type PushTokenRepository interface {
	Create(ctx context.Context, f *firestore.Client, p PushTokenRecord) error
}
