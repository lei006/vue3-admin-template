package controller

import (
	"net/http"
	"time"
	"vue3-admin-template/pkg/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/rand"
)

var JwtSigningKey = ""

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
	ctx.Abort()
}

func RetPage(ctx *gin.Context, data_list interface{}, total int64) {

	val := JsonPageInfo{
		Items: data_list,
		Total: total,
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

func GetClaims(c *gin.Context) (*utils.CustomClaims, error) {
	token := c.Request.Header.Get("token")
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

// RandomString 生成指定长度和字符集的随机字符串
func RandomString(length int, numbers, letters, specials bool) string {
	rand.Seed(uint64(time.Now().UnixNano()))

	var charSet []rune

	if numbers {
		charSet = append(charSet, '0', '1', '2', '3', '4', '5', '6', '7', '8', '9')
	}
	if letters {
		charSet = append(charSet, 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z')
	}
	if specials {
		charSet = append(charSet, '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '+', '=', '{', '}', '[', ']', '|', '\\', ';', ':', ',', '.', '<', '>', '/', '?', '`', '~')
	}

	if len(charSet) == 0 {
		panic("At least one character type (numbers, letters, or specials) must be enabled.")
	}

	result := make([]rune, length)
	for i := range result {
		result[i] = charSet[rand.Intn(len(charSet))]
	}
	return string(result)
}
