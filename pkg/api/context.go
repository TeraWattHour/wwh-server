package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Context interface {
	Abort(status int, err FailedResponse) error
}

type context struct {
	GinContext *gin.Context
}

func (c context) Abort(status int, err FailedResponse) error {
	c.GinContext.AbortWithStatusJSON(status, err)
	return nil
}

func W(h func(c *gin.Context, C Context) error) func(c *gin.Context) {
	return func(c *gin.Context) {
		wrapper := context{GinContext: c}
		err := h(c, wrapper)
		if err != nil {
			if c.Writer.Size() <= 0 {
				c.AbortWithStatusJSON(500, FailedResponse{
					Code: InternalError,
				})
				InternalLogger.Println("Unhandled internal error thrown... ", err)
				return
			}
			fmt.Println("Unhandled error... ", err)
		}
	}
}
