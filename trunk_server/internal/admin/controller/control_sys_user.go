package controller

import (
	"net/http"
	"vue3-admin-template/internal/admin/model"
	"vue3-admin-template/pkg/utils"

	"github.com/gin-gonic/gin"
)

var modelUser model.SysUser

type SysUserControl struct {
	ControllerBase
}

func (control *SysUserControl) Create(ctx *gin.Context) {

	user_info := model.SysUser{}
	err := ctx.ShouldBindJSON(&user_info)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}
	log.Debugf("user_info: %+v \n", user_info)
	err = modelUser.Create(&user_info)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	RetData(ctx, user_info)
}

func (control *SysUserControl) DeleteOne(ctx *gin.Context) {

	id := ctx.Param("id")

	err := modelUser.DeleteOne(id)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	RetOK(ctx)
}

func (control *SysUserControl) DeleteMany(ctx *gin.Context) {
	var ids []uint
	err := ctx.ShouldBindJSON(&ids)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}
	err = modelUser.DeleteMany(ids)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	RetOK(ctx)
}

func (control *SysUserControl) PutOne(ctx *gin.Context) {

	var reportItem model.SysUser
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
	log.Debug("patch:", id, req, data)

	RetData(ctx, data)
}

func (control *SysUserControl) GetOne(ctx *gin.Context) {
	id := ctx.Param("id")
	item, err := modelUser.GetOne(id)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	RetData(ctx, item)
}

func (control *SysUserControl) GetPage(ctx *gin.Context) {

	user_list, _, err := modelUser.GetPage()
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	RetList(ctx, user_list)

}
