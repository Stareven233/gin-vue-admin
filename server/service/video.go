package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateVideo
//@description: 创建Video记录
//@param: video model.Video
//@return: err error

func CreateVideo(video model.Video) (err error) {
	err = global.GVA_DB.Create(&video).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteVideo
//@description: 删除Video记录
//@param: video model.Video
//@return: err error

func DeleteVideo(video model.Video) (err error) {
	err = global.GVA_DB.Delete(&video).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteVideoByIds
//@description: 批量删除Video记录
//@param: ids request.IdsReq
//@return: err error

func DeleteVideoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Video{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateVideo
//@description: 更新Video记录
//@param: video *model.Video
//@return: err error

func UpdateVideo(video model.Video) (err error) {
	err = global.GVA_DB.Save(&video).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetVideo
//@description: 根据id获取Video记录
//@param: id uint
//@return: err error, video model.Video

func GetVideo(id uint) (err error, video model.Video) {
	err = global.GVA_DB.Where("id = ?", id).First(&video).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetVideoInfoList
//@description: 分页获取Video记录
//@param: info request.VideoSearch
//@return: err error, list interface{}, total int64

func GetVideoInfoList(info request.VideoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Video{})
    var videos []model.Video
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Title != "" {
        db = db.Where("`title` LIKE ?","%"+ info.Title+"%")
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&videos).Error
	return err, videos, total
}