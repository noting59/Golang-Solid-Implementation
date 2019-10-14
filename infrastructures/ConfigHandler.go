package infrastructures

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
}

type Conf struct {
	PostgreSQLConn string `yaml:"postresql_conn"`
}

func (c *Config) GetConf() *Conf {

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	conf := Conf{}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return &conf
}
