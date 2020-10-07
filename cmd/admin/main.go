package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"syscall"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ssh/terminal"
	"gorm.io/gorm"

	"github.com/YusukeKishino/go-blog/model"
	"github.com/YusukeKishino/go-blog/registry"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	fmt.Println("Input admin name: ")
	stdin.Scan()
	name := stdin.Text()
	if name == "" {
		log.Fatalln("Please input admin name")
	}

	fmt.Println("Input password: ")
	password, err := terminal.ReadPassword(syscall.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln(err)
	}
	admin := model.Admin{
		Name:     name,
		Password: string(hashedPassword),
	}

	c, err := registry.BuildContainer()
	if err != nil {
		log.Fatalln(err)
	}
	err = c.Invoke(func(db *gorm.DB) {
		if err := db.Create(&admin).Error; err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Success")
	})
	if err != nil {
		log.Fatalln(err)
	}
}
