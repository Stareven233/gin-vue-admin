// 自动生成模板Users
package model

import (
	"gin-vue-admin/global"
	"time"
)

// 如果含有time.Time 请自行import time包
type Users struct {
      global.GVA_MODEL
      Username  string `json:"username" form:"username" gorm:"column:username;comment:;type:varchar(20);"`
      Password  string `json:"password" form:"password" gorm:"column:password;comment:;type:varchar(64);"`
      Email  string `json:"email" form:"email" gorm:"column:email;comment:;type:varchar(64);"`
      Avatar  string `json:"avatar" form:"avatar" gorm:"column:avatar;comment:;type:varchar(20);"`
      Nickname  string `json:"nickname" form:"nickname" gorm:"column:nickname;comment:;type:varchar(25);"`
      AboutMe  string `json:"aboutMe" form:"aboutMe" gorm:"column:about_me;comment:;type:text;"`
      MemberSince  time.Time `json:"memberSince" form:"memberSince" gorm:"column:member_since;comment:;type:datetime;"`
      LastSeen  time.Time `json:"lastSeen" form:"lastSeen" gorm:"column:last_seen;comment:;type:datetime;"`
      Permissions  int `json:"permissions" form:"permissions" gorm:"column:permissions;comment:;type:int;size:10;"`
      Coins  int `json:"coins" form:"coins" gorm:"column:coins;comment:;type:int;size:10;"`
}


func (Users) TableName() string {
  return "users"
}

