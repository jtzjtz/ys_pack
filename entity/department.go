package entity

import (
	"database/sql"
	"time"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
)

//go:generate gormgen -structs Department -output department_gen.go
type Department struct {
	DeptID       int    `gorm:"column:DeptId;primary_key" json:"DeptId" form:"DeptId"`       //
	DeptName     string `gorm:"column:DeptName" json:"DeptName" form:"DeptName"`             //
	DeptCode     string `gorm:"column:DeptCode" json:"DeptCode" form:"DeptCode"`             //
	DeptType     int    `gorm:"column:DeptType" json:"DeptType" form:"DeptType"`             // 1=公司 2=财务 3=行政 4=人力 5=销售 6=采购 7=生产部 0=虚拟部门
	ParentDeptID int    `gorm:"column:ParentDeptId" json:"ParentDeptId" form:"ParentDeptId"` //
	Status       int    `gorm:"column:Status" json:"Status" form:"Status"`                   //
	CreateTime   string `gorm:"column:CreateTime" json:"CreateTime" form:"CreateTime"`       //
	OrderID      int    `gorm:"column:OrderId" json:"OrderId" form:"OrderId"`                //

}

// TableName sets the insert table name for this struct type
func (d *Department) TableName() string {
	return "department"
}
