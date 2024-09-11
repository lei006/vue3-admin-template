package controller

import (
	"net/http"
	"time"
	"vue3-admin-template/pkg/utils"

	"github.com/gin-gonic/gin"

	"github.com/lei006/zlog"
)

var log = zlog.New("admin-controller")

var JwtSigningKey = ""

type IdsReq struct {
	Ids []uint `json:"ids" form:"ids"`
}

type PageInfo struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   // 关键字
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
	ERROR              = 7
	SUCCESS            = 0
	ERROR_Unauthorized = 401
	ERROR_Forbidden    = 403
	ERROR_500          = 500
)

//401   （未授权）请求要求身份验证。对于需要登录的网页，服务器可能返回此响应。

// 基类，提供基本响应方法
type ControllerBase struct{}

var ControllerBaseApp = new(ControllerBase)

type ResMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type FieldDataRequest struct {
	Field string      `json:"field"`
	Data  interface{} `json:"data"`
}

type JsonItemList struct {
	List interface{} `json:"list"` //Data字段需要设置为interface类型以便接收任意数据
}

type JsonPageInfo struct {
	Items interface{} `json:"items"` //Data字段需要设置为interface类型以便接收任意数据
	Total int64       `json:"total"`
}

type JsonReturn struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"` //Data字段需要设置为interface类型以便接收任意数据
	Msg  string      `json:"message"`
	Now  int64       `json:"now"`
}

func _ret_data(ctx *gin.Context, code int, data interface{}, msg string) {
	var _ret JsonReturn
	_ret.Msg = msg
	_ret.Code = code
	_ret.Data = data
	_ret.Now = time.Now().Unix()
	ctx.JSON(http.StatusOK, _ret)
}

func RetPage(ctx *gin.Context, data_list interface{}, total int64) {

	val := JsonPageInfo{
		Items: data_list,
		Total: total,
	}

	_ret_data(ctx, http.StatusOK, val, "success")
}

func RetList(ctx *gin.Context, list interface{}) {

	val := JsonItemList{
		List: list,
	}

	_ret_data(ctx, http.StatusOK, val, "success")
}

func RetOK(ctx *gin.Context) {
	_ret_data(ctx, http.StatusOK, nil, "success")
}

func RetData(ctx *gin.Context, data interface{}) {
	_ret_data(ctx, http.StatusOK, data, "success")
}

func RetErr(ctx *gin.Context, code int, message string) {
	_ret_data(ctx, code, nil, message)
}

func RetMsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, ResMsg{Code: SUCCESS, Msg: msg})
}

func GetClaims(c *gin.Context) (*utils.CustomClaims, error) {
	token := c.Request.Header.Get("x-token")
	j := utils.NewJWT(JwtSigningKey)
	claims, err := j.ParseToken(token)
	if err != nil {
		return nil, err
	}
	return claims, err
}

// GetUserID 从Gin的Context中获取从jwt解析出来的用户ID
func GetUserID(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.BaseClaims.ID
		}
	} else {
		waitUse := claims.(*utils.CustomClaims)
		return waitUse.BaseClaims.ID
	}
}
