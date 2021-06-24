package request

import "gin-vue-admin/model"

type DanmakuSearch struct{
    model.Danmaku
    PageInfo
}