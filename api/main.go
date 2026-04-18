package api

import (
	"net/http"

	. "github.com/tbxark/g4vercel"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := New()

	server.GET("/", func(c *Context) {
		c.JSON(http.StatusOK, H{
			"message": "hello",
		})
	})

	server.Handle(w, r)
}
