package utils

import (
	"github.com/gin-gonic/gin"
)

func GinRouterHome2Folder(engine *gin.Engine, folder_path string) {
	engine.StaticFile("/index.html", folder_path+"/index.html")
	engine.StaticFile("/favicon.ico", folder_path+"/favicon.ico")
	engine.StaticFile("/", folder_path+"/index.html")
	GinRouterView2Folder(engine, "/assets", folder_path+"/assets")
}

func GinRouterView2Folder(engine *gin.Engine, view_path string, folder_path string) {
	engine.Static(view_path, folder_path)
}
