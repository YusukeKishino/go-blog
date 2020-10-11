package main

import (
	"flag"

	"github.com/sirupsen/logrus"

	// TODO: Fixme
	"github.com/YusukeKishino/go-blog/registry"
	"github.com/YusukeKishino/go-blog/server"
)

var (
	addr string
)

func main() {
	flag.StringVar(&addr, "a", ":8080", "address to use")
	flag.Parse()

	container, err := registry.BuildContainer()
	if err != nil {
		logrus.Fatalln(err)
	}

	err = container.Invoke(func(s *server.Server) {
		err := s.Run(addr)
		if err != nil {
			logrus.Fatalln(err)
		}
	})
	if err != nil {
		logrus.Fatalln(err)
	}
}
