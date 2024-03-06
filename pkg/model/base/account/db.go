package account

type DatabaseAccount struct {
	mysql SqlAccountInterface
	redis RdsAccountInterface
	mongo MgoAccountInterface
}

func NewAccountDatabase() {
}
