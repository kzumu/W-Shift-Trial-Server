package controller

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v1"
)

type DBConfig struct {
	Development struct {
		Dialect    string
		Datasource string
		Dir        string
	}
	Production struct {
		Dialect    string
		Datasource string
		Dir        string
	}
}

func GetDBconfig() (DBConfig, error) {
	c := DBConfig{}
	bytes, err := ioutil.ReadFile("dbconfig.yml")
	if err != nil {
		return c, err
	}
	err = yaml.Unmarshal(bytes, &c)
	if err != nil {
		return c, err
	}
	return c, nil
}
