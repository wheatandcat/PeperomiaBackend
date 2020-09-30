package repository

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"

	"github.com/wheatandcat/PeperomiaBackend/domain"
)

// ItemRepository is repository for item
type ItemRepository struct {
}

// NewItemRepository is Create new ItemRepository
func NewItemRepository() domain.ItemRepository {
	return &ItemRepository{}
}

// GetItemCollection アイテムのコレクションを取得する
func GetItemCollection(f *firestore.Client, uID string, itemID string) *firestore.DocumentRef {
	iDoc := getItemDocID(uID, itemID)

	return f.Collection("items").Doc(iDoc)
}

func getItemDocID(uID string, itemID string) string {
	doc := uID + "_" + itemID

	return doc
}

// Create アイテムを作成する
func (re *ItemRepository) Create(ctx context.Context, f *firestore.Client, i domain.ItemRecord) error {
	iDoc := getItemDocID(i.UID, i.ID)
	i.CreatedAt = time.Now()
	i.UpdatedAt = time.Now()

	_, err := f.Collection("items").Doc(iDoc).Set(ctx, i)

	return err
}

// Update アイテムを更新する
func (re *ItemRepository) Update(ctx context.Context, f *firestore.Client, i domain.ItemRecord) error {
	iDoc := getItemDocID(i.UID, i.ID)
	i.UpdatedAt = time.Now()

	_, err := f.Collection("items").Doc(iDoc).Set(ctx, i)

	return err
}

// Delete アイテムを削除する
func (re *ItemRepository) Delete(ctx context.Context, f *firestore.Client, i domain.ItemRecord) error {
	iDoc := getItemDocID(i.UID, i.ID)

	_, err := f.Collection("items").Doc(iDoc).Delete(ctx)

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
	iDoc := getItemDocID(uid, itemID)

	dsnap, err := f.Collection("items").Doc(iDoc).Get(ctx)
	if err != nil {
		return ir, err
	}

	dsnap.DataTo(&ir)
	return ir, nil
}

// FindByPublicAndID 公開中かつIDから取得する
func (re *ItemRepository) FindByPublicAndID(ctx context.Context, f *firestore.Client, id string) (domain.ItemRecord, error) {
	var item domain.ItemRecord
	matchItem := f.Collection("items").Where("id", "==", id).Where("public", "==", true).Limit(1).Documents(ctx)
	docs, err := matchItem.GetAll()
	if err != nil {
		return item, err
	}

	docs[0].DataTo(&item)

	return item, nil
}

// DeleteByUID ユーザーIDから削除する
func (re *ItemRepository) DeleteByUID(ctx context.Context, f *firestore.Client, uid string) error {
	matchItem := f.Collection("items").Where("uid", "==", uid).Documents(ctx)
	docs, err := matchItem.GetAll()
	if err != nil {
		return err
	}

	for _, doc := range docs {
		if _, err := doc.Ref.Delete(ctx); err != nil {
			return err
		}
	}

	return nil
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
