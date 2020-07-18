package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/wheatandcat/PeperomiaBackend/domain"
)

// PushTokenRepository is repository for push token
type PushTokenRepository struct {
}

// NewPushTokenRepository is Create new CalendarRepository
func NewPushTokenRepository() domain.PushTokenRepository {
	return &PushTokenRepository{}
}

func getPushTokenDocID(uID string, DeviceID string, pushTokenID string) string {
	doc := uID + "_" + DeviceID + "_" + pushTokenID
	return doc
}

// Create カレンダーを作成する
func (re *PushTokenRepository) Create(ctx context.Context, f *firestore.Client, p domain.PushTokenRecord) error {
	idDoc := getPushTokenDocID(p.UID, p.DeviceID, p.Token)

	_, err := f.Collection("expoPushTokens").Doc(idDoc).Set(ctx, p)

	return err
}

// FindByUID ユーザーIDから取得する
func (re *PushTokenRepository) FindByUID(ctx context.Context, f *firestore.Client, uid string) ([]domain.PushTokenRecord, error) {
	var items []domain.PushTokenRecord
	matchItem := f.Collection("expoPushTokens").Where("uid", "==", uid).Documents(ctx)
	docs, err := matchItem.GetAll()
	if err != nil {
		return items, err
	}

	for _, doc := range docs {
		var item domain.PushTokenRecord
		doc.DataTo(&item)
		items = append(items, item)
	}

	return items, nil
}

// FindAll 全て取得する
func (re *PushTokenRepository) FindAll(ctx context.Context, f *firestore.Client) ([]domain.PushTokenRecord, error) {
	var items []domain.PushTokenRecord
	matchItem := f.Collection("expoPushTokens").Documents(ctx)
	docs, err := matchItem.GetAll()
	if err != nil {
		return items, err
	}

	for _, doc := range docs {
		var item domain.PushTokenRecord
		doc.DataTo(&item)
		items = append(items, item)
	}

	return items, nil
}
