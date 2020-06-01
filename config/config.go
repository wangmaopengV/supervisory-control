package config

import "supervisory-control/sql"

type SCPortMeta struct {
	Tcp  int32
	Http int32
	GRpc int32
}

type SCAgreementMeta struct {
	TaskSize   int32
	TaskThread int32
	LinkSize   int32
}

type SCDBMeta struct {
	UserName string
	Password string
	Endpoint string
	KeySpace string
}

type SupervisoryControlConfig struct {
	AgreementMeta SCAgreementMeta
	PortMeta      SCPortMeta
	LogLevel      string

	DBMeta   SCDBMeta
	DBClient *sql.DBClient
}

var GlobalConfig *SupervisoryControlConfig
