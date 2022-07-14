package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}





var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

const UserCollection = "user"
var (
       errNotExist        = errors.New("Users are not exist")
       errInvalidID       = errors.New("Invalid ID")
       errInvalidBody     = errors.New("Invalid request body")
       errInsertionFailed = errors.New("Error in the user insertion")
       errUpdationFailed  = errors.New("Error in the user updation")
       errDeletionFailed  = errors.New("Error in the user deletion")
)

func getAcsess(c *gin.Context){

	c.IndentedJSON(http.StatusOK ,albums)
	}

func postAcsess(c *gin.Context){
	var newData album
	c.BindJSON(&newData)
	albums=append(albums,newData)
	c.IndentedJSON(http.StatusCreated,  newData)

}

func getItemById(c *gin.Context){
	id := c.Param("id")
	for _,a:=range albums{
		if a.ID == id {
		c.IndentedJSON(http.StatusOK,a)
		}
	}

}

func recordDelete(c *gin.Context){
	id := c.Param("id")
	for k,a:= range albums{
		if  a.ID==id {
			albums =append(albums[:k],albums[k+1:]... )
			c.IndentedJSON(http.StatusOK,albums)
		}
	}
}

func updateRecord(c *gin.Context){
	id := c.Param("id")
	var newDate album
	c.BindJSON(&newDate)
	fmt.Println(newDate)
	for k,a:= range albums{
		if a.ID == id {
			albums[k]= newDate
			c.IndentedJSON(http.StatusOK,albums)
		}
	}

}
func main(){

	router := gin.Default()
	router.GET("/AcsessGet", getAcsess)
	router.POST("/AcsessPost", postAcsess)
	router.GET("/AcsessGet/:id",getItemById)
	router.DELETE("/AcsessDelete/:id",recordDelete)
	router.PUT("/updateRecord/:id",updateRecord)
	router.Run("localhost:8080")
}