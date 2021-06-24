import service from '@/utils/request'

// @Tags Danmaku
// @Summary 创建Danmaku
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Danmaku true "创建Danmaku"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /danmaku/createDanmaku [post]
export const createDanmaku = (data) => {
  return service({
    url: '/danmaku/createDanmaku',
    method: 'post',
    data
  })
}

// @Tags Danmaku
// @Summary 删除Danmaku
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Danmaku true "删除Danmaku"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /danmaku/deleteDanmaku [delete]
export const deleteDanmaku = (data) => {
  return service({
    url: '/danmaku/deleteDanmaku',
    method: 'delete',
    data
  })
}

// @Tags Danmaku
// @Summary 删除Danmaku
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Danmaku"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /danmaku/deleteDanmaku [delete]
export const deleteDanmakuByIds = (data) => {
  return service({
    url: '/danmaku/deleteDanmakuByIds',
    method: 'delete',
    data
  })
}

// @Tags Danmaku
// @Summary 更新Danmaku
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Danmaku true "更新Danmaku"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /danmaku/updateDanmaku [put]
export const updateDanmaku = (data) => {
  return service({
    url: '/danmaku/updateDanmaku',
    method: 'put',
    data
  })
}

// @Tags Danmaku
// @Summary 用id查询Danmaku
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Danmaku true "用id查询Danmaku"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /danmaku/findDanmaku [get]
export const findDanmaku = (params) => {
  return service({
    url: '/danmaku/findDanmaku',
    method: 'get',
    params
  })
}

// @Tags Danmaku
// @Summary 分页获取Danmaku列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Danmaku列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /danmaku/getDanmakuList [get]
export const getDanmakuList = (params) => {
  return service({
    url: '/danmaku/getDanmakuList',
    method: 'get',
    params
  })
}