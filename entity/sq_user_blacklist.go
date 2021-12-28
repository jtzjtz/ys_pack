package entity

import (
	"database/sql"
	"time"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
)

//go:generate gormgen -structs SqUserBlacklist -output sq_user_blacklist_gen.go
type SqUserBlacklist struct {
	ID         int    `gorm:"column:id;primary_key" json:"id" form:"id"`             //
	Userid     int    `gorm:"column:userid" json:"userid" form:"userid"`             //
	Username   string `gorm:"column:username" json:"username" form:"username"`       //
	Createtime string `gorm:"column:createtime" json:"createtime" form:"createtime"` //

}

// TableName sets the insert table name for this struct type
func (s *SqUserBlacklist) TableName() string {
	return "sq_user_blacklist"
}
