package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateDanmaku
//@description: 创建Danmaku记录
//@param: danmaku model.Danmaku
//@return: err error

func CreateDanmaku(danmaku model.Danmaku) (err error) {
	err = global.GVA_DB.Create(&danmaku).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDanmaku
//@description: 删除Danmaku记录
//@param: danmaku model.Danmaku
//@return: err error

func DeleteDanmaku(danmaku model.Danmaku) (err error) {
	err = global.GVA_DB.Delete(&danmaku).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDanmakuByIds
//@description: 批量删除Danmaku记录
//@param: ids request.IdsReq
//@return: err error

func DeleteDanmakuByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Danmaku{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateDanmaku
//@description: 更新Danmaku记录
//@param: danmaku *model.Danmaku
//@return: err error

func UpdateDanmaku(danmaku model.Danmaku) (err error) {
	err = global.GVA_DB.Save(&danmaku).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDanmaku
//@description: 根据id获取Danmaku记录
//@param: id uint
//@return: err error, danmaku model.Danmaku

func GetDanmaku(id uint) (err error, danmaku model.Danmaku) {
	err = global.GVA_DB.Where("id = ?", id).First(&danmaku).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDanmakuInfoList
//@description: 分页获取Danmaku记录
//@param: info request.DanmakuSearch
//@return: err error, list interface{}, total int64

func GetDanmakuInfoList(info request.DanmakuSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Danmaku{})
    var danmakus []model.Danmaku
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&danmakus).Error
	return err, danmakus, total
}