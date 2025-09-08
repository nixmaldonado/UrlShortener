package main

import (
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

var conf Conf

type Conf struct {
	Port     string `default:":8081"`
	CodeSize int    `default:"7" split_words:"true"`
}

func InitConfig() {
	err := envconfig.Process("SHORT_URL_APP", &conf)
	if err != nil {
		panic(err)
	}

	log.Info("Configs initialized", zap.Int("CodeSize", conf.CodeSize))
}
