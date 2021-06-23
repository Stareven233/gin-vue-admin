package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitVideoRouter(Router *gin.RouterGroup) {
	VideoRouter := Router.Group("video").Use(middleware.OperationRecord())
	{
		VideoRouter.POST("createVideo", v1.CreateVideo)   // 新建Video
		VideoRouter.DELETE("deleteVideo", v1.DeleteVideo) // 删除Video
		VideoRouter.DELETE("deleteVideoByIds", v1.DeleteVideoByIds) // 批量删除Video
		VideoRouter.PUT("updateVideo", v1.UpdateVideo)    // 更新Video
		VideoRouter.GET("findVideo", v1.FindVideo)        // 根据ID获取Video
		VideoRouter.GET("getVideoList", v1.GetVideoList)  // 获取Video列表
	}
}
