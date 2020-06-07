package domain

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
)

// UserRecord is user data
type UserRecord struct {
	UID       string    `json:"uid" firestore:"uid"`
	CreatedAt time.Time `json:"createdAt" firestore:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" firestore:"updatedAt"`
}

// UserRepository is repository interface
type UserRepository interface {
	Create(ctx context.Context, f *firestore.Client, u UserRecord) error
	FindByUID(ctx context.Context, f *firestore.Client, uid string) (UserRecord, error)
	ExistsByUID(ctx context.Context, f *firestore.Client, uid string) (bool, error)
}
