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

// @Tags Video
// @Summary 创建Video
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Video true "创建Video"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /video/createVideo [post]
func CreateVideo(c *gin.Context) {
	var video model.Video
	_ = c.ShouldBindJSON(&video)
	if err := service.CreateVideo(video); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Video
// @Summary 删除Video
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Video true "删除Video"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /video/deleteVideo [delete]
func DeleteVideo(c *gin.Context) {
	var video model.Video
	_ = c.ShouldBindJSON(&video)
	if err := service.DeleteVideo(video); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Video
// @Summary 批量删除Video
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Video"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /video/deleteVideoByIds [delete]
func DeleteVideoByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteVideoByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Video
// @Summary 更新Video
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Video true "更新Video"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /video/updateVideo [put]
func UpdateVideo(c *gin.Context) {
	var video model.Video
	_ = c.ShouldBindJSON(&video)
	if err := service.UpdateVideo(video); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Video
// @Summary 用id查询Video
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Video true "用id查询Video"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /video/findVideo [get]
func FindVideo(c *gin.Context) {
	var video model.Video
	_ = c.ShouldBindQuery(&video)
	if err, revideo := service.GetVideo(video.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"revideo": revideo}, c)
	}
}

// @Tags Video
// @Summary 分页获取Video列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.VideoSearch true "分页获取Video列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /video/getVideoList [get]
func GetVideoList(c *gin.Context) {
	var pageInfo request.VideoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetVideoInfoList(pageInfo); err != nil {
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
