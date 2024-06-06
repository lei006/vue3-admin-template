package controller

import (
	"net/http"
	"os"
	"yc-webreport-server/config"
	"yc-webreport-server/utils"

	"github.com/gin-gonic/gin"
	"github.com/sohaha/zlsgo/zlog"
)

type ReportPrintControl struct {
	BaseController
}

func (control *ReportPrintControl) Create(ctx *gin.Context) {

	file, err := ctx.FormFile("file")
	if err != nil {
		control.RetError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// 获取文件名
	fileName := file.Filename

	upload_path := config.ReportCfg.Api.UploadPrintPath
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

func (control *ReportPrintControl) DeleteOne(ctx *gin.Context) {

	id := ctx.Param("id")
	upload_path := config.ReportCfg.Api.UploadPrintPath
	file_pathname := upload_path + "/" + id + ".doc"

	err := os.Remove(file_pathname)
	if err != nil {
		zlog.Error("error:", file_pathname, err)
		control.RetError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	control.RetOK(ctx)
}

func (control *ReportPrintControl) DeleteMany(ctx *gin.Context) {
	var ids []string
	err := ctx.ShouldBindJSON(&ids)
	if err != nil {
		control.RetError(ctx, ERROR, err.Error())
		return
	}

	for i := 0; i < len(ids); i++ {
		id := ids[i]
		upload_path := config.ReportCfg.Api.UploadPrintPath
		file_pathname := upload_path + "/" + id + ".docx"

		err := os.Remove(file_pathname)
		if err != nil {
			zlog.Error("error:", file_pathname, err)
			control.RetError(ctx, http.StatusInternalServerError, err.Error())
			return
		}
	}

	control.RetOK(ctx)

}

func (control *ReportPrintControl) GetOne(ctx *gin.Context) {
	id := ctx.Param("id")

	upload_path := config.ReportCfg.Api.UploadPrintPath
	//file_pathname := upload_path + "/" + id + ".docx"
	file_pathname := upload_path + "/" + id

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

type PrintFileInfo struct {
	Id       string `json:"id"`
	FilePath string `json:"file_path"`
	Size     int64  `json:"size"`
}

func (control *ReportPrintControl) GetList(ctx *gin.Context) {

	// 根据目录取得文件列表
	upload_path := config.ReportCfg.Api.UploadPrintPath

	// 判断目录 upload_path 是否存在
	if _, err := os.Stat(upload_path); os.IsNotExist(err) {
		control.RetError(ctx, http.StatusInternalServerError, "目录不存在")
		return
	}

	// 从 upload_path 遍历文件夹，取得.doc文件列表
	docFiles, err := utils.GetFiles(upload_path, ".docx")
	if err != nil {
		control.RetError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	zlog.Debug(docFiles)

	list_info := []PrintFileInfo{}
	total := len(docFiles)

	for i := 0; i < total; i++ {
		// 取得文件名
		file_pathname := docFiles[i]

		fileInfo, err := os.Stat(file_pathname)
		if err != nil {
			zlog.Error("error:", file_pathname, err)
			control.RetError(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		item := PrintFileInfo{
			//Id:       utils.GetFileNameWithoutExt(fileInfo.Name()),
			Id:       fileInfo.Name(),
			FilePath: file_pathname,
			Size:     fileInfo.Size(),
		}

		list_info = append(list_info, item)
	}

	zlog.Debug("docFiles:", list_info)

	control.RetOkList(ctx, list_info)
}
