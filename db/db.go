package db

import (
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	// TODO: fixme
	"github.com/YusukeKishino/go-blog/config"
)

func ConnectDB(conf *config.AppConfig) (*gorm.DB, error) {
	db, err := gorm.Open(
		mysql.New(mysql.Config{
			DSN: conf.DBUrl,
		}),
		&gorm.Config{
			AllowGlobalUpdate: false,
		})
	if err != nil {
		return nil, errors.Wrap(err, "failed to connecting database")
	}

	return db, nil
}
