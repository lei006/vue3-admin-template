package middleware

import (
	"net/http"
	"vue3-admin-template/internal/admin/controller"
	"vue3-admin-template/internal/admin/model"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/lei006/zlog"
)

func LimitIp() gin.HandlerFunc {
	return func(c *gin.Context) {

		client_ip := c.ClientIP()

		var modelLimitIp model.SysLimitIp
		item, err := modelLimitIp.GetOneIp(client_ip)
		if err == gorm.ErrRecordNotFound {
			//没找到ip 直接放行
			c.Next()
			return
		}
		if err != nil {
			zlog.Error(err.Error())
			controller.RetErr(c, http.StatusInternalServerError, err.Error())
			return
		}

		if !item.IsLimit {

			zlog.Errorf("LimitIP: %+v \n", item)

			// 存在，但没有限制
			c.Next()
			return
		} else {
			zlog.Error("IP:", client_ip, "is limit")
			model.SysOptionLog("limit", item.Ip, c.Request.URL.String(), "", c.Request.UserAgent(), c.ClientIP())
			c.Abort()
		}
	}
}
