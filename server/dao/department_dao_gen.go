package dao

import (
	"fmt"
	"github.com/jtzjtz/kit/database"
	"github.com/jtzjtz/ys_pack/entity"
	"gorm.io/gorm"
)

type DepartmentDAO struct {
}

func (instance *DepartmentDAO) Create(department entity.Department) (entity.Department, error) {

	db := SqlDBWrite
	err := db.Create(&department).Error
	return department, err
}

//Update(map[string]interface{}{"name": "hello", "age": 18, "actived": false})
func (instance *DepartmentDAO) Update(data map[string]interface{}, conditions []database.SqlCondition) error {

	db := SqlDBWrite
	return instance.buildQuery(db, conditions).Model(&entity.Department{}).Updates(data).Error

}

//Update(map[string]interface{}{"name": "hello", "age": 18, "actived": false})
func (instance *DepartmentDAO) UpdateBySql(data map[string]interface{}, sqlQuery string) error {

	db := SqlDBWrite
	return db.Where(sqlQuery).Model(&entity.Department{}).Updates(data).Error

}
func (instance *DepartmentDAO) Save(department entity.Department) error {
	db := SqlDBWrite
	return db.Save(&department).Error
}

func (instance *DepartmentDAO) Delete(conditions []database.SqlCondition) error {
	db := SqlDBWrite
	return instance.buildQuery(db, conditions).Model(&entity.Department{}).Delete(entity.Department{}).Error
}

func (instance *DepartmentDAO) DeleteBySql(sqlQuery string) error {
	db := SqlDBWrite
	return db.Where(sqlQuery).Model(&entity.Department{}).Delete(entity.Department{}).Error
}

func (instance *DepartmentDAO) GetByDeptID(deptid int, isWriteDB bool, options ...database.SqlOptions) *entity.Department {
	db := SqlDBRead
	if isWriteDB {
		db = SqlDBWrite
	}
	department := &entity.Department{}
	condition := []database.SqlCondition{}
	condition = append(condition, struct {
		QueryName string
		Predicate database.SqlPredicate
		Value     interface{}
	}{QueryName: "DeptId", Value: deptid, Predicate: database.SqlEqualPredicate})
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
	if instance.buildQuery(db, condition).First(department).RecordNotFound() {
		return nil
	}
	return department
}

func (instance *DepartmentDAO) First(condition []database.SqlCondition, isWriteDB bool, options ...database.SqlOptions) *entity.Department {
	db := SqlDBRead
	if isWriteDB {
		db = SqlDBWrite
	}
	department := &entity.Department{}
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
	if instance.buildQuery(db, condition).First(department).RecordNotFound() {
		return nil
	}

	return department
}

func (instance *DepartmentDAO) FirstBySql(sqlQuery string, isWriteDB bool, options ...database.SqlOptions) *entity.Department {
	db := SqlDBRead
	if isWriteDB {
		db = SqlDBWrite
	}
	department := &entity.Department{}
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
	if db.Where(sqlQuery).First(department).RecordNotFound() {
		return nil
	}

	return department
}

func (instance *DepartmentDAO) Count(condition []database.SqlCondition, isWriteDB bool) int {
	var c int
	db := SqlDBRead

	if isWriteDB {
		db = SqlDBWrite
	}
	instance.buildQuery(db, condition).Model(&entity.Department{}).Count(&c)

	return c
}

func (instance *DepartmentDAO) CountBySql(sqlQuery string, isWriteDB bool) int {
	var c int
	db := SqlDBRead

	if isWriteDB {
		db = SqlDBWrite
	}
	db.Where(sqlQuery).Model(&entity.Department{}).Count(&c)

	return c
}

func (instance *DepartmentDAO) GetList(conditions []database.SqlCondition, isWriteDB bool, options ...database.SqlOptions) []entity.Department {

	db := SqlDBRead
	departments := []entity.Department{}
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
	instance.buildQuery(db, conditions).Find(&departments)
	return departments
}

func (instance *DepartmentDAO) GetListBySql(sqlQuery string, isWriteDB bool, options ...database.SqlOptions) []entity.Department {

	db := SqlDBRead
	departments := []entity.Department{}
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
	db.Where(sqlQuery).Find(&departments)
	return departments
}

func (instance *DepartmentDAO) GetPageList(conditions []database.SqlCondition, start int, limit int, orderBy map[string]string, isWriteDB bool, options ...database.SqlOptions) []entity.Department {
	db := SqlDBRead
	if isWriteDB {
		db = SqlDBWrite
	}
	departments := []entity.Department{}
	for orderKey, ascDes := range orderBy {
		db = db.Order(orderKey + " " + ascDes)
	}
	for _, option := range options {
		if option.SelectField != "" {
			db = db.Select(option.SelectField)
		}
	}
	instance.buildQuery(db, conditions).Offset(start).Limit(limit).Find(&departments)
	return departments
}

func (instance *DepartmentDAO) GetPageListBySql(sqlQuery string, start int, limit int, orderBy map[string]string, isWriteDB bool, options ...database.SqlOptions) []entity.Department {
	db := SqlDBRead
	if isWriteDB {
		db = SqlDBWrite
	}
	departments := []entity.Department{}
	for orderKey, ascDes := range orderBy {
		db = db.Order(orderKey + " " + ascDes)
	}
	for _, option := range options {
		if option.SelectField != "" {
			db = db.Select(option.SelectField)
		}
	}
	db.Where(sqlQuery).Offset(start).Limit(limit).Find(&departments)
	return departments
}

func (instance *DepartmentDAO) buildQuery(db *gorm.DB, condition []database.SqlCondition) *gorm.DB {

	for _, where := range condition {
		if where.Predicate == "" {
			where.Predicate = database.SqlEqualPredicate
		}
		db = db.Where(fmt.Sprintf("%v %v ", where.QueryName, where.Predicate), where.Value)
	}

	return db
}
