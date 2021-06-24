package request

import "gin-vue-admin/model"

type CommentSearch struct{
    model.Comment
    PageInfo
}