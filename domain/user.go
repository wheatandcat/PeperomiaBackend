package domain

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/wheatandcat/PeperomiaBackend/graph/model"
)

// UserRoleAdmin 管理者
const UserRoleAdmin = 1

// UserRecord is user data
type UserRecord struct {
	UID        string             `json:"uid" firestore:"uid"`
	Role       int                `json:"role" firestore:"role"`
	CreatedAt  time.Time          `json:"createdAt" firestore:"createdAt"`
	UpdatedAt  time.Time          `json:"updatedAt" firestore:"updatedAt"`
	PushTokens []*PushTokenRecord `json:"pushTokens" firestore:"pushTokens"`
}

// UserRepository is repository interface
type UserRepository interface {
	Create(ctx context.Context, f *firestore.Client, u UserRecord) error
	FindByUID(ctx context.Context, f *firestore.Client, uid string) (UserRecord, error)
	ExistsByUID(ctx context.Context, f *firestore.Client, uid string) (bool, error)
}

// ToModel Modelに変換する
func (r *UserRecord) ToModel() *model.User {
	var epts []*model.ExpoPushToken

	for _, pt := range r.PushTokens {
		epts = append(epts, pt.ToModel())
	}

	u := &model.User{
		UID:            r.UID,
		Role:           r.Role,
		ExpoPushTokens: epts,
	}

	return u
}
