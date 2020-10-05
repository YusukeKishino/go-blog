package registry

import (
    "go.uber.org/dig"

    // TODO: Fixme
    "github.com/YusukeKishino/go-blog/config"
    "github.com/YusukeKishino/go-blog/db"
    "github.com/YusukeKishino/go-blog/server"
)

func BuildContainer() (*dig.Container, error) {
    c := dig.New()

    providers := []*provider{
        newProvider(config.GetConfig),
        newProvider(db.ConnectDB),
        newProvider(server.NewServer),
    }

    if err := setProviders(c, providers); err != nil {
        return nil, err
    }

    return c, nil
}

type provider struct {
    target interface{}
    opts   []dig.ProvideOption
}

func newProvider(target interface{}, opts ...dig.ProvideOption) *provider {
    return &provider{target: target, opts: opts}
}

func setProviders(container *dig.Container, providers []*provider) error {
    for _, p := range providers {
        if err := container.Provide(p.target, p.opts...); err != nil {
            return err
        }
    }
    return nil
}
