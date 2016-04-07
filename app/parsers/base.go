package parsers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func parse(c *gin.Context, model interface{}) error {
	if c.Bind(model) != nil {
		return fmt.Errorf("Bad Request: bad params")
	}
	return nil
}
