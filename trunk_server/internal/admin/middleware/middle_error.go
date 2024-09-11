package middleware

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"runtime/debug"
	"vue3-admin-template/internal/admin/controller"

	"github.com/gin-gonic/gin"
)

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
				var stack_msg string

				if stack {
					stack_msg = string(debug.Stack())
					println(stack_msg)
				}

				controller.RetErr(ctx, http.StatusInternalServerError, string(httpRequest)+"-服务内部异常: "+fmt.Sprintf("%v", err))
			}
		}()
		ctx.Next()
	}
}
