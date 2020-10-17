package domain

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/wheatandcat/PeperomiaBackend/graph/model"
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
	FindByUID(ctx context.Context, f *firestore.Client, uid string) ([]PushTokenRecord, error)
	FindAll(ctx context.Context, f *firestore.Client) ([]PushTokenRecord, error)
}

// ToModel Modelに変換する
func (r *PushTokenRecord) ToModel() *model.ExpoPushToken {

	item := &model.ExpoPushToken{
		ID:       r.ID,
		UID:      r.UID,
		Token:    r.Token,
		DeviceID: r.DeviceID,
	}

	return item
}
