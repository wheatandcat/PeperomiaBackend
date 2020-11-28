package repository

import (
	"context"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/wheatandcat/PeperomiaBackend/domain"
)

// DictionaryRepository is repository for dictionary
type DictionaryRepository struct {
}

// NewDictionaryRepository is Create new CalendarRepository
func NewDictionaryRepository() domain.DictionaryRepository {
	return &DictionaryRepository{}
}

func ngram(targetText string, n int) ([]string, error) {
	sepText := strings.Split(targetText, "")
	var ngrams []string

	if len(sepText) < n {
		r := []string{}
		r = append(r, targetText)
		return r, nil
	}

	for i := 0; i < (len(sepText) - n + 1); i++ {
		ngrams = append(ngrams, strings.Join(sepText[i:i+n], ""))
	}
	return ngrams, nil
}

// Find テキストからサジェストを取得する
func (re *DictionaryRepository) Find(ctx context.Context, f *firestore.Client, text string) ([]string, error) {

	bigrams, err := ngram(text, 2)
	if err != nil {
		return nil, err
	}

	var query = f.Collection("version/1/dictionary").Limit(100)

	for _, bigram := range bigrams {
		key := "bigrams." + bigram
		query = query.Where(key, "==", true)
	}

	matchItem := query.Documents(ctx)
	docs, err := matchItem.GetAll()
	if err != nil {
		return nil, err
	}

	var textList []string

	for _, doc := range docs {
		text := doc.Data()["text"]

		switch v := text.(type) {
		case string:
			textList = append(textList, v)
		}

	}

	return textList, nil
}
