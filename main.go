package main

import (
	"github.com/gin-gonic/gin"

  "github.com/hcnp/trireg2-import-clockify/cmd"
)

func main() {
  run()
}

func run() {
  router := gin.Default()
  router.GET("/", cmd.FrontPage)
  router.GET("/:name", cmd.FrontPage)

  router.Run(":8080")
}

