package mysql

import (
	"ApscIM/pkg/model/base/user"
	"context"
	"gorm.io/gorm"
)

type UserSql struct {
	conn *gorm.DB
}

func NewUserSql() (user.SqlUserInterface, error) {
	return &UserSql{
		conn: &gorm.DB{},
	}, nil
}

func (u *UserSql) CreateUserSql(ctx context.Context, user *user.User) (err error) {
	u.conn.Create(user)
	return nil
}

func (u *UserSql) UpdateUserSql(ctx context.Context, user *user.User) (err error) {
	return nil
}

func (u *UserSql) GetUsersSql(ctx context.Context, user *user.User) (users *[]user.User, err error) {
	return nil, nil
}

func (u *UserSql) CreateUserSettingSql(ctx context.Context, setting user.SettingUser) (err error) {
	return nil
}

func (u *UserSql) UpdateUserSettingSql(ctx context.Context, setting user.SettingUser) (err error) {
	return nil
}

func (u *UserSql) GetUserSettingSql(ctx context.Context, userID int32) (setting *user.SettingUser, err error) {
	return nil, nil
}
