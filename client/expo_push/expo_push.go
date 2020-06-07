package expopush

import (
	"fmt"

	expo "github.com/oliveroneill/exponent-server-sdk-golang/sdk"
)

// ExpoPushClientGenerator ExpoPushClient
type ExpoPushClientGenerator interface {
	Send(req SendRequest) error
}

// ExpoPushClient is expo push client
type ExpoPushClient struct {
	PushClient expo.PushClient
}

// SendRequest is send request
type SendRequest struct {
	Body  string
	Data  map[string]string
	Title string
	Token string
}

// NewExpoPushClient is Create new NewExpoPushClient
func NewExpoPushClient() (*ExpoPushClient, error) {

	client := expo.NewPushClient(nil)

	return &ExpoPushClient{
		PushClient: *client,
	}, nil
}

// Send Push通知を送信
func (c *ExpoPushClient) Send(req SendRequest) error {
	pushToken, err := expo.NewExponentPushToken(req.Token)
	if err != nil {
		return err
	}

	response, err := c.PushClient.Publish(
		&expo.PushMessage{
			To:       pushToken,
			Body:     req.Body,
			Data:     req.Data,
			Sound:    "default",
			Title:    req.Title,
			Priority: expo.DefaultPriority,
		},
	)

	if err != nil {
		return err
	}

	if response.ValidateResponse() != nil {
		fmt.Println(response.PushMessage.To, "failed")
	}

	return nil
}
