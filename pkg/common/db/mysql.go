package db

import (
	"ApscIM/pkg/common/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlService struct {
	conn *gorm.DB
}

func NewMysqlService() (*MysqlService, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Conf.Database.Mysql.User,
		config.Conf.Database.Mysql.Password,
		config.Conf.Database.Mysql.Host,
		config.Conf.Database.Mysql.Port,
		config.Conf.Database.Mysql.DbName,
	)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})

	if err != nil {
		panic(fmt.Errorf("when connecting the mysql service occure the error: %v", err))
		return nil, err
	}
	return &MysqlService{
		conn: db,
	}, nil
}
