package main

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	HappyMeter struct {
		Url  string `yaml:"url"`
		Tags string `yaml:"tags"`
	}
}

func ReadConfig(file string) (*Config, error) {
	config := &Config{}
	configData, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
		return nil, err
	}

	return config, nil
}
