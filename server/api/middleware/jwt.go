package middleware

import (
	"github.com/sohaha/zlsgo/zlog"

	"yc-webreport-server/api/controller"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {

	return func(c *gin.Context) {

		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中
		// 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("x-token")
		zlog.Debug("x-token", token)

		if token == "" {
			controller.BaseControllerApp.RetError(c, controller.ERROR, "未登录或非法访问")
			c.Abort()
			return
		}
		/*
			zlog.Debug("xxxxxxxxxxxxxxxx")

			j := utils.NewJWT(config.ReportCfg.JWT.SigningKey)
			// parseToken 解析token包含的信息
			claims, err := j.ParseToken(token)
			if err != nil {
				if errors.Is(err, utils.TokenExpired) {
					controller.BaseControllerApp.RetError(c, controller.ERROR, "授权已过期")
					c.Abort()
					return
				}
				controller.BaseControllerApp.RetError(c, controller.ERROR, err.Error())
				c.Abort()
				return
			}
			zlog.Debug("xxxxxxxxxxxxxxxx  11")

			// 已登录用户被管理员禁用 需要使该用户的jwt失效 此处比较消耗性能 如果需要 请自行打开
			// 用户被删除的逻辑 需要优化 此处比较消耗性能 如果需要 请自行打开

			//if user, err := userService.FindUserByUuid(claims.UUID.String()); err != nil || user.Enable == 2 {
			//	_ = jwtService.JsonInBlacklist(model.JwtBlacklist{Jwt: token})
			//	response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			//	c.Abort()
			//}
			if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
				dr, _ := utils.ParseDuration(config.ReportCfg.JWT.ExpiresTime)
				claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(dr))
				newToken, _ := j.CreateTokenByOldToken(token, *claims)
				newClaims, _ := j.ParseToken(newToken)
				c.Header("new-token", newToken)
				c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
				if config.ReportCfg.System.UseMultipoint {

				}
			}
			c.Set("claims", claims)
		*/
		c.Next()
	}
}
