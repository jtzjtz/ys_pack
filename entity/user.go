package entity

import (
	"database/sql"
	"time"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
)

//go:generate gormgen -structs User -output user_gen.go
type User struct {
	UserID      int    `gorm:"column:UserId;primary_key" json:"UserId" form:"UserId"`    //
	UserName    string `gorm:"column:UserName" json:"UserName" form:"UserName"`          //
	Pwd         string `gorm:"column:Pwd" json:"Pwd" form:"Pwd"`                         //
	RealName    string `gorm:"column:RealName" json:"RealName" form:"RealName"`          //
	Address     string `gorm:"column:Address" json:"Address" form:"Address"`             //
	Birthday    string `gorm:"column:Birthday" json:"Birthday" form:"Birthday"`          //
	Sex         int    `gorm:"column:Sex" json:"Sex" form:"Sex"`                         //
	Mobile      string `gorm:"column:Mobile" json:"Mobile" form:"Mobile"`                //
	OfficePhone string `gorm:"column:OfficePhone" json:"OfficePhone" form:"OfficePhone"` //
	HomePhone   string `gorm:"column:HomePhone" json:"HomePhone" form:"HomePhone"`       //
	Email       string `gorm:"column:Email" json:"Email" form:"Email"`                   //
	CreateTime  string `gorm:"column:CreateTime" json:"CreateTime" form:"CreateTime"`    //
	Status      int    `gorm:"column:Status" json:"Status" form:"Status"`                //
	PosIds      string `gorm:"column:PosIds" json:"PosIds" form:"PosIds"`                //
	Fax         string `gorm:"column:Fax" json:"Fax" form:"Fax"`                         //

}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "user"
}
