package controller

import (
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shimokp/W-Shift-Trial-Server/model"
)

func (d Database) CreateUser(c *gin.Context) {
	a, _ := c.GetPostForm("uuid")
	log.Println("Key::", a)

	id, err := model.CreateUserByUUID(d.DB, a)
	if err != nil {
		log.Println(err)
	}

	log.Println("UserId::", *id)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"id":     *id,
	})
}

func (d Database) GetOneUser(c *gin.Context) {

}
