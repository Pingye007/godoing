package config

import (
	"github.com/BurntSushi/toml"
)

var Cfg Config

type DB struct {
	User         string `toml:"user"`
	Password     string `toml:"password"`
	SqlType      string `toml:"sql_type"`
	ServerAddr   string `toml:"server_addr"`
	DatabaseName string `toml:"database_name"`
	Port         int    `toml:"port"`
}

type SystemConfig struct {
	Version string `toml:"version"`
}

type Log struct {
	Level string `toml:"level"`
}

type Config struct {
	DB        DB           `toml:"db"`
	Log       Log          `toml:"log"`
	SysConfig SystemConfig `toml:"system_config"`
}

func ParseConfigFile() error {
	_, err := toml.DecodeFile("config/config.toml", &Cfg)
	if err != nil {
		panic("decode configuration failed")
	}
	return nil
}

func init() {
	ParseConfigFile()
}
