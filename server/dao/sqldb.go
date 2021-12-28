package dao

import "github.com/jinzhu/gorm"

//以下两个变量务必在项目里初始化
var (
	SqlDBWrite *gorm.DB
	SqlDBRead  *gorm.DB
)
