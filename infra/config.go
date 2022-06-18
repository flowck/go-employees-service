package infra

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DbDriver string `envconfig:"DB_DRIVER" required:"true"`
	DbUrl    string `envconfig:"DB_URL" required:"true"`
	PORT     int    `envconfig:"PORT" required:"true"`
	HOST     string `envconfig:"HOST" required:"true"`
}

var Configs Config

func init() {
	Configs = Config{}
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	err := envconfig.Process("", &Configs)

	if err != nil {
		panic(err)
	}
}
