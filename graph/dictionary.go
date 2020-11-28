package graph

import (
	"context"
)

// GetSuggestionText サジェストリストを取得する
func (g *Graph) GetSuggestionText(ctx context.Context, text string) ([]string, error) {
	h := g.Handler

	list, err := h.App.DictionaryRepository.Find(ctx, h.FirestoreClient, text)
	if err != nil {
		return nil, err
	}

	return list, nil
}
