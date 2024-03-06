package friend

type DatabaseFriend struct {
	mysql SqlFriendInterface
	redis RdsFriendInterface
	mongo MgoFriendInterface
}

func NewFriendDatabase() {

}
