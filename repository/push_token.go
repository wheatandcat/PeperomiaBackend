package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/wheatandcat/PeperomiaBackend/backend/domain"
)

// PushTokenRepository is repository for push token
type PushTokenRepository struct {
}

// NewPushTokenRepository is Create new CalendarRepository
func NewPushTokenRepository() domain.PushTokenRepository {
	return &PushTokenRepository{}
}

func getPushTokenDocID(uID string, pushTokenID string) string {
	doc := uID + "_" + pushTokenID
	return doc
}

// Create カレンダーを作成する
func (re *PushTokenRepository) Create(ctx context.Context, f *firestore.Client, p domain.PushTokenRecord) error {
	idDoc := getPushTokenDocID(p.UID, p.ID)

	_, err := f.Collection("expoPushToken").Doc(idDoc).Set(ctx, p)

	return err
}
