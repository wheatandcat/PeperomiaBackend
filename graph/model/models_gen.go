// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Calendar struct {
	ID string `json:"id"`
	// 日付
	Date string `json:"date"`
	// true: パブリック、false: プライベート
	Public bool `json:"public"`
	// スケジュール
	Item *Item `json:"item"`
}

type ExpoPushToken struct {
	ID  string `json:"id"`
	UID string `json:"uid"`
	// デバイスID
	DeviceID string `json:"deviceId"`
	// トークン
	Token string `json:"token"`
}

type Item struct {
	ID string `json:"id"`
	// タイトル
	Title string `json:"title"`
	// 種別
	Kind string `json:"kind"`
	// スケジュール詳細
	ItemDetails []*ItemDetail `json:"itemDetails"`
}

type ItemDetail struct {
	ID string `json:"id"`
	// タイトル
	Title string `json:"title"`
	// 種類
	Kind  string `json:"kind"`
	Place string `json:"place"`
	// URL
	URL string `json:"url"`
	// メモ
	Memo string `json:"memo"`
	// 表示順
	Priority int `json:"priority"`
}

type NewCalendar struct {
	// 日付
	Date string `json:"date"`
	// スケジュール
	Item *NewItem `json:"item"`
}

type NewItem struct {
	// タイトル
	Title string `json:"title"`
	// 種類
	Kind  string `json:"kind"`
	Place string `json:"place"`
	URL   string `json:"url"`
	Memo  string `json:"memo"`
}

type NewItemDetail struct {
	// 日付
	Date   string `json:"date"`
	ItemID string `json:"itemId"`
	// タイトル
	Title string `json:"title"`
	// 種類
	Kind     string `json:"kind"`
	Place    string `json:"place"`
	URL      string `json:"url"`
	Memo     string `json:"memo"`
	Priority int    `json:"priority"`
}

type ShareItem struct {
	ID     string `json:"id"`
	ItemID string `json:"itemId"`
	// 日付
	Date string `json:"date"`
	// スケジュール
	Item *Item `json:"item"`
}

type User struct {
	ID string `json:"id"`
	// ユーザーID
	UID string `json:"uid"`
	// 役割:(管理権限: admin)
	Role int `json:"role"`
	// PUSH通知設定
	ExpoPushTokens []*ExpoPushToken `json:"expoPushTokens"`
}
