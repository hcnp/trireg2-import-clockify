package main

//go:generate oapi-codegen --config=api/config.yaml trireg2-api/openapi/trireg2.yaml

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

  router.Run(":8080")
}
