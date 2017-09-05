package controller

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (d Database) CreateUser(c *gin.Context) {
	a, _ := c.GetPostForm("a")
	log.Println(a)

}

func (d Database) GetOneUser(c *gin.Context) {

}
