package v1

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Users
// @Summary 创建Users
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Users true "创建Users"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /users/createUsers [post]
func CreateUsers(c *gin.Context) {
	var users model.Users
	_ = c.ShouldBindJSON(&users)
	if err := service.CreateUsers(users); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Users
// @Summary 删除Users
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Users true "删除Users"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /users/deleteUsers [delete]
func DeleteUsers(c *gin.Context) {
	var users model.Users
	_ = c.ShouldBindJSON(&users)
	if err := service.DeleteUsers(users); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Users
// @Summary 批量删除Users
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Users"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /users/deleteUsersByIds [delete]
func DeleteUsersByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteUsersByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Users
// @Summary 更新Users
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Users true "更新Users"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /users/updateUsers [put]
func UpdateUsers(c *gin.Context) {
	var users model.Users
	_ = c.ShouldBindJSON(&users)
	if err := service.UpdateUsers(users); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Users
// @Summary 用id查询Users
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Users true "用id查询Users"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /users/findUsers [get]
func FindUsers(c *gin.Context) {
	var users model.Users
	_ = c.ShouldBindQuery(&users)
	// 如http://localhost:8888/users/findUsers?ID=2，ID居然是区分大小写的
	if err, reusers := service.GetUsers(users.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reusers": reusers}, c)
	}
}

// @Tags Users
// @Summary 分页获取Users列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UsersSearch true "分页获取Users列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /users/getUsersList [get]
func GetUsersList(c *gin.Context) {
	var pageInfo request.UsersSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetUsersInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}

func GetUserByAssociation(c *gin.Context) {
	var user model.Users
	_ = c.ShouldBindQuery(&user)
	if err, reuser := service.GetUsers(user.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		fmt.Printf("GetUserByAssociation user: %v\n", user)
		ass := global.GVA_DB.Model(&reuser).Association("VCollect")

		err2 := ass.Find(&reuser.VCollect)
		cnt := ass.Count()
		//err1, v := service.GetVideo(205)
		//err1 = ass.Append(&v)
		//cnt = ass.Count()
		//err2 := global.GVA_DB.Model(&reuser).Association("Videos").Find(&reuser.Videos)
		fmt.Println(err2, cnt)
		response.OkWithData(gin.H{"user": reuser}, c)
	}
}

func GetUserByPreload(c *gin.Context) {
	var user model.Users
	var user1 []model.Users
	_ = c.ShouldBindQuery(&user)
	fmt.Printf("GetUserByPreload user: %v\n", user)
	//err1 := global.GVA_DB.Preload("Danmaku").Preload("Videos").nl
	err1 := global.GVA_DB.Preload("VCollect").Preload("VLike").Preload("VCoin").Find(&user1)
	//where里限制的是最终reuser的id
	fmt.Println(err1)
	response.OkWithData(gin.H{"user": user1}, c)
}
