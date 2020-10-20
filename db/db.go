package db

import (
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/YusukeKishino/go-blog/config"
)

func ConnectDB(conf *config.AppConfig) (*gorm.DB, error) {
	var logLebel logger.LogLevel
	if config.IsDev() {
		logLebel = logger.Info
	} else {
		logLebel = logger.Error
	}
	db, err := gorm.Open(
		mysql.New(mysql.Config{
			DSN: conf.DBUrl,
		}),
		&gorm.Config{
			AllowGlobalUpdate: false,
			Logger:            logger.Default.LogMode(logLebel),
		})
	if err != nil {
		return nil, errors.Wrap(err, "failed to connecting database")
	}

	return db, nil
}
