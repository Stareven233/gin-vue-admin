package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitUsersRouter(Router *gin.RouterGroup) {
	UsersRouter := Router.Group("users").Use(middleware.OperationRecord())
	{
		UsersRouter.POST("createUsers", v1.CreateUsers)   // 新建Users
		UsersRouter.DELETE("deleteUsers", v1.DeleteUsers) // 删除Users
		UsersRouter.DELETE("deleteUsersByIds", v1.DeleteUsersByIds) // 批量删除Users
		UsersRouter.PUT("updateUsers", v1.UpdateUsers)    // 更新Users
		UsersRouter.GET("findUsers", v1.FindUsers)        // 根据ID获取Users
		UsersRouter.GET("getUsersList", v1.GetUsersList)  // 获取Users列表
	}
}
