import service from '@/utils/request'

// @Tags Video
// @Summary 创建Video
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Video true "创建Video"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /video/createVideo [post]
export const createVideo = (data) => {
  return service({
    url: '/video/createVideo',
    method: 'post',
    data
  })
}

// @Tags Video
// @Summary 删除Video
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Video true "删除Video"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /video/deleteVideo [delete]
export const deleteVideo = (data) => {
  return service({
    url: '/video/deleteVideo',
    method: 'delete',
    data
  })
}

// @Tags Video
// @Summary 删除Video
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Video"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /video/deleteVideo [delete]
export const deleteVideoByIds = (data) => {
  return service({
    url: '/video/deleteVideoByIds',
    method: 'delete',
    data
  })
}

// @Tags Video
// @Summary 更新Video
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Video true "更新Video"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /video/updateVideo [put]
export const updateVideo = (data) => {
  return service({
    url: '/video/updateVideo',
    method: 'put',
    data
  })
}

// @Tags Video
// @Summary 用id查询Video
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Video true "用id查询Video"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /video/findVideo [get]
export const findVideo = (params) => {
  return service({
    url: '/video/findVideo',
    method: 'get',
    params
  })
}

// @Tags Video
// @Summary 分页获取Video列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Video列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /video/getVideoList [get]
export const getVideoList = (params) => {
  return service({
    url: '/video/getVideoList',
    method: 'get',
    params
  })
}