package controller

import (
	"log"
	"strconv"

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
		return
	}

	log.Println("UserId::", *id)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"id":     *id,
	})
}

func (d Database) GetOneUser(c *gin.Context) {
	strid := c.Param("id")

	id, err := strconv.Atoi(strid)
	if err != nil {
		log.Println(err)
	}

	user, err := model.FindUserById(d.DB, id)
	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"uuid": user.Uuid,
	})
}
