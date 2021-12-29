package dao

import (
	"fmt"
	"github.com/jtzjtz/kit/database"
	"github.com/jtzjtz/ys_pack/entity"
	"gorm.io/gorm"
)

type WorkflowDAO struct {
}

func (instance *WorkflowDAO) Create(workflow entity.Workflow) (entity.Workflow, error) {

	db := SqlDBWrite
	err := db.Create(&workflow).Error
	return workflow, err
}

//Update(map[string]interface{}{"name": "hello", "age": 18, "actived": false})
func (instance *WorkflowDAO) Update(data map[string]interface{}, conditions []database.SqlCondition) error {

	db := SqlDBWrite
	return instance.buildQuery(db, conditions).Model(&entity.Workflow{}).Updates(data).Error

}

//Update(map[string]interface{}{"name": "hello", "age": 18, "actived": false})
func (instance *WorkflowDAO) UpdateBySql(data map[string]interface{}, sqlQuery string) error {

	db := SqlDBWrite
	return db.Where(sqlQuery).Model(&entity.Workflow{}).Updates(data).Error

}
func (instance *WorkflowDAO) Save(workflow entity.Workflow) error {
	db := SqlDBWrite
	return db.Save(&workflow).Error
}

func (instance *WorkflowDAO) Delete(conditions []database.SqlCondition) error {
	db := SqlDBWrite
	return instance.buildQuery(db, conditions).Model(&entity.Workflow{}).Delete(entity.Workflow{}).Error
}

func (instance *WorkflowDAO) DeleteBySql(sqlQuery string) error {
	db := SqlDBWrite
	return db.Where(sqlQuery).Model(&entity.Workflow{}).Delete(entity.Workflow{}).Error
}

func (instance *WorkflowDAO) GetByWorkFlowID(workflowid int, isWriteDB bool, options ...database.SqlOptions) *entity.Workflow {
	db := SqlDBRead
	if isWriteDB {
		db = SqlDBWrite
	}
	workflow := &entity.Workflow{}
	condition := []database.SqlCondition{}
	condition = append(condition, struct {
		QueryName string
		Predicate database.SqlPredicate
		Value     interface{}
	}{QueryName: "WorkFlowId", Value: workflowid, Predicate: database.SqlEqualPredicate})
	for _, option := range options {
		if option.SelectField != "" {
			db = db.Select(option.SelectField)
		}
		if option.OrderBy != nil {
			for orderKey, ascDes := range option.OrderBy {
				db = db.Order(orderKey + " " + ascDes)
			}
		}
	}
	if instance.buildQuery(db, condition).First(workflow).RecordNotFound() {
		return nil
	}
	return workflow
}

func (instance *WorkflowDAO) First(condition []database.SqlCondition, isWriteDB bool, options ...database.SqlOptions) *entity.Workflow {
	db := SqlDBRead
	if isWriteDB {
		db = SqlDBWrite
	}
	workflow := &entity.Workflow{}
	for _, option := range options {
		if option.SelectField != "" {
			db = db.Select(option.SelectField)
		}
		if option.OrderBy != nil {
			for orderKey, ascDes := range option.OrderBy {
				db = db.Order(orderKey + " " + ascDes)
			}
		}
	}
	if instance.buildQuery(db, condition).First(workflow).RecordNotFound() {
		return nil
	}

	return workflow
}

func (instance *WorkflowDAO) FirstBySql(sqlQuery string, isWriteDB bool, options ...database.SqlOptions) *entity.Workflow {
	db := SqlDBRead
	if isWriteDB {
		db = SqlDBWrite
	}
	workflow := &entity.Workflow{}
	for _, option := range options {
		if option.SelectField != "" {
			db = db.Select(option.SelectField)
		}
		if option.OrderBy != nil {
			for orderKey, ascDes := range option.OrderBy {
				db = db.Order(orderKey + " " + ascDes)
			}
		}
	}
	if db.Where(sqlQuery).First(workflow).RecordNotFound() {
		return nil
	}

	return workflow
}

func (instance *WorkflowDAO) Count(condition []database.SqlCondition, isWriteDB bool) int {
	var c int
	db := SqlDBRead

	if isWriteDB {
		db = SqlDBWrite
	}
	instance.buildQuery(db, condition).Model(&entity.Workflow{}).Count(&c)

	return c
}

func (instance *WorkflowDAO) CountBySql(sqlQuery string, isWriteDB bool) int {
	var c int
	db := SqlDBRead

	if isWriteDB {
		db = SqlDBWrite
	}
	db.Where(sqlQuery).Model(&entity.Workflow{}).Count(&c)

	return c
}

func (instance *WorkflowDAO) GetList(conditions []database.SqlCondition, isWriteDB bool, options ...database.SqlOptions) []entity.Workflow {

	db := SqlDBRead
	workflows := []entity.Workflow{}
	if isWriteDB {
		db = SqlDBWrite
	}
	for _, option := range options {
		if option.SelectField != "" {
			db = db.Select(option.SelectField)
		}
		if option.OrderBy != nil {
			for orderKey, ascDes := range option.OrderBy {
				db = db.Order(orderKey + " " + ascDes)
			}
		}
	}
	instance.buildQuery(db, conditions).Find(&workflows)
	return workflows
}

func (instance *WorkflowDAO) GetListBySql(sqlQuery string, isWriteDB bool, options ...database.SqlOptions) []entity.Workflow {

	db := SqlDBRead
	workflows := []entity.Workflow{}
	if isWriteDB {
		db = SqlDBWrite
	}
	for _, option := range options {
		if option.SelectField != "" {
			db = db.Select(option.SelectField)
		}
		if option.OrderBy != nil {
			for orderKey, ascDes := range option.OrderBy {
				db = db.Order(orderKey + " " + ascDes)
			}
		}
	}
	db.Where(sqlQuery).Find(&workflows)
	return workflows
}

func (instance *WorkflowDAO) GetPageList(conditions []database.SqlCondition, start int, limit int, orderBy map[string]string, isWriteDB bool, options ...database.SqlOptions) []entity.Workflow {
	db := SqlDBRead
	if isWriteDB {
		db = SqlDBWrite
	}
	workflows := []entity.Workflow{}
	for orderKey, ascDes := range orderBy {
		db = db.Order(orderKey + " " + ascDes)
	}
	for _, option := range options {
		if option.SelectField != "" {
			db = db.Select(option.SelectField)
		}
	}
	instance.buildQuery(db, conditions).Offset(start).Limit(limit).Find(&workflows)
	return workflows
}

func (instance *WorkflowDAO) GetPageListBySql(sqlQuery string, start int, limit int, orderBy map[string]string, isWriteDB bool, options ...database.SqlOptions) []entity.Workflow {
	db := SqlDBRead
	if isWriteDB {
		db = SqlDBWrite
	}
	workflows := []entity.Workflow{}
	for orderKey, ascDes := range orderBy {
		db = db.Order(orderKey + " " + ascDes)
	}
	for _, option := range options {
		if option.SelectField != "" {
			db = db.Select(option.SelectField)
		}
	}
	db.Where(sqlQuery).Offset(start).Limit(limit).Find(&workflows)
	return workflows
}

func (instance *WorkflowDAO) buildQuery(db *gorm.DB, condition []database.SqlCondition) *gorm.DB {

	for _, where := range condition {
		if where.Predicate == "" {
			where.Predicate = database.SqlEqualPredicate
		}
		db = db.Where(fmt.Sprintf("%v %v ", where.QueryName, where.Predicate), where.Value)
	}

	return db
}
