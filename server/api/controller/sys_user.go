package controller

import (
	"net/http"
	"os"
	"yc-webreport-server/api/utils"
	"yc-webreport-server/config"
	"yc-webreport-server/model"

	"github.com/gin-gonic/gin"
	"github.com/sohaha/zlsgo/zlog"
	"go.uber.org/zap"
)

var modelUser model.SysUser

type SysUserControl struct {
	BaseController
}

func (control *SysUserControl) Create(ctx *gin.Context) {

	user_info := model.SysUser{}
	err := ctx.ShouldBindJSON(&user_info)
	if err != nil {
		zlog.Debug("login:", zap.Error(err))
		control.RetErrorParam(ctx, "")
		return
	}

	err = modelUser.Create(&user_info)
	if err != nil {
		zlog.Debug("login:", zap.Error(err))
		control.RetErrorParam(ctx, "")
		return
	}

	control.RetOkData(ctx, user_info)
}

func (control *SysUserControl) DeleteOne(ctx *gin.Context) {

	id := ctx.Param("id")
	upload_path := config.ReportCfg.System.UploadPrintPath
	file_pathname := upload_path + "/" + id + ".doc"

	err := os.Remove(file_pathname)
	if err != nil {
		zlog.Error("error:", file_pathname, err)
		control.RetError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	control.RetOK(ctx)
}

func (control *SysUserControl) DeleteMany(ctx *gin.Context) {
	var ids []uint
	err := ctx.ShouldBindJSON(&ids)
	if err != nil {
		control.RetError(ctx, ERROR, err.Error())
		return
	}
	err = modelUser.DeleteMany(ids)
	if err != nil {
		control.RetError(ctx, ERROR, err.Error())
		return
	}

	control.RetOK(ctx)
}

func (control *SysUserControl) PutOne(ctx *gin.Context) {

	var reportItem model.SysUser
	err := ctx.ShouldBindJSON(&reportItem)
	if err != nil {
		control.RetError(ctx, ERROR, err.Error())
		return
	}
	verify := utils.Rules{
		"ID": {utils.NotEmpty()},
	}
	if err := utils.Verify(reportItem, verify); err != nil {
		control.RetError(ctx, ERROR, err.Error())
		return
	}

	if err := modelUser.UpdateOne(&reportItem); err != nil {
		zlog.Error("更新失败!", zap.Error(err))
		control.RetError(ctx, ERROR, err.Error())
		return
	}
	control.RetOK(ctx)
}

func (control *SysUserControl) PatchOne(ctx *gin.Context) {
	id := ctx.Param("id")
	// 把id 转为 uint

	req := PatchReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		control.RetError(ctx, ERROR, err.Error())
		return
	}

	err = modelUser.PatchOne(id, req.Field, req.Data)
	if err != nil {
		control.RetError(ctx, ERROR, err.Error())
		return
	}

	control.RetOK(ctx)
}

func (control *SysUserControl) GetOne(ctx *gin.Context) {
	id := ctx.Param("id")

	upload_path := config.ReportCfg.System.UploadPrintPath
	file_pathname := upload_path + "/" + id + ".doc"

	//ctx.
	//	w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(filePath))
	//w.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))
	if _, err := os.Stat(file_pathname); os.IsNotExist(err) {
		ctx.AbortWithStatus(http.StatusNotFound) // File not found
		return
	}
	ctx.File(file_pathname)
	//control.RetOK(ctx)
}

func (control *SysUserControl) GetList(ctx *gin.Context) {

	user_list, _, err := modelUser.GetList()
	if err != nil {
		control.RetError(ctx, ERROR, err.Error())
		return
	}

	control.ReturnList(ctx, user_list)
}

func (control *SysUserControl) GetPage(ctx *gin.Context) {

	control.RetOK(ctx)
}
