// 自动生成模板Video
package model

import (
      "gin-vue-admin/global"
      "time"
)

type Video struct {
      global.GVA_MODEL
      Title  string `json:"title" form:"title" gorm:"column:title;comment:;type:varchar(32);"`
      File  string `json:"file" form:"file" gorm:"column:file;comment:;type:varchar(20);"`
      Date  time.Time `json:"date" form:"date" gorm:"column:date;comment:;type:datetime;"`
      AuthorId  uint `json:"authorId" form:"authorId" gorm:"column:author_id;comment:;type:int;size:10;"`
      Type  int `json:"type" form:"type" gorm:"column:type;comment:;type:smallint;"`
      Face  string `json:"face" form:"face" gorm:"column:face;comment:;type:varchar(20);"`
      Desc  string `json:"desc" form:"desc" gorm:"column:desc;comment:;type:varchar(64);"`
      Coin  int `json:"coin" form:"coin" gorm:"column:coin;comment:;type:int;size:10;"`
      Collect  int `json:"collect" form:"collect" gorm:"column:collect;comment:;type:int;size:10;"`
      Like  int `json:"like" form:"like" gorm:"column:like;comment:;type:int;size:10;"`
      Status  int `json:"status" form:"status" gorm:"column:status;comment:;type:smallint;"`
      Inspector  int `json:"inspector" form:"inspector" gorm:"column:inspector;comment:;type:int;size:10;"`
      Danmaku []Danmaku `json:"danmaku" gorm:"foreignKey:VideoId"`
}

func (Video) TableName() string {
      return "videos"
}
