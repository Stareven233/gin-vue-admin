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

// @Tags Danmaku
// @Summary 创建Danmaku
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Danmaku true "创建Danmaku"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /danmaku/createDanmaku [post]
func CreateDanmaku(c *gin.Context) {
	var danmaku model.Danmaku
	_ = c.ShouldBindJSON(&danmaku)
	if err := service.CreateDanmaku(danmaku); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Danmaku
// @Summary 删除Danmaku
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Danmaku true "删除Danmaku"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /danmaku/deleteDanmaku [delete]
func DeleteDanmaku(c *gin.Context) {
	var danmaku model.Danmaku
	_ = c.ShouldBindJSON(&danmaku)
	if err := service.DeleteDanmaku(danmaku); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Danmaku
// @Summary 批量删除Danmaku
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Danmaku"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /danmaku/deleteDanmakuByIds [delete]
func DeleteDanmakuByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteDanmakuByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Danmaku
// @Summary 更新Danmaku
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Danmaku true "更新Danmaku"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /danmaku/updateDanmaku [put]
func UpdateDanmaku(c *gin.Context) {
	var danmaku model.Danmaku
	_ = c.ShouldBindJSON(&danmaku)
	if err := service.UpdateDanmaku(danmaku); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Danmaku
// @Summary 用id查询Danmaku
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Danmaku true "用id查询Danmaku"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /danmaku/findDanmaku [get]
func FindDanmaku(c *gin.Context) {
	var danmaku model.Danmaku
	_ = c.ShouldBindQuery(&danmaku)
	if err, redanmaku := service.GetDanmaku(danmaku.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redanmaku": redanmaku}, c)
	}
}

// @Tags Danmaku
// @Summary 分页获取Danmaku列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DanmakuSearch true "分页获取Danmaku列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /danmaku/getDanmakuList [get]
func GetDanmakuList(c *gin.Context) {
	var pageInfo request.DanmakuSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetDanmakuInfoList(pageInfo); err != nil {
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
