package cmd

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func FrontPage(c *gin.Context) {
  fmt.Fprint(c.Writer, "Trireg2 Clockify Importer")
}
