package request

import "gin-vue-admin/model"

type VideoSearch struct{
    model.Video
    PageInfo
}