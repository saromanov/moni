package main

import
(
	"time"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"os"
)

//Config provides basic configuration for moni
type Config struct {
	timeout time*Time
	email   string
	telegram string
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
	conf.timeout = time.Duration(1) * time.Minutes
	return conf
}