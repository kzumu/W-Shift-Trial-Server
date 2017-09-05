package wsct

import (
	"database/sql"

	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shimokp/W-Shift-Trial-Server/controller"
)

type Server struct {
	DB *sql.DB
}

func New() *Server {
	return &Server{}
}

func (s *Server) SetupRouter() {
	database := controller.Database{DB: s.DB}

	router := gin.Default()
	v1 := router.Group("/api/v1/wsct")
	{
		v1.POST("/users", database.CreateUser)
		v1.GET("/users/:id", database.GetOneUser)
	}
	err := router.Run()
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func (s *Server) SetupDB() error {
	c, err := controller.GetDBconfig()
	if err != nil {
		return err
	}
	db, err := sql.Open(c.Dialect, c.Datasource)
	if err != nil {
		return err
	}
	s.DB = db
	return nil
}

//func (s *Server) DbTest() error {
//	_, err = db.DB.Exec(
//		`CREATE TABLE HOGE (
//	 id INTEGER,
//	 hoge VARCHAR(255))`,
//	)
//	if err != nil {
//		return err
//	}
//
//	_, err = db.Exec(
//		`INSERT INTO HOGE (id, hoge)
//		VALUES (?, ?) `,
//		1,
//		"hogege",
//	)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
