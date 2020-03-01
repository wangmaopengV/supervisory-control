package main

import (
	"encoding/json"
	"flag"
	"github.com/onrik/logrus/filename"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"supervisory-control/config"
)

var
(
	configPath = flag.String("configPath", "config/config.json", "config file path")
)

func ParseConfig(configPath string) () {

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal("ReadFile Failed  ", err, configPath)
	}

	config.GlobalConfig = &config.SupervisoryControlConfig{
	}
	if err := json.Unmarshal(data, config.GlobalConfig); err != nil {
		log.Fatal("Unmarshal Failed  ", err, configPath)
	}
}

func main() {

	//log init
	flag.Parse()
	log.AddHook(filename.NewHook())

	//config init
	ParseConfig(*configPath)
}
