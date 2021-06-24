package v1

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/response"
	"github.com/gin-gonic/gin"
)

func GetCatByAssociation(c *gin.Context) {
	method := c.Request.Method
	println(method)
	var cat model.Cat
	_ = c.ShouldBindQuery(&cat)

	//bell := []model.Bell{{Name: "mmm"}}
	//global.GVA_DB.Create(bell)
	//cat := &model.Cat{Bells: bell}
	//global.GVA_DB.Create(cat)

	cat.ID = 0
	global.GVA_DB.First(&cat, 2)
	fmt.Printf("cat: %v\n", cat)
	ass := global.GVA_DB.Model(&cat).Association("Bells")

	cnt := ass.Count()
	var b1 []model.Bell
	err1 := ass.Find(&b1)
	// 完全按照示例的association不行
	var c1 []model.Cat
	err2 := global.GVA_DB.Preload("Bells").Find(&c1)
	// preload可以
	fmt.Println(cnt, err1, err2)
	response.OkWithData(gin.H{"user": cat}, c)
}
