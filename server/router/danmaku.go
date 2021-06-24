package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDanmakuRouter(Router *gin.RouterGroup) {
	DanmakuRouter := Router.Group("danmaku").Use(middleware.OperationRecord())
	{
		DanmakuRouter.POST("createDanmaku", v1.CreateDanmaku)   // 新建Danmaku
		DanmakuRouter.DELETE("deleteDanmaku", v1.DeleteDanmaku) // 删除Danmaku
		DanmakuRouter.DELETE("deleteDanmakuByIds", v1.DeleteDanmakuByIds) // 批量删除Danmaku
		DanmakuRouter.PUT("updateDanmaku", v1.UpdateDanmaku)    // 更新Danmaku
		DanmakuRouter.GET("findDanmaku", v1.FindDanmaku)        // 根据ID获取Danmaku
		DanmakuRouter.GET("getDanmakuList", v1.GetDanmakuList)  // 获取Danmaku列表
	}
}
