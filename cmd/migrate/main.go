package main

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/YusukeKishino/go-blog/registry"
)

func main() {
	container, err := registry.BuildContainer()
	if err != nil {
		logrus.Fatalln(err)
	}

	err = container.Invoke(func(db *gorm.DB) {
		// Add models to migrate
		err := db.AutoMigrate()
		if err != nil {
			logrus.Fatalln(err)
		}
	})
	if err != nil {
		logrus.Fatalln(err)
	}
}
