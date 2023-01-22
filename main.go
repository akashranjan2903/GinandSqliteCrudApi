package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ginSqliteCrud/controllers"
	_ "modernc.org/sqlite"
)

func main() {
	db := intializeDb()
	defer db.Close()
	newservice := controllers.Service(db)
	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	r.POST("/blog/create", newservice.Createblog)
	r.GET("/blog/read", newservice.Getblog)
	r.GET("/blog/getbyid/:id", newservice.Getblogbyid)
	r.DELETE("/blog/delete/:id", newservice.Deleteblog)
	r.PATCH("/blog/update/:id", newservice.Updateblog)
	r.PATCH("/blog/changeStatus/:id", newservice.Changestatus)

	r.Run(":3030")
}

func intializeDb() *sql.DB {
	db, err := sql.Open("sqlite", "test.db")
	if err != nil {
		panic(err)
	}
	// Create tables

	sql := `create table if not exists blog (id integer primary key autoincrement,title text,body text,iscomplete Boolean DEFAULT 0);`
	_, err = db.Exec(sql)
	if err != nil {
		log.Printf(" %q %s", err, sql)
		panic(err)
	}
	return db
}
