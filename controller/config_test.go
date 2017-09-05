package controller

import (
	"io/ioutil"
	"testing"

	"github.com/gin-gonic/gin"
	yaml "gopkg.in/yaml.v1"
)

func testGetDBTest(t *testing.T) {

	te := TempDBConfig{}
	c := DBConfig{}

	bytes, err := ioutil.ReadFile("../dbconfig.yml")
	if err != nil {
		t.Error(err)
		return
	}

	err = yaml.Unmarshal(bytes, &te)
	if err != nil {
		t.Error(err)
		return
	}
	if gin.IsDebugging() {
		c = te.Development
	} else {
		c = te.Production
	}

	if c.Datasource == "" || c.Dialect == "" || c.Dir == "" {
		t.Error("Datasource::", c.Datasource, " Dialect::", c.Dialect, " Dir::", c.Dir)
		return
	}
}
