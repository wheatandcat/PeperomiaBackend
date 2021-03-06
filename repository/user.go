package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/wheatandcat/PeperomiaBackend/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// UserRepository is repository for user
type UserRepository struct {
}

// NewUserRepository is Create new UserRepository
func NewUserRepository() domain.UserRepository {
	return &UserRepository{}
}

// Create ユーザーを作成する
func (re *UserRepository) Create(ctx context.Context, f *firestore.Client, u domain.UserRecord) error {

	_, err := f.Collection("version/1/users").Doc(u.UID).Set(ctx, u)

	return err
}

// FindByUID ユーザーIDから取得する
func (re *UserRepository) FindByUID(ctx context.Context, f *firestore.Client, uid string) (domain.UserRecord, error) {
	var u domain.UserRecord
	dsnap, err := f.Collection("version/1/users").Doc(uid).Get(ctx)
	if err != nil {
		return u, err
	}

	dsnap.DataTo(&u)

	pts, err := GetPushTokensByDocument(ctx, dsnap)
	if err != nil {
		return u, err
	}

	u.PushTokens = pts

	return u, nil
}

// ExistsByUID ユーザーIDが存在するか判定
func (re *UserRepository) ExistsByUID(ctx context.Context, f *firestore.Client, uid string) (bool, error) {
	dsnap, err := f.Collection("version/1/users").Doc(uid).Get(ctx)

	if err != nil {
		if grpc.Code(err) == codes.NotFound {
			return false, nil
		}
		return false, err
	}

	return dsnap.Exists(), nil
}
