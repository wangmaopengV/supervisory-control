package main

import (
	"encoding/json"
	"flag"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"supervisory-control/config"
	"supervisory-control/server"
	"supervisory-control/sql"
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
	//log.AddHook(filename.NewHook())

	//config init
	ParseConfig(*configPath)

	//link db
	db, err := sql.NewDBClient(config.GlobalConfig.DBMeta.Endpoint, config.GlobalConfig.DBMeta.KeySpace, config.GlobalConfig.DBMeta.UserName, config.GlobalConfig.DBMeta.Password)
	if err != nil {
		log.Fatal("NewDBClient Failed ", err, config.GlobalConfig.DBMeta)
	}
	config.GlobalConfig.DBClient = db

	//tcp server
	go server.TcpRun(config.GlobalConfig.PortMeta.Tcp)

	//http server
	server.RunServer(config.GlobalConfig.PortMeta.GRpc, config.GlobalConfig.PortMeta.Http)
}
