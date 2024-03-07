package mysql

import (
	"ApscIM/pkg/model/base/user"
	"context"
	"gorm.io/gorm"
)

func NewUserSql() (user.SqlUserInterface, error) {
	return &UserSql{
		conn: &gorm.DB{},
	}, nil
}

type UserSql struct {
	conn *gorm.DB
}

func (u *UserSql) CreateUsers(ctx context.Context, user *user.User) (err error) {
	u.conn.Create(user)
	return nil
}

func (u *UserSql) UpdateUser(ctx context.Context, user *user.User) (err error) {
	return nil
}

func (u *UserSql) GetUsers(ctx context.Context, user *user.User) (users *[]user.User, err error) {
	return nil, nil
}

func (u *UserSql) CreateUserSetting(ctx context.Context, setting user.SettingUser) (err error) {
	return nil
}

func (u *UserSql) UpdateUserSetting(ctx context.Context, setting user.SettingUser) (err error) {
	return nil
}

func (u *UserSql) GetUserSetting(ctx context.Context, userID int32) (setting *user.SettingUser, err error) {
	return nil, nil
}
