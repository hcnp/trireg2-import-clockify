package cmd

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func FrontPage(c *gin.Context) {
  fmt.Fprintf(c.Writer, "Hello, %s!", c.Param("name"))
}
