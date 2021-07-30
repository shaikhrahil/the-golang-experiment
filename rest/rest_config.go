package rest

import (
	"fmt"
	"log"

	"github.com/tkanos/gonfig"
)

func GetConfig(path string) Configuration {
	configuration := Configuration{}
	file := fmt.Sprintf(`./%s_config.json`, path)
	if err := gonfig.GetConf(file, &configuration); err != nil {
		log.Fatalln(err.Error())
	}
	return configuration
}

type Configuration struct {
	APP     AppConf
	DB      DBConf
	AUTH    AuthConf
	ACCOUNT AccountConf
	TODO    TodoConf
}

type AppConf struct {
	NAME    string
	PORT    string
	VERSION string
}

type DBConf struct {
	USERNAME string
	PASSWORD string
	PORT     string
	HOST     string
	NAME     string
}

type AuthConf struct {
	PREFIX          string
	JWT_TTL         string
	JWT_REFRESH_TTL string
	JWT_SECRET      string
}

type AccountConf struct {
	PREFIX string
}

type TodoConf struct {
	PREFIX string
}
