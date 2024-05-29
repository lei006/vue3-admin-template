package controller

import (
	"net/http"
	"time"
	"yc-webreport-server/api/utils"
	"yc-webreport-server/config"

	"github.com/gin-gonic/gin"
	"github.com/sohaha/zlsgo/zlog"

	"github.com/gofrs/uuid/v5"
)

type IdsReq struct {
	Ids []uint `json:"ids" form:"ids"`
}

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

// 基类，提供基本响应方法
type BaseController struct{}

var BaseControllerApp = new(BaseController)

type ResMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type JsonItemList struct {
	Items interface{} `json:"items"` //Data字段需要设置为interface类型以便接收任意数据
}

type JsonReturn struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"` //Data字段需要设置为interface类型以便接收任意数据
	Msg  string      `json:"message"`
	Now  int64       `json:"now"`
}

func (base *BaseController) _ret_data(ctx *gin.Context, code int, data interface{}, msg string) {
	var _ret JsonReturn
	_ret.Msg = msg
	_ret.Code = code
	_ret.Data = data
	_ret.Now = time.Now().Unix()
	ctx.JSON(http.StatusOK, _ret)
}

func (base *BaseController) ReturnList(ctx *gin.Context, data_list interface{}) {

	val := JsonItemList{
		Items: data_list,
	}

	base._ret_data(ctx, http.StatusOK, val, "success")
}

func (base *BaseController) RetOK(ctx *gin.Context) {
	base._ret_data(ctx, http.StatusOK, nil, "success")
}

func (base *BaseController) RetOkData(ctx *gin.Context, data interface{}) {
	base._ret_data(ctx, http.StatusOK, data, "success")
}

func (base *BaseController) RetError(ctx *gin.Context, code int, message string) {
	base._ret_data(ctx, code, nil, message)
}

func (base *BaseController) RetOkMessage(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, ResMsg{Code: SUCCESS, Msg: msg})
}

func (base *BaseController) GetClaims(c *gin.Context) (*utils.CustomClaims, error) {
	token := c.Request.Header.Get("x-token")
	j := utils.NewJWT(config.ReportCfg.JWT.SigningKey)
	claims, err := j.ParseToken(token)
	if err != nil {
		zlog.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
		return nil, err
	}
	return claims, err
}

// GetUserID 从Gin的Context中获取从jwt解析出来的用户ID
func (base *BaseController) GetUserID(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := base.GetClaims(c); err != nil {
			return 0
		} else {
			return cl.BaseClaims.ID
		}
	} else {
		waitUse := claims.(*utils.CustomClaims)
		return waitUse.BaseClaims.ID
	}
}

// GetUserUuid 从Gin的Context中获取从jwt解析出来的用户UUID
func (base *BaseController) GetUserUuid(c *gin.Context) uuid.UUID {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := base.GetClaims(c); err != nil {
			return uuid.UUID{}
		} else {
			return cl.UUID
		}

	} else {
		waitUse := claims.(*utils.CustomClaims)
		return waitUse.UUID
	}
}

// GetUserAuthorityId 从Gin的Context中获取从jwt解析出来的用户角色id
func (base *BaseController) GetUserAuthorityId(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := base.GetClaims(c); err != nil {
			return 0
		} else {
			return cl.AuthorityId
		}
	} else {
		waitUse := claims.(*utils.CustomClaims)
		return waitUse.AuthorityId
	}
}

// GetUserInfo 从Gin的Context中获取从jwt解析出来的用户角色id
func (base *BaseController) GetClaimsUserInfo(c *gin.Context) *utils.CustomClaims {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := base.GetClaims(c); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := claims.(*utils.CustomClaims)
		return waitUse
	}
}

// GetUserName 从Gin的Context中获取从jwt解析出来的用户名
func (base *BaseController) GetUserName(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := base.GetClaims(c); err != nil {
			return ""
		} else {
			return cl.Username
		}
	} else {
		waitUse := claims.(*utils.CustomClaims)
		return waitUse.Username
	}
}
