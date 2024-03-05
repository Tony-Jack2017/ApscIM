package db

type DataStorageManage struct {
	perStorage   MysqlService
	shortStorage MongoService
	objStorage   MinioService
	cache        CacheRedis
}

func RunDSM() (*DataStorageManage, error) {
	NewMysqlService()
	NewRedisService()
	NewMinioService()
	NewMongoService()
	return &DataStorageManage{}, nil
}
