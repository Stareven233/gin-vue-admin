package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitTestsRouter(Router *gin.RouterGroup) {
	TestsRouter := Router.Group("tests").Use(middleware.OperationRecord())
	{
		TestsRouter.GET("getUserByAssociation", v1.GetUserByAssociation)  // 测试association的使用
		TestsRouter.GET("getUserByPreload", v1.GetUserByPreload)  // 测试preload
		TestsRouter.Any("testCatByAssociation", v1.GetCatByAssociation)  // 新建两个类测试
	}
}
