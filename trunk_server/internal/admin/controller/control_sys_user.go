package controller

import (
	"net/http"
	"strconv"
	"vue3-admin-template/internal/db_model"
	"vue3-admin-template/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/sohaha/zlsgo/zlog"
)

type SysUserControl struct {
	ControllerBase
}

func (control *SysUserControl) Create(ctx *gin.Context) {
	var modelUser db_model.SysUser

	user_info := db_model.SysUser{}
	err := ctx.ShouldBindJSON(&user_info)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	zlog.Debugf("user_info: %+v \n", user_info)
	err = modelUser.Create(&user_info)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	RetData(ctx, user_info)
}

func (control *SysUserControl) DeleteOne(ctx *gin.Context) {

	id := ctx.Param("id")
	var modelUser db_model.SysUser

	err := modelUser.DeleteOne(id)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	RetOK(ctx)
}

func (control *SysUserControl) DeleteMany(ctx *gin.Context) {
	var ids []int
	err := ctx.ShouldBindJSON(&ids)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}
	var modelUser db_model.SysUser

	err = modelUser.DeleteMany(ids)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	RetOK(ctx)
}

func (control *SysUserControl) PutOne(ctx *gin.Context) {

	var reportItem db_model.SysUser
	err := ctx.ShouldBindJSON(&reportItem)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}
	verify := utils.Rules{
		"ID": {utils.NotEmpty()},
	}
	if err := utils.Verify(reportItem, verify); err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var modelUser db_model.SysUser

	if err := modelUser.UpdateOne(&reportItem); err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}
	RetOK(ctx)
}

func (control *SysUserControl) PatchOne(ctx *gin.Context) {
	id := ctx.Param("id")
	// 把id 转为 uint

	req := FieldDataRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}
	var modelUser db_model.SysUser

	err = modelUser.PatchOne(id, req.Field, req.Data)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}
	data, err := modelUser.GetOne(id)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}
	zlog.Debug("patch:", id, req, data)

	RetData(ctx, data)
}

func (control *SysUserControl) GetOne(ctx *gin.Context) {
	id := ctx.Param("id")

	var modelUser db_model.SysUser

	item, err := modelUser.GetOne(id)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	RetData(ctx, item)
}

func (control *SysUserControl) GetPage(ctx *gin.Context) {

	reqPageInfo := db_model.PageInfo{}
	// 从参数中获取分页信息
	page := ctx.Query("page")
	pageSize := ctx.Query("pageSize")
	keyword := ctx.Query("keyword")

	var err error

	// page 转为 int
	reqPageInfo.Page, err = strconv.Atoi(page)
	if err != nil {
		reqPageInfo.Page = 1
	}
	// pageSize 转为 int
	reqPageInfo.PageSize, err = strconv.Atoi(pageSize)
	if err != nil {
		reqPageInfo.PageSize = 10
	}

	// keyword 转为 string
	reqPageInfo.Keyword = keyword

	var modelUser db_model.SysUser
	items, total, err := modelUser.GetPage(reqPageInfo)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	RetPage(ctx, items, total)

}
