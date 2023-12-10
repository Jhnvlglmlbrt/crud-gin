package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	log.Print("Server starting...")
	r := gin.Default()

	r.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, "welcome home")
	})

	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
