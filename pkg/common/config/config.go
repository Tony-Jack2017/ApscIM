package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf Config

type Config struct {
	Database   `json:"database" mapstructure:"database"`
	ObjStorage `json:"objStorage" mapstructure:"objStorage"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Driver   string `mapstructure:"driver"`
	DbName   string `mapstructure:"dbname"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	DbName   string `mapstructure:"dbname"`
	Password string `mapstructure:"password"`
}

type MongoDBConfig struct {
	Host   string `mapstructure:"host"`
	Port   string `mapstructure:"port"`
	DbName string `mapstructure:"dbname"`
}

type Database struct {
	Mysql   MysqlConfig   `json:"mysql" mapstructure:"mysql"`
	Redis   RedisConfig   `json:"redis" mapstructure:"redis"`
	MongoDB MongoDBConfig `json:"mongodb" mapstructure:"mongodb"`
}

type ObjStorage struct {
	Host   string `mapstructure:"host"`
	Port   string `mapstructure:"port"`
	DbName string `mapstructure:"dbname"`
}

func Init() error {
	viper.SetConfigFile("../config/config-dev.yaml")
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("The config file is changed ... ")
		err := viper.Unmarshal(&Conf)
		if err != nil {
			panic(fmt.Errorf("config file update occure the error: %v", err))
		}
	})

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("when read config file occure the error: %v", err))
	}

	if errUnmarshal := viper.Unmarshal(&Conf); errUnmarshal != nil {
		panic(fmt.Errorf("unmarshal the config to conf struc occure the error: %v", errUnmarshal))
	}
	return nil
}
