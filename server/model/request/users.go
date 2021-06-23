package request

import "gin-vue-admin/model"

type UsersSearch struct{
    model.Users
    PageInfo
}