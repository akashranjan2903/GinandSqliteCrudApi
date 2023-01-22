package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ginSqliteCrud/utils"
)

type Blog struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	IsComplete bool   `json:"iscomplete"`
}
type bloglist struct {
	blogStore []Blog
	db        *sql.DB
}

func Service(db *sql.DB) *bloglist {
	return &bloglist{
		blogStore: []Blog{},
		db:        db,
	}
}
func (b *bloglist) Createblog(c *gin.Context) {

	if !utils.Checkmethod(c.Request.Method, utils.POST) {
		utils.ResponseWriter(c, http.StatusMethodNotAllowed, nil, "Method not match")
	}
	// for json
	// b.LoadFromJson()
	var newblog Blog
	err := json.NewDecoder(c.Request.Body).Decode(&newblog)
	if err != nil {
		log.Fatal("error found in decoding data")
		panic(err)
	}
	stmt, err := b.db.Prepare("INSERT INTO blog (title, body) VALUES (?,?);")

	if err != nil {
		panic(err)
	}
	row, err := stmt.Exec(newblog.Title, newblog.Body)

	if err != nil {
		panic(err)
	}
	id, err := row.LastInsertId()
	if err != nil {
		panic(err)
	}
	newblog.Id = int(id)
	// for saving to json file
	// b.blogStore = append(b.blogStore, newblog)
	// b.SavetoJson()

	utils.ResponseWriter(c, http.StatusOK, newblog, "New Blog created Successfully:")
}
func (b *bloglist) Getblog(c *gin.Context) {
	if !utils.Checkmethod(c.Request.Method, utils.GET) {
		utils.ResponseWriter(c, http.StatusMethodNotAllowed, nil, "Method not match")
	}
	// by json
	// b.LoadFromJson()
	// data := b.blogStore
	// if len(data) > 0 {
	// 	utils.ResponseWriter(c, http.StatusOK, data, "Data Present is:")
	// } else {
	// 	utils.ResponseWriter(c, http.StatusNotFound, nil, "Data Not Found")
	// }

	// by database
	rows, err := b.db.Query("SELECT * FROM blog;")
	if err != nil {
		panic(err)
	}
	var data = []Blog{}

	for rows.Next() {
		var newblog Blog
		rows.Scan(&newblog.Id, &newblog.Title, &newblog.Body, &newblog.IsComplete)
		data = append(data, newblog)
	}
	if len(data) > 0 {
		utils.ResponseWriter(c, http.StatusOK, data, "Data Present is")
	} else {
		utils.ResponseWriter(c, http.StatusNotFound, nil, "Data Not Found")
	}

}
func (b *bloglist) Deleteblog(c *gin.Context) {

	if !utils.Checkmethod(c.Request.Method, utils.DELETE) {
		utils.ResponseWriter(c, http.StatusMethodNotAllowed, nil, "Method not match")
	}
	// for json
	// b.LoadFromJson()
	id := utils.Getidfromurl(c.Param("id"))
	stmt, err := b.db.Prepare("DELETE FROM blog WHERE id=?;")
	if err != nil {
		panic(err)
	}

	row, e := stmt.Exec(id)
	if e != nil {
		utils.ResponseWriter(c, http.StatusOK, nil, "Invalid data id")
		panic(e)
	}
	// for json
	// for key, v := range b.blogStore {

	// 	if v.Id == id {
	// 		b.blogStore = append(b.blogStore[:key], b.blogStore[key+1:]...)
	// 		b.SavetoJson()
	// 		utils.ResponseWriter(c, http.StatusOK, id, "Data Deleted with id is:")
	// 		return
	// 	}

	// }
	rowaffect, _ := row.RowsAffected()
	if rowaffect > 0 {
		utils.ResponseWriter(c, http.StatusOK, id, "Data Deleted with given id is")
	} else {
		utils.ResponseWriter(c, http.StatusNotFound, nil, "Data Not Found")
	}

}

func (b *bloglist) Updateblog(c *gin.Context) {

	if !utils.Checkmethod(c.Request.Method, utils.PATCH) {
		utils.ResponseWriter(c, http.StatusMethodNotAllowed, nil, "Method not match")
	}
	// for json
	// b.LoadFromJson()
	id := utils.Getidfromurl(c.Param("id"))

	var newblog Blog
	err := json.NewDecoder(c.Request.Body).Decode(&newblog)
	utils.Errorhandlefordataconversion(err)

	//  For json
	// for key, v := range b.blogStore {

	// 	if v.Id == id {
	// 		v.Body = newblog.Body
	// 		v.Title = newblog.Title
	// 		b.blogStore[key] = v
	// 		b.SavetoJson()
	// 		utils.ResponseWriter(c, http.StatusOK, id, "Data Updates with id is:")
	// 		return
	// 	}

	// }

	//  for Db
	stmt, err := b.db.Prepare("UPDATE blog SET title = ? ,body = ? WHERE id=?;")
	if err != nil {
		panic(err)
	}
	row, e := stmt.Exec(&newblog.Title, &newblog.Body, id)
	if e != nil {
		utils.ResponseWriter(c, http.StatusOK, nil, "Invalid statement")
		panic(e)
	}
	rowaffect, _ := row.RowsAffected()
	if rowaffect > 0 {
		utils.ResponseWriter(c, http.StatusOK, id, "Data Updates with given id is")
	} else {
		utils.ResponseWriter(c, http.StatusNotFound, nil, "Data Not Found")
	}
}

func (b *bloglist) Getblogbyid(c *gin.Context) {

	if !utils.Checkmethod(c.Request.Method, utils.GET) {
		utils.ResponseWriter(c, http.StatusMethodNotAllowed, nil, "Method not match")
	}
	// for json
	// b.LoadFromJson()
	id := utils.Getidfromurl(c.Param("id"))
	row := b.db.QueryRow("SELECT * FROM blog WHERE id=?;", id)

	var newblog Blog
	err := row.Scan(&newblog.Id, &newblog.Title, &newblog.Body, &newblog.IsComplete)
	if err != nil {
		utils.ResponseWriter(c, http.StatusOK, nil, "No blog with the given id")
		return
	}
	utils.ResponseWriter(c, http.StatusOK, newblog, "Data with the given id")

	// For json
	// var newblog Blog
	// for _, v := range b.blogStore {

	// 	if v.Id == id {
	// 		newblog = v
	// 		utils.ResponseWriter(c, http.StatusOK, newblog, "Data with the given id")
	// 		return
	// 	}

	// }

}
func (b *bloglist) Changestatus(c *gin.Context) {
	if !utils.Checkmethod(c.Request.Method, utils.PATCH) {
		utils.ResponseWriter(c, http.StatusMethodNotAllowed, nil, "Method not match")
	}
	id := utils.Getidfromurl(c.Param("id"))

	var newblog Blog
	err := json.NewDecoder(c.Request.Body).Decode(&newblog)
	utils.Errorhandlefordataconversion(err)

	stmt, err := b.db.Prepare("UPDATE blog SET iscomplete = ? WHERE id=?;")
	if err != nil {
		panic(err)
	}
	row, e := stmt.Exec(&newblog.IsComplete, id)
	if e != nil {
		utils.ResponseWriter(c, http.StatusOK, nil, "Invalid statement")
		panic(e)
	}
	rowaffect, _ := row.RowsAffected()
	if rowaffect > 0 {
		utils.ResponseWriter(c, http.StatusOK, id, "Data Updates with given id is")
	} else {
		utils.ResponseWriter(c, http.StatusNotFound, nil, "Data Not Found")
	}
}
