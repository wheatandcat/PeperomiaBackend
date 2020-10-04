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

func getPushTokenDocID(uID string, deviceID string) string {
	doc := uID + "_" + deviceID
	return doc
}

func pushTokenCollectionRef(f *firestore.Client, uid string) *firestore.CollectionRef {
	return f.Collection("version/1/users/" + uid + "/expoPushTokens")
}

func pushTokenCollection(f *firestore.Client, uid string, deviceID string) *firestore.DocumentRef {
	idDoc := getPushTokenDocID(uid, deviceID)

	return pushTokenCollectionRef(f, uid).Doc(idDoc)
}

// Create カレンダーを作成する
func (re *PushTokenRepository) Create(ctx context.Context, f *firestore.Client, p domain.PushTokenRecord) error {
	_, err := pushTokenCollection(f, p.UID, p.DeviceID).Set(ctx, p)

	return err
}

// FindByUID ユーザーIDから取得する
func (re *PushTokenRepository) FindByUID(ctx context.Context, f *firestore.Client, uid string) ([]domain.PushTokenRecord, error) {
	var items []domain.PushTokenRecord
	matchItem := pushTokenCollectionRef(f, uid).Documents(ctx)
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
	matchItem := f.CollectionGroup("expoPushTokens").Documents(ctx)
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

// GetPushTokenDocs DocumentからPushTokenDocsを取得する
func GetPushTokenDocs(ctx context.Context, doc *firestore.DocumentSnapshot) ([]*firestore.DocumentSnapshot, error) {
	matchItems := doc.Ref.Collection("expoPushTokens").Documents(ctx)
	docs, err := matchItems.GetAll()
	if err != nil {
		return nil, err
	}

	return docs, nil
}

// GetPushTokensByDocument DocumentからPushTokensを取得する
func GetPushTokensByDocument(ctx context.Context, doc *firestore.DocumentSnapshot) ([]*domain.PushTokenRecord, error) {
	docs, err := GetPushTokenDocs(ctx, doc)
	if err != nil {
		return nil, err
	}

	var items []*domain.PushTokenRecord
	for _, doc := range docs {
		var r *domain.PushTokenRecord
		doc.DataTo(&r)
		items = append(items, r)
	}

	return items, nil
}
