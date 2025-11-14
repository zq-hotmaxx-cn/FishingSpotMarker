package data

import (
	"FishingSpotMarker/core/conf"
	"FishingSpotMarker/core/log"
	"FishingSpotMarker/pkg/config"
	"FishingSpotMarker/pkg/model"
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	fishing_spot_marker_database *gorm.DB
)

func Init() error {
	// 创建 Gorm 配置
	gormConfig := &gorm.Config{
		Logger: log.GetZapGormLogger(),
	}

	// 创建数据库连接
	var err error
	switch conf.GetConf().Database.Driver {
	case config.DATABASE_DRIVER_SQLITE:
		fishing_spot_marker_database, err = gorm.Open(sqlite.Open(conf.GetConf().Database.DSN), gormConfig)
	case config.DATABASE_DRIVER_MYSQL:
		fishing_spot_marker_database, err = gorm.Open(mysql.Open(conf.GetConf().Database.DSN), gormConfig)
	case config.DATABASE_DRIVER_POSTGRES:
		fishing_spot_marker_database, err = gorm.Open(postgres.Open(conf.GetConf().Database.DSN), gormConfig)
	default:
		return errors.New("Invalid database driver: " + conf.GetConf().Database.Driver)
	}
	if err != nil {
		return err
	}

	tables := []any{
		&model.User{},
		&model.Category{},
		&model.FishingSpotMarker{},
	}

	err = fishing_spot_marker_database.AutoMigrate(tables...)
	if err != nil {
		return err
	}

	return nil
}

func GetDataBase() *gorm.DB {
	return fishing_spot_marker_database
}
