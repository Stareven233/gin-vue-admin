// 自动生成模板Comment
package model

import (
	"gin-vue-admin/global"
	"time"
)

type Comment struct {
      global.GVA_MODEL
      Body  string `json:"body" form:"body" gorm:"column:body;comment:;type:text;"`
      Date  time.Time `json:"date" form:"date" gorm:"column:date;comment:;type:datetime;"`
      Disabled  *bool `json:"disabled" form:"disabled" gorm:"column:disabled;comment:;type:tinyint"`
      AuthorId  int `json:"authorId" form:"authorId" gorm:"column:author_id;comment:;type:int;size:10;"`
      VideoId  int `json:"videoId" form:"videoId" gorm:"column:video_id;comment:;type:int;size:10;"`
      TipUser  int `json:"tipUser" form:"tipUser" gorm:"column:tip_user;comment:;type:int;size:10;"`
      Tipped  *bool `json:"tipped" form:"tipped" gorm:"column:tipped;comment:;type:tinyint"`
}


func (Comment) TableName() string {
  return "comments"
}

