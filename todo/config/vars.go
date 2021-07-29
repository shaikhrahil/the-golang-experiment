package config

import (
	"fmt"
	"log"

	"github.com/shaikhrahil/the-golang-experiment/rest"
	"github.com/tkanos/gonfig"
)

type config struct {
	rest.Configuration
	JWT_TOKEN_EXPIRY         string
	JWT_REFRESH_TOKEN_EXPIRY string
	JWT_SECRET               string
}

func GetConfig(path string) config {
	configuration := config{}
	fileName := fmt.Sprintf("./%s_config.json", path)
	if err := gonfig.GetConf(fileName, &configuration); err != nil {
		log.Fatalln(err.Error())
	}
	return configuration
}
