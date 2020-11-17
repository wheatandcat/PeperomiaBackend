package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/wheatandcat/PeperomiaBackend/domain"
)

// ItemDetailRepository is repository for itemDetail
type ItemDetailRepository struct {
}

// NewItemDetailRepository is Create new ItemDetailRepository
func NewItemDetailRepository() domain.ItemDetailRepository {
	return &ItemDetailRepository{}
}

// getItemDetailCollection アイテムのコレクションを取得する
func getItemDetailCollection(f *firestore.Client, key domain.ItemDetailKey) *firestore.DocumentRef {
	iDoc := getItemDocID(key.ItemDetailID)
	date := key.Date.Format("2006-01-02")

	return f.Collection("version/1/users/" + key.UID + "/calendars/" + date + "/items/" + key.ItemID + "/itemDetails").Doc(iDoc)
}

func getItemDetailDocID(itemDetailID string) string {
	doc := itemDetailID

	return doc
}

// Create アイテム詳細を作成する
func (re *ItemDetailRepository) Create(ctx context.Context, f *firestore.Client, i domain.ItemDetailRecord, key domain.ItemDetailKey) error {
	_, err := getItemDetailCollection(f, key).Set(ctx, i)

	return err
}

// Update アイテム詳細を更新する
func (re *ItemDetailRepository) Update(ctx context.Context, f *firestore.Client, i domain.ItemDetailRecord, key domain.ItemDetailKey) error {
	_, err := getItemDetailCollection(f, key).Set(ctx, i)

	return err
}

// Delete アイテム詳細を削除する
func (re *ItemDetailRepository) Delete(ctx context.Context, f *firestore.Client, key domain.ItemDetailKey) error {
	_, err := getItemDetailCollection(f, key).Delete(ctx)

	return err
}

// Get アイテム詳細を取得する
func (re *ItemDetailRepository) Get(ctx context.Context, f *firestore.Client, key domain.ItemDetailKey) (domain.ItemDetailRecord, error) {
	var idr domain.ItemDetailRecord

	snap, err := getItemDetailCollection(f, key).Get(ctx)
	if err != nil {
		return idr, err
	}

	snap.DataTo(&idr)
	return idr, nil
}

// FindByItemID ItemIDから取得する
func (re *ItemDetailRepository) FindByItemID(ctx context.Context, f *firestore.Client, itemID string) ([]domain.ItemDetailRecord, error) {
	var ids []domain.ItemDetailRecord
	matchItem := f.Collection("itemDetails").Where("itemId", "==", itemID).OrderBy("priority", firestore.Asc).Documents(ctx)
	docs, err := matchItem.GetAll()
	if err != nil {
		return ids, err
	}

	for _, doc := range docs {
		var id domain.ItemDetailRecord
		doc.DataTo(&id)
		ids = append(ids, id)
	}

	return ids, nil
}

// GetItemDetailDocs DocumentからItemDetailsDocを取得する
func GetItemDetailDocs(ctx context.Context, doc *firestore.DocumentSnapshot) ([]*firestore.DocumentSnapshot, error) {
	matchItems := doc.Ref.Collection("itemDetails").OrderBy("priority", firestore.Asc).Documents(ctx)
	docs, err := matchItems.GetAll()
	if err != nil {
		return nil, err
	}

	return docs, nil
}

// GetItemDetailsByDocument DocumentからItemDetailsを取得する
func GetItemDetailsByDocument(ctx context.Context, doc *firestore.DocumentSnapshot) ([]*domain.ItemDetailRecord, error) {
	docs, err := GetItemDetailDocs(ctx, doc)
	if err != nil {
		return nil, err
	}

	var items []*domain.ItemDetailRecord
	for _, did := range docs {
		var id *domain.ItemDetailRecord
		did.DataTo(&id)
		items = append(items, id)
	}

	return items, nil
}

// DeleteItemDetailsDoc DocumentからItemDetailsDocを削除する
func DeleteItemDetailsDoc(ctx context.Context, doc *firestore.DocumentSnapshot) error {
	iddocs, err := GetItemDetailDocs(ctx, doc)
	if err != nil {
		return err
	}

	for _, iddoc := range iddocs {
		_, err := iddoc.Ref.Delete(ctx)
		if err != nil {
			return err
		}

	}

	return nil
}
