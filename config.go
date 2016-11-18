package main

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	LogFile string `yaml:"logfile"`
}

func LoadConfig() Config {
	file, err := ioutil.ReadFile("/etc/sendmail-logger/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	conf := Config{}
	err = yaml.Unmarshal([]byte(file), &conf)
	return conf
}
