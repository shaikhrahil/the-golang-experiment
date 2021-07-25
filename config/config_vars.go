package config

import (
	"fmt"
	"log"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
}

func GetConfig(config string) Configuration {
	configuration := Configuration{}
	fileName := fmt.Sprintf("./%s_config.json", config)
	if err := gonfig.GetConf(fileName, &configuration); err != nil {
		log.Fatalln(err.Error())
	}
	return configuration
}
