package moni

import (
	"github.com/saromanov/goconfig"
	"log"
	"time"
)

//Config provides basic configuration for moni
type Config struct {
	Timeout  time.Duration
	Email    string
	Telegram string
	Hosts    []*Host
}

type Host struct {
	Username string
	Password string
}

//LoadConfigData provides load configuration or set default params
func LoadConfigData(path string)*Config {
	conf := Config{}
	err := goconfig.Load(path, &conf)
	if err != nil {
		log.Printf("Config file %s is not found", path)
		return nil
	}

	conf.setMissedValues()
	return &conf
}

//TODO
func (conf *Config) setMissedValues() {

}

func setDefaultParams() *Config {
	conf := new(Config)
	conf.Timeout = 60 * time.Second
	return conf
}
