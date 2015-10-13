package main

import
(
	"time"
	"github.com/saromanov/goconfig"
	"log"
)

//Config provides basic configuration for moni
type Config struct {
	Timeout time.Duration
	Email   string
	Telegram string
	Hosts    []*Host
}

type Host struct {
	Username  string
	Password  string
}

//LoadConfigData provides load configuration or set default params
func LoadConfigData(path string) *Config {
	conf := Config{}
	err := goconfig.Load(path, &conf)
	if err != nil {
		log.Fatal(err)
	}
	
	conf.setMissedValues()
	return &conf
}


//TODO
func (conf *Config) setMissedValues() {
	
}

func setDefaultParams() *Config {
	conf := new(Config)
	conf.Timeout = 60* time.Second
	return conf
}