package main

import g "github.com/gin-gonic/gin"

func main() {
	r := g.Default()
	r.GET("/ping", func(c *g.Context) {
		c.JSON(200, g.H{
			"message": "pong",
		})
	})

	r.Run()
}
