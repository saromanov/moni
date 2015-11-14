package moni

import (
	"log"
	"time"

	"github.com/saromanov/goconfig"
	"github.com/saromanov/dirwatcher"
)

//Config provides basic configuration for moni
type Config struct {
	Timeout  time.Duration
	//Write totification to Email
	Email    string
	//Write notification to Telegram
	Telegram string
	//Write result to the outpath
	Outpath string
	//ReloadConfig provides reloading config during running
	ReloadConfig bool 
	//Hosts for monitoring
	Hosts    []*Host
}

// reload provides reloading config
func (conf *Config) reload(path string) {
	watcher := dirwatcher.Init()
	watcher.AddFile(path, func(item string, d *dirwatcher.DirWatcher){
		newconfig := LoadConfigData(path)
		conf = newconfig
	})
	watcher.Run()
}

//LoadConfigData provides load configuration or set default params
func LoadConfigData(path string)*Config {
	conf := Config{}
	err := goconfig.Load(path, &conf)
	if err != nil {
		log.Printf("Config file %s is not found", path)
	}

	conf.setMissedValues()
	return &conf
}

//TODO
func (conf *Config) setMissedValues() {
	conf.Timeout = 1 * time.Second
	conf.Hosts = []*Host{}
	conf.Outpath = "moniout"
}

func setDefaultParams() *Config {
	conf := new(Config)
	conf.Timeout = 60 * time.Second
	return conf
}
