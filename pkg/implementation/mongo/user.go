package mongo

import (
	"ApscIM/pkg/model/base/user"
	"ApscIM/pkg/tools/db"
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

func (u *UserMgo) CreateUsers(ctx context.Context, users []user.User) (err error) {
	return db.InsertMany(ctx, u.col, users)
}

func (u *UserMgo) UpdateUser(ctx context.Context, user *user.User) (err error) {
	return nil
}
