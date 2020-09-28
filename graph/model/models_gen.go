// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Calendar struct {
	ID     string `json:"id"`
	ItemID string `json:"itemId"`
	Date   string `json:"date"`
	Public bool   `json:"public"`
	Item   *Item  `json:"item"`
}

type ExpoPushToken struct {
	ID       string `json:"id"`
	UID      string `json:"uid"`
	DeviceID string `json:"deviceId"`
	Token    string `json:"token"`
}

type Item struct {
	ID          string        `json:"id"`
	Title       string        `json:"title"`
	Kind        string        `json:"kind"`
	ItemDetails []*ItemDetail `json:"itemDetails"`
}

type ItemDetail struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	ItemID      string `json:"itemId"`
	Kind        string `json:"kind"`
	MoveMinutes int    `json:"moveMinutes"`
	Place       string `json:"place"`
	URL         string `json:"url"`
	Memo        string `json:"memo"`
	Priority    int    `json:"priority"`
}

type ShareItem struct {
	ID     string `json:"id"`
	ItemID string `json:"itemId"`
	Date   string `json:"date"`
	Item   *Item  `json:"item"`
}

type User struct {
	ID   string `json:"id"`
	UID  string `json:"uid"`
	Role int    `json:"role"`
}
