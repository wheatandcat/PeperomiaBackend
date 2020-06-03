package domain

import (
	"context"

	"cloud.google.com/go/firestore"
)

// PushTokenRecord is push token data
type PushTokenRecord struct {
	ID       string `json:"id" firestore:"id" binding:"required"`
	UID      string `json:"uid" firestore:"uid"`
	Token    string `json:"token" firestore:"token" binding:"required"`
	DeviceID string `json:"deviceId" firestore:"deviceId" binding:"required"`
}

// PushTokenRepository is repository interface
type PushTokenRepository interface {
	Create(ctx context.Context, f *firestore.Client, i PushTokenRecord) error
}
