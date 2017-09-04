package controller

import (
	"io/ioutil"

	"log"

	"github.com/gin-gonic/gin"
	yaml "gopkg.in/yaml.v1"
)

type TempDBConfig struct {
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

type DBConfig struct {
	Dialect    string
	Datasource string
	Dir        string
}

func GetDBconfig() (DBConfig, error) {
	t := TempDBConfig{}
	c := DBConfig{}
	bytes, err := ioutil.ReadFile("dbconfig.yml")
	if err != nil {
		return c, err
	}
	err = yaml.Unmarshal(bytes, &t)
	if err != nil {
		return c, err
	}

	if gin.IsDebugging() {
		log.Println("ISDEBUG!!")
		c = t.Development
	} else {
		log.Println("NOTDEBUG!!")
		c = t.Production
	}
	return c, nil
}
