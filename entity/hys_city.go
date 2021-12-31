package entity

import (
	"database/sql"
	"time"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
)

//go:generate gormgen -structs HysCity -output hys_city_gen.go
type HysCity struct {
	ID         int    `gorm:"column:id;primary_key" json:"id" form:"id"`             //
	Parentid   int    `gorm:"column:parentid" json:"parentid" form:"parentid"`       //
	Cityname   string `gorm:"column:cityname" json:"cityname" form:"cityname"`       //
	Regiontype int    `gorm:"column:regiontype" json:"regiontype" form:"regiontype"` //
	Agencyid   int    `gorm:"column:agencyid" json:"agencyid" form:"agencyid"`       //

}

// TableName sets the insert table name for this struct type
func (h *HysCity) TableName() string {
	return "hys_city"
}
