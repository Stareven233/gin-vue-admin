package model

import (
	"gin-vue-admin/global"
)

type Cat struct {
	global.GVA_MODEL
	Bells []Bell `gorm:"many2many:cat_bells;"`
}

func (Cat) TableName() string {
	return "cats"
}

type Bell struct {
	global.GVA_MODEL
	Name string
	//Cats []Cat `gorm:"many2many:cat_bells;"`
}

func (Bell) TableName() string {
	return "bells"
}
