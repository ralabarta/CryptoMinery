package config

import (
	"github.com/yuchenfw/gocrypt/rsa"
)

type AppConfig struct {
	Server       Server
	MysqlMaster  MySQL `mapstructure:"mysql-master"`
	MysqlSlave   MySQL `mapstructure:"mysql-slave"`
	Redis        Redis
	Rsa          rsa.RSASecret `mapstructure:"rsa"`
	Log          Log
	Scheduler    Scheduler
	InternalHost string `mapstructure:"internal-host"`
	InternalIp   string `mapstructure:"internal-ip"`
}

// Server 服务器配置
type Server struct {
	Port uint16
}

// MySQL mysql配置
type MySQL struct {
	Host         string
	Port         uint16
	Db           string
	Username     string
	Password     string
	Charset      string
	Timeout      string
	ParseTime    bool
	Loc          string
	MaxIdleConns int
	MaxOpenConns int
}

// Redis redis配置
type Redis struct {
	Host     string
	Port     uint16
	Db       int
	Password string
}

// Log log配置
type Log struct {
	Level      string
	Filename   string
	MaxSize    int `mapstructure:"maxsize"`
	MaxAge     int `mapstructure:"max_age"`
	MaxBackups int `mapstructure:"max_backups"`
}

// Scheduler 任务调度器
type Scheduler map[string]string

type Hoo struct {
	ClientId     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
	HOST         string `mapstructure:"host"`
	ApiHost      string `mapstructure:"api_host"`
}

type HtmlConf struct {
	Protocol    string
	Host        string
	ProcessType int64
}

type PoolConf struct {
	Puid  string
	Token string
}
