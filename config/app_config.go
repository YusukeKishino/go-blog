package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

const Dev = "development"
const Test = "test"
const Prod = "production"

type AppConfig struct {
	DBUrl    string `yaml:"database_url"`
	Salt     string `yaml:"salt"`
	S3Bucket string `yaml:"bucket"`
}

func GetConfig() (*AppConfig, error) {
	env := GetEnv()

	conf, err := readSettings(env)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

func GetEnv() string {
	env := os.Getenv("GO_APP_ENV")
	if env == "" {
		env = Dev
	}

	return env
}

func IsDev() bool {
	return GetEnv() == Dev
}

func IsTest() bool {
	return GetEnv() == Test
}

func IsProd() bool {
	return GetEnv() == Prod
}

func readSettings(env string) (*AppConfig, error) {
	file, err := ioutil.ReadFile("config/settings.yml")
	if err != nil {
		return nil, err
	}

	file = []byte(os.ExpandEnv(string(file)))
	confs := make(map[string]*AppConfig)
	if err := yaml.Unmarshal(file, confs); err != nil {
		return nil, err
	}

	conf := confs[env]
	if conf == nil {
		return nil, fmt.Errorf("environment '%s' is not found on config/settings.yml", env)
	}

	return conf, nil
}
