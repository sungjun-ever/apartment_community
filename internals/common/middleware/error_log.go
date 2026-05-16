package middleware

import (
	"apart_community/internals/common/errUtils"
	"apart_community/internals/common/response"
	"errors"

	"github.com/gin-gonic/gin"
)

func ErrorLogMiddleWare(em *errUtils.ErrorManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			var appErr *errUtils.AppError

			em.Handle(c, err)

			if errors.As(err, &appErr) {
				response.Fail(c, appErr.Status, appErr.Code.String(), appErr.Code.GetMessage())
			} else {
				response.Fail(c, 500, errUtils.S001.String(), "INTERNAL_SERVER_ERROR")
			}
		}
	}
}
