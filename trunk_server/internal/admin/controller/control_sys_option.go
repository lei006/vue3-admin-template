package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SysRecordControl struct {
	ControllerBase
}

var SysRecordControler = new(SysRecordControl)

func (control *SysRecordControl) Create(ctx *gin.Context) {

	RetOK(ctx)
}

func (control *SysRecordControl) DeleteMany(ctx *gin.Context) {
	var ids []string
	err := ctx.ShouldBindJSON(&ids)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	RetOK(ctx)

}

func (control *SysRecordControl) GetPage(ctx *gin.Context) {

	RetOK(ctx)
}
