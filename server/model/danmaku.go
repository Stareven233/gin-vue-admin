// 自动生成模板Danmaku
package model

import (
      "gin-vue-admin/global"
      "time"
)

type Danmaku struct {
      global.GVA_MODEL
      Date  time.Time `json:"date" form:"date" gorm:"column:date;comment:;type:datetime;"`
      Time  float64 `json:"time" form:"time" gorm:"column:time;comment:弹幕出现的视频时刻;type:float;"`
      Text  string `json:"text" form:"text" gorm:"column:text;comment:"`
      Color  int `json:"color" form:"color" gorm:"column:color;comment:;type:int;size:10;"`
      Type  int `json:"type" form:"type" gorm:"column:type;comment:;type:smallint;"`
      AuthorId  uint `json:"authorId" form:"authorId" gorm:"column:author_id;comment:;type:int;size:10;"`
      VideoId  uint `json:"videoId" form:"videoId" gorm:"column:video_id;comment:;type:int;size:10;"`
}
// gorm在struct tag里设置的size总是要比数据库中实际的大，非常恶心，如上面的Color，mysql里长度是11
// 实测去掉comment:以8位十进制数表示rgb颜色;的文本就没事了，纯nt


func (Danmaku) TableName() string {
  return "danmaku"
}

