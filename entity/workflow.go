package entity

import (
	"database/sql"
	"time"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
)

//go:generate gormgen -structs Workflow -output workflow_gen.go
type Workflow struct {
	WorkFlowID int    `gorm:"column:WorkFlowId;primary_key" json:"WorkFlowId" form:"WorkFlowId"` //
	FlowName   string `gorm:"column:FlowName" json:"FlowName" form:"FlowName"`                   //
	FlowType   int    `gorm:"column:FlowType" json:"FlowType" form:"FlowType"`                   // 流程类型
	ManagerID  int    `gorm:"column:ManagerId" json:"ManagerId" form:"ManagerId"`                //
	Steps      string `gorm:"column:Steps" json:"Steps" form:"Steps"`                            // 流程步骤
	Forms      string `gorm:"column:Forms" json:"Forms" form:"Forms"`                            // 表单
	Events     string `gorm:"column:Events" json:"Events" form:"Events"`                         //
	FlowGUID   string `gorm:"column:FlowGuid" json:"FlowGuid" form:"FlowGuid"`                   // 流程模板guid
	Status     int    `gorm:"column:Status" json:"Status" form:"Status"`                         //
	CreateTime string `gorm:"column:CreateTime" json:"CreateTime" form:"CreateTime"`             //

}

// TableName sets the insert table name for this struct type
func (w *Workflow) TableName() string {
	return "workflow"
}
