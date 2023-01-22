package utils

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type httpMethod string

const (
	GET    httpMethod = "GET"
	POST   httpMethod = "POST"
	DELETE httpMethod = "DELETE"
	PATCH  httpMethod = "PATCH"
)

func ResponseWriter(c *gin.Context, status int, data interface{}, message string) {

	c.JSON(status, gin.H{
		"code":    status,
		"message": message,
		"data":    data})

}
func Checkmethod(method string, checkmethod httpMethod) bool {
	return method == string(checkmethod)
}

func Getidfromurl(str string) int {

	id, e := strconv.Atoi(str)
	if e != nil {
		log.Fatal("Conversion of string into int failed")
		panic(e)
	}
	return id
}

func Errorhandlefordataconversion(err error) {

	if err != nil {
		log.Fatal("error found in Marshaling the data in json")
		panic(err)
	}
}
