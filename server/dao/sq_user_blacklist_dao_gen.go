package dao

import (
	"fmt"
	"github.com/jtzjtz/kit/database"
	"github.com/jtzjtz/ys_pack/entity"
	"gorm.io/gorm"
)

type SqUserBlacklistDAO struct {
}

func (instance *SqUserBlacklistDAO) Create(squserblacklist entity.SqUserBlacklist) (entity.SqUserBlacklist, error) {

	db := SqlDBWrite
	err := db.Create(&squserblacklist).Error
	return squserblacklist, err
}

//Update(map[string]interface{}{"name": "hello", "age": 18, "actived": false})
func (instance *SqUserBlacklistDAO) Update(data map[string]interface{}, conditions []database.SqlCondition) error {

	db := SqlDBWrite
	return instance.buildQuery(db, conditions).Model(&entity.SqUserBlacklist{}).Updates(data).Error

}

//Update(map[string]interface{}{"name": "hello", "age": 18, "actived": false})
func (instance *SqUserBlacklistDAO) UpdateBySql(data map[string]interface{}, sqlQuery string) error {

	db := SqlDBWrite
	return db.Where(sqlQuery).Model(&entity.SqUserBlacklist{}).Updates(data).Error

}
func (instance *SqUserBlacklistDAO) Save(squserblacklist entity.SqUserBlacklist) error {
	db := SqlDBWrite
	return db.Save(&squserblacklist).Error
}

func (instance *SqUserBlacklistDAO) Delete(conditions []database.SqlCondition) error {
	db := SqlDBWrite
	return instance.buildQuery(db, conditions).Model(&entity.SqUserBlacklist{}).Delete(entity.SqUserBlacklist{}).Error
}

func (instance *SqUserBlacklistDAO) DeleteBySql(sqlQuery string) error {
	db := SqlDBWrite
	return db.Where(sqlQuery).Model(&entity.SqUserBlacklist{}).Delete(entity.SqUserBlacklist{}).Error
}

func (instance *SqUserBlacklistDAO) GetByID(id int, isWriteDB bool, options ...database.SqlOptions) *entity.SqUserBlacklist {
	db := SqlDBRead
	if isWriteDB {
		db = SqlDBWrite
	}
	squserblacklist := &entity.SqUserBlacklist{}
	condition := []database.SqlCondition{}
	condition = append(condition, struct {
		QueryName string
		Predicate database.SqlPredicate
		Value     interface{}
	}{QueryName: "id", Value: id, Predicate: database.SqlEqualPredicate})
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
	if instance.buildQuery(db, condition).First(squserblacklist).RecordNotFound() {
		return nil
	}
	return squserblacklist
}

func (instance *SqUserBlacklistDAO) First(condition []database.SqlCondition, isWriteDB bool, options ...database.SqlOptions) *entity.SqUserBlacklist {
	db := SqlDBRead
	if isWriteDB {
		db = SqlDBWrite
	}
	squserblacklist := &entity.SqUserBlacklist{}
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
	if instance.buildQuery(db, condition).First(squserblacklist).RecordNotFound() {
		return nil
	}

	return squserblacklist
}

func (instance *SqUserBlacklistDAO) FirstBySql(sqlQuery string, isWriteDB bool, options ...database.SqlOptions) *entity.SqUserBlacklist {
	db := SqlDBRead
	if isWriteDB {
		db = SqlDBWrite
	}
	squserblacklist := &entity.SqUserBlacklist{}
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
	if db.Where(sqlQuery).First(squserblacklist).RecordNotFound() {
		return nil
	}

	return squserblacklist
}

func (instance *SqUserBlacklistDAO) Count(condition []database.SqlCondition, isWriteDB bool) int {
	var c int
	db := SqlDBRead

	if isWriteDB {
		db = SqlDBWrite
	}
	instance.buildQuery(db, condition).Model(&entity.SqUserBlacklist{}).Count(&c)

	return c
}

func (instance *SqUserBlacklistDAO) CountBySql(sqlQuery string, isWriteDB bool) int {
	var c int
	db := SqlDBRead

	if isWriteDB {
		db = SqlDBWrite
	}
	db.Where(sqlQuery).Model(&entity.SqUserBlacklist{}).Count(&c)

	return c
}

func (instance *SqUserBlacklistDAO) GetList(conditions []database.SqlCondition, isWriteDB bool, options ...database.SqlOptions) []entity.SqUserBlacklist {

	db := SqlDBRead
	squserblacklists := []entity.SqUserBlacklist{}
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
	instance.buildQuery(db, conditions).Find(&squserblacklists)
	return squserblacklists
}

func (instance *SqUserBlacklistDAO) GetListBySql(sqlQuery string, isWriteDB bool, options ...database.SqlOptions) []entity.SqUserBlacklist {

	db := SqlDBRead
	squserblacklists := []entity.SqUserBlacklist{}
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
	db.Where(sqlQuery).Find(&squserblacklists)
	return squserblacklists
}

func (instance *SqUserBlacklistDAO) GetPageList(conditions []database.SqlCondition, start int, limit int, orderBy map[string]string, isWriteDB bool, options ...database.SqlOptions) []entity.SqUserBlacklist {
	db := SqlDBRead
	if isWriteDB {
		db = SqlDBWrite
	}
	squserblacklists := []entity.SqUserBlacklist{}
	for orderKey, ascDes := range orderBy {
		db = db.Order(orderKey + " " + ascDes)
	}
	for _, option := range options {
		if option.SelectField != "" {
			db = db.Select(option.SelectField)
		}
	}
	instance.buildQuery(db, conditions).Offset(start).Limit(limit).Find(&squserblacklists)
	return squserblacklists
}

func (instance *SqUserBlacklistDAO) GetPageListBySql(sqlQuery string, start int, limit int, orderBy map[string]string, isWriteDB bool, options ...database.SqlOptions) []entity.SqUserBlacklist {
	db := SqlDBRead
	if isWriteDB {
		db = SqlDBWrite
	}
	squserblacklists := []entity.SqUserBlacklist{}
	for orderKey, ascDes := range orderBy {
		db = db.Order(orderKey + " " + ascDes)
	}
	for _, option := range options {
		if option.SelectField != "" {
			db = db.Select(option.SelectField)
		}
	}
	db.Where(sqlQuery).Offset(start).Limit(limit).Find(&squserblacklists)
	return squserblacklists
}

func (instance *SqUserBlacklistDAO) buildQuery(db *gorm.DB, condition []database.SqlCondition) *gorm.DB {

	for _, where := range condition {
		if where.Predicate == "" {
			where.Predicate = database.SqlEqualPredicate
		}
		db = db.Where(fmt.Sprintf("%v %v ", where.QueryName, where.Predicate), where.Value)
	}

	return db
}
