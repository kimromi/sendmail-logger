package main

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	LogFile string `yaml:"logfile"`
}

func LoadConfig() (Config, error) {
	file, err := ioutil.ReadFile("/etc/sendmail-logger/config.yaml")
	if err != nil {
		return Config{}, err
	}

	conf := Config{}
	err = yaml.Unmarshal([]byte(file), &conf)
	if err != nil {
		return Config{}, err
	}
	return conf, nil
}
