package main

import
(
	"time"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

//Config provides basic configuration for moni
type Config struct {
	Timeout *time.Time
	Email   string
	Telegram string
}

//LoadConfigData provides load configuration or set default params
func LoadConfigData(path string) *Config {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return setDefaultParams()
	}
	var conf Config
	yamlerr := yaml.Unmarshal(data, &conf)
	if yamlerr != nil {
		panic(yamlerr)
	}

	conf.setMissedValues()
	return &conf
}


//TODO
func (conf *Config) setMissedValues() {
	
}

func setDefaultParams() *Config {
	conf := new(Config)
	conf.Timeout = time.Duration(1) * time.Minutes
	return conf
}