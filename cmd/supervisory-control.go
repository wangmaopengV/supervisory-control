package main

import (
	"encoding/json"
	"flag"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"supervisory-control/config"
	"supervisory-control/server"
	"supervisory-control/sql"
	"github.com/onrik/logrus/filename"
	"supervisory-control/agreement"
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

	//log init
	if len(config.GlobalConfig.LogLevel) > 0 {
		if logLevel, err := log.ParseLevel(config.GlobalConfig.LogLevel); err == nil {
			log.SetLevel(logLevel)
		} else {
			log.Warn("Parse log level config string failed: ", err)
		}
	}

	//link db
	db, err := sql.NewDBClient(config.GlobalConfig.DBMeta.Endpoint, config.GlobalConfig.DBMeta.KeySpace, config.GlobalConfig.DBMeta.UserName, config.GlobalConfig.DBMeta.Password)
	if err != nil {
		log.Fatal("NewDBClient Failed ", err, config.GlobalConfig.DBMeta)
	}
	config.GlobalConfig.DBClient = db

	//run task
	agreement.RunTask(config.GlobalConfig.AgreementMeta)

	//tcp server
	go server.TcpRun(config.GlobalConfig.PortMeta.Tcp, config.GlobalConfig.AgreementMeta.LinkSize)

	//http server
	server.RunServer(config.GlobalConfig.PortMeta.GRpc, config.GlobalConfig.PortMeta.Http)
}
