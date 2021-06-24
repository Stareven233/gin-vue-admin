package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateComment
//@description: 创建Comment记录
//@param: comment model.Comment
//@return: err error

func CreateComment(comment model.Comment) (err error) {
	err = global.GVA_DB.Create(&comment).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteComment
//@description: 删除Comment记录
//@param: comment model.Comment
//@return: err error

func DeleteComment(comment model.Comment) (err error) {
	err = global.GVA_DB.Delete(&comment).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteCommentByIds
//@description: 批量删除Comment记录
//@param: ids request.IdsReq
//@return: err error

func DeleteCommentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Comment{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateComment
//@description: 更新Comment记录
//@param: comment *model.Comment
//@return: err error

func UpdateComment(comment model.Comment) (err error) {
	err = global.GVA_DB.Save(&comment).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetComment
//@description: 根据id获取Comment记录
//@param: id uint
//@return: err error, comment model.Comment

func GetComment(id uint) (err error, comment model.Comment) {
	err = global.GVA_DB.Where("id = ?", id).First(&comment).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetCommentInfoList
//@description: 分页获取Comment记录
//@param: info request.CommentSearch
//@return: err error, list interface{}, total int64

func GetCommentInfoList(info request.CommentSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Comment{})
    var comments []model.Comment
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&comments).Error
	return err, comments, total
}