package mongo

import (
	"ApscIM/pkg/model/base/user"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewUserMongo(db *mongo.Database) (user.MgoUserInterface, error) {
	col := db.Collection("user")
	_, err := col.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{
			{Key: "user_id", Value: 1},
		},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return nil, err
	}
	return &UserMgo{
		col: col,
	}, nil
}

type UserMgo struct {
	col *mongo.Collection
}

func (u *UserMgo) CreateUserMgo(ctx context.Context, users *user.User) (err error) {
	return nil
}
func (u *UserMgo) UpdateUserMgo(ctx context.Context, user *user.User) (err error) {
	return nil
}
func (u *UserMgo) GetUsersMgo(ctx context.Context, user *user.User) (users *[]user.User, err error) {
	return nil, nil
}

func (u *UserMgo) CreateUserSettingMgo(ctx context.Context, setting user.SettingUser) (err error) {
	return nil
}
func (u *UserMgo) UpdateUserSettingMgo(ctx context.Context, setting user.SettingUser) (err error) {
	return nil
}
func (u *UserMgo) GetUserSettingMgo(ctx context.Context, userID int32) (setting *user.SettingUser, err error) {
	return nil, nil
}
