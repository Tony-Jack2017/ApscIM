package user

type DatabaseUser struct {
	mysql SqlUserInterface
	redis RdsUserInterface
	mongo MgoUserInterface
}

func NewUserDatabase() {
}

// Redis

func (d *DatabaseUser) AddUserKey() {}

// Mongo

func (d *DatabaseUser) CreateUser() {}

func (d *DatabaseUser) UpdateUser() {}

func (d *DatabaseUser) GetUserList() {}

func (d *DatabaseUser) SetUserSetting() {}

func (d *DatabaseUser) UpdateUserSetting() {}

func (d *DatabaseUser) GetUserSettings() {}

// Mysql
