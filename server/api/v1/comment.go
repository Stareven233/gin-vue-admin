package v1

import (
	"gin-vue-admin/global"
    "gin-vue-admin/model"
    "gin-vue-admin/model/request"
    "gin-vue-admin/model/response"
    "gin-vue-admin/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// @Tags Comment
// @Summary 创建Comment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Comment true "创建Comment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /comment/createComment [post]
func CreateComment(c *gin.Context) {
	var comment model.Comment
	_ = c.ShouldBindJSON(&comment)
	if err := service.CreateComment(comment); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Comment
// @Summary 删除Comment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Comment true "删除Comment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /comment/deleteComment [delete]
func DeleteComment(c *gin.Context) {
	var comment model.Comment
	_ = c.ShouldBindJSON(&comment)
	if err := service.DeleteComment(comment); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Comment
// @Summary 批量删除Comment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Comment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /comment/deleteCommentByIds [delete]
func DeleteCommentByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteCommentByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Comment
// @Summary 更新Comment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Comment true "更新Comment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /comment/updateComment [put]
func UpdateComment(c *gin.Context) {
	var comment model.Comment
	_ = c.ShouldBindJSON(&comment)
	if err := service.UpdateComment(comment); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Comment
// @Summary 用id查询Comment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Comment true "用id查询Comment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /comment/findComment [get]
func FindComment(c *gin.Context) {
	var comment model.Comment
	_ = c.ShouldBindQuery(&comment)
	if err, recomment := service.GetComment(comment.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recomment": recomment}, c)
	}
}

// @Tags Comment
// @Summary 分页获取Comment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CommentSearch true "分页获取Comment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /comment/getCommentList [get]
func GetCommentList(c *gin.Context) {
	var pageInfo request.CommentSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetCommentInfoList(pageInfo); err != nil {
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
