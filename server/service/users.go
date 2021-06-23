package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateUsers
//@description: 创建Users记录
//@param: users model.Users
//@return: err error

func CreateUsers(users model.Users) (err error) {
	err = global.GVA_DB.Create(&users).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUsers
//@description: 删除Users记录
//@param: users model.Users
//@return: err error

func DeleteUsers(users model.Users) (err error) {
	err = global.GVA_DB.Delete(&users).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUsersByIds
//@description: 批量删除Users记录
//@param: ids request.IdsReq
//@return: err error

func DeleteUsersByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Users{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateUsers
//@description: 更新Users记录
//@param: users *model.Users
//@return: err error

func UpdateUsers(users model.Users) (err error) {
	err = global.GVA_DB.Save(&users).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUsers
//@description: 根据id获取Users记录
//@param: id uint
//@return: err error, users model.Users

func GetUsers(id uint) (err error, users model.Users) {
	err = global.GVA_DB.Where("id = ?", id).First(&users).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUsersInfoList
//@description: 分页获取Users记录
//@param: info request.UsersSearch
//@return: err error, list interface{}, total int64

func GetUsersInfoList(info request.UsersSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Users{})
    var users []model.Users
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Username != "" {
        db = db.Where("`username` LIKE ?","%"+ info.Username+"%")
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&users).Error
	return err, users, total
}