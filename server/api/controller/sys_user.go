package controller

import (
	"net/http"
	"os"
	"yc-webreport-server/config"

	"github.com/gin-gonic/gin"
	"github.com/sohaha/zlsgo/zlog"
)

type SysUserControl struct {
	BaseController
}

var SysUserControler = new(SysUserControl)

func (control *SysUserControl) Create(ctx *gin.Context) {

	file, err := ctx.FormFile("file")
	if err != nil {
		control.RetError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// 获取文件名
	fileName := file.Filename

	upload_path := config.ReportCfg.System.UploadPrintPath
	upload_filename_path := upload_path + "/" + fileName

	// 保存文件
	err = ctx.SaveUploadedFile(file, upload_filename_path)
	if err != nil {
		control.RetError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	zlog.Debug("upload_filename_path :", upload_filename_path)

	control.RetOK(ctx)
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
	var ids []string
	err := ctx.ShouldBindJSON(&ids)
	if err != nil {
		control.RetError(ctx, ERROR, err.Error())
		return
	}

	for i := 0; i < len(ids); i++ {
		id := ids[i]
		upload_path := config.ReportCfg.System.UploadPrintPath
		file_pathname := upload_path + "/" + id + ".doc"

		err := os.Remove(file_pathname)
		if err != nil {
			zlog.Error("error:", file_pathname, err)
			control.RetError(ctx, http.StatusInternalServerError, err.Error())
			return
		}
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

	control.RetOK(ctx)
}

func (control *SysUserControl) GetPage(ctx *gin.Context) {

	control.RetOK(ctx)
}
