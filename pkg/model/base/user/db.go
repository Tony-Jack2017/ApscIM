package user

type DatabaseUser struct {
	mysql SqlUserInterface
	redis RdsUserInterface
	mongo MgoUserInterface
}

func NewUserDatabase() {
}
