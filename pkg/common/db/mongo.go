package db

type MongoService struct {
}

func NewMongoService() (*MongoService, error) {
	return &MongoService{}, nil
}
