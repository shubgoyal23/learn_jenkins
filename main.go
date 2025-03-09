package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Post struct {
	Name     string         `json:"name"`
	ID       string         `json:"id"`
	Response *http.Response `json:"response"`
	Error    error          `json:"error"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}
	router := setupRouter()
	router.Run(port)
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/post", func(c *gin.Context) {
		obj := make(map[string]interface{})
		if err := c.ShouldBindJSON(&obj); err != nil {
			c.JSON(400, gin.H{
				"message": "error",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "post",
			"data":    obj,
		})
	})
	return router
}

func test() {
	all := make([]Post, 0)
	t := time.Now()
	for time.Since(t) < 10*time.Second {
		var post Post
		post.Name = "test"
		post.ID = uuid.New().String()

		jsonBody, _ := json.Marshal(post)
		res, err := http.Post("http://localhost:3000/post", "application/json", bytes.NewBuffer(jsonBody))
		if err != nil {
			post.Error = err
		} else {
			post.Response = res
		}
		all = append(all, post)
	}
	for _, post := range all {
		if post.Error != nil {
			println(post.Error.Error())
		} else {
			println(post.Response.Status)
		}
	}
	print("done", len(all))
}
