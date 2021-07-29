package config

import (
	"fmt"
	"log"

	"github.com/shaikhrahil/the-golang-experiment/rest"
	"github.com/tkanos/gonfig"
)

type config struct {
	rest.Configuration
}

func GetConfig(path string) config {
	configuration := config{}
	fileName := fmt.Sprintf("./%s_config.json", path)
	if err := gonfig.GetConf(fileName, &configuration); err != nil {
		log.Fatalln(err.Error())
	}
	return configuration
}
