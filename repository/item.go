package repository

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"

	"github.com/wheatandcat/PeperomiaBackend/domain"
)

// ItemRepository is repository for item
type ItemRepository struct{}

// NewItemRepository is Create new ItemRepository
func NewItemRepository() domain.ItemRepository {
	return &ItemRepository{}
}

// GetItemCollection アイテムのコレクションを取得する
func getItemCollection(f *firestore.Client, itemID string, key domain.ItemKey) *firestore.DocumentRef {
	iDoc := getItemDocID(itemID)
	date := key.Date.Format("2006-01-02")

	return f.Collection("version/1/" + key.UID + "/calendars/" + date + "/items").Doc(iDoc)
}

func getItemDocID(itemID string) string {
	doc := itemID

	return doc
}

// Create アイテムを作成する
func (re *ItemRepository) Create(ctx context.Context, f *firestore.Client, i domain.ItemRecord, key domain.ItemKey) error {
	i.CreatedAt = time.Now()
	i.UpdatedAt = time.Now()
	_, err := getItemCollection(f, i.ID, key).Set(ctx, i)

	return err
}

// Update アイテムを更新する
func (re *ItemRepository) Update(ctx context.Context, f *firestore.Client, i domain.ItemRecord, key domain.ItemKey) error {
	i.UpdatedAt = time.Now()
	_, err := getItemCollection(f, i.ID, key).Set(ctx, i)

	return err
}

// FindByUID ユーザーIDから取得する
func (re *ItemRepository) FindByUID(ctx context.Context, f *firestore.Client, uid string) ([]domain.ItemRecord, error) {
	var items []domain.ItemRecord
	matchItem := f.Collection("items").Where("uid", "==", uid).Documents(ctx)
	docs, err := matchItem.GetAll()
	if err != nil {
		return items, err
	}

	for _, doc := range docs {
		var item domain.ItemRecord
		doc.DataTo(&item)
		items = append(items, item)
	}

	return items, nil
}

// FindByDoc ドキュメントから取得する
func (re *ItemRepository) FindByDoc(ctx context.Context, f *firestore.Client, uid string, itemID string) (domain.ItemRecord, error) {
	var ir domain.ItemRecord
	iDoc := getItemDocID(itemID)

	dsnap, err := f.Collection("items").Doc(iDoc).Get(ctx)
	if err != nil {
		return ir, err
	}

	dsnap.DataTo(&ir)
	return ir, nil
}

// GetItemDoc DocumentからItemDocを取得する
func GetItemDoc(ctx context.Context, doc *firestore.DocumentSnapshot) (*firestore.DocumentSnapshot, error) {
	matchItems := doc.Ref.Collection("items").Documents(ctx)
	docs, err := matchItems.GetAll()
	if err != nil {
		return nil, err
	}

	return docs[0], nil
}

// DeleteItemDoc DocumentからItemDocを削除する
func DeleteItemDoc(ctx context.Context, doc *firestore.DocumentSnapshot) error {
	idoc, err := GetItemDoc(ctx, doc)
	if err != nil {
		return err
	}

	err = DeleteItemDetailsDoc(ctx, idoc)
	if err != nil {
		return err
	}

	_, err = idoc.Ref.Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}
