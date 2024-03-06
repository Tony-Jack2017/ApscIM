package admin

type DatabaseAdmin struct {
	mysql SqlAdminInterface
	redis RdsAdminInterface
	mongo MgoAdminInterface
}

func NewAdminDatabase() {

}
