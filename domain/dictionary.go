package domain

import (
	"context"

	"cloud.google.com/go/firestore"
)

// DictionaryRecord is item data
type DictionaryRecord struct {
	Text    string   `json:"text" firestore:"text"`
	Bigrams []string `json:"bigrams" firestore:"bigrams"`
}

// DictionaryRepository is repository interface
type DictionaryRepository interface {
	Find(ctx context.Context, f *firestore.Client, text string) ([]string, error)
}
