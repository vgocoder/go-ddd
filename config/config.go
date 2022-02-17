package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

var Conf *Config

type Config struct {
	App   AppConfig   `json:"app,omitempty"`
	Mysql MysqlConfig `json:"mysql,omitempty"`
	Redis RedisConfig `json:"redis,omitempty"`
}

type AppConfig struct {
	Name string `json:"name,omitempty"`
	Port string `json:"port,omitempty"`
}

type MysqlConfig struct {
	Host    string `json:"host,omitempty"`
	Port    string `json:"port,omitempty"`
	User    string `json:"user,omitempty"`
	Pass    string `json:"pass,omitempty"`
	DbName  string `json:"db_name,omitempty"`
	Charset string `json:"charset,omitempty"`
}

type RedisConfig struct {
	Host string `json:"host,omitempty"`
	Port string `json:"port,omitempty"`
	Db   string `json:"db,omitempty"`
}

func init() {
	f := "config.toml"
	if _, err := os.Stat(f); err != nil {
		f = "config.toml"
	}

	_, err := toml.DecodeFile(f, &Conf)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
