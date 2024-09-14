package middleware

import (
	"net/http"
	"vue3-admin-template/internal/admin/model"

	"github.com/gin-gonic/gin"
)

func RecordOption() gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.Request.Method != http.MethodGet {
			model.SysOptionLog("api", c.Request.Method, c.Request.URL.String(), "", c.Request.UserAgent(), c.ClientIP())
		} else {
			//model.SysOptionLog("api", c.Request.Method, c.Request.URL.String(), "", c.Request.UserAgent(), c.ClientIP())
		}

		c.Next()
	}
}
