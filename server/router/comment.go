package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitCommentRouter(Router *gin.RouterGroup) {
	CommentRouter := Router.Group("comment").Use(middleware.OperationRecord())
	{
		CommentRouter.POST("createComment", v1.CreateComment)   // 新建Comment
		CommentRouter.DELETE("deleteComment", v1.DeleteComment) // 删除Comment
		CommentRouter.DELETE("deleteCommentByIds", v1.DeleteCommentByIds) // 批量删除Comment
		CommentRouter.PUT("updateComment", v1.UpdateComment)    // 更新Comment
		CommentRouter.GET("findComment", v1.FindComment)        // 根据ID获取Comment
		CommentRouter.GET("getCommentList", v1.GetCommentList)  // 获取Comment列表
	}
}
