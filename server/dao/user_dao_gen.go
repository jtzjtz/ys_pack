package dao

import (
	"fmt"
	"github.com/jtzjtz/kit/database"
	"github.com/jtzjtz/ys_pack/entity"
	"gorm.io/gorm"
)

type UserDAO struct {
}

func (instance *UserDAO) Create(user entity.User) (entity.User, error) {

	db := SqlDBWrite
	err := db.Create(&user).Error
	return user, err
}

//Update(map[string]interface{}{"name": "hello", "age": 18, "actived": false})
func (instance *UserDAO) Update(data map[string]interface{}, conditions []database.SqlCondition) error {

	db := SqlDBWrite
	return instance.buildQuery(db, conditions).Model(&entity.User{}).Updates(data).Error

}

//Update(map[string]interface{}{"name": "hello", "age": 18, "actived": false})
func (instance *UserDAO) UpdateBySql(data map[string]interface{}, sqlQuery string) error {

	db := SqlDBWrite
	return db.Where(sqlQuery).Model(&entity.User{}).Updates(data).Error

}
func (instance *UserDAO) Save(user entity.User) error {
	db := SqlDBWrite
	return db.Save(&user).Error
}

func (instance *UserDAO) Delete(conditions []database.SqlCondition) error {
	db := SqlDBWrite
	return instance.buildQuery(db, conditions).Model(&entity.User{}).Delete(entity.User{}).Error
}

func (instance *UserDAO) DeleteBySql(sqlQuery string) error {
	db := SqlDBWrite
	return db.Where(sqlQuery).Model(&entity.User{}).Delete(entity.User{}).Error
}

func (instance *UserDAO) GetByUserID(userid int, isWriteDB bool, options ...database.SqlOptions) *entity.User {
	db := SqlDBRead
	if isWriteDB {
		db = SqlDBWrite
	}
	user := &entity.User{}
	condition := []database.SqlCondition{}
	condition = append(condition, struct {
		QueryName string
		Predicate database.SqlPredicate
		Value     interface{}
	}{QueryName: "UserId", Value: userid, Predicate: database.SqlEqualPredicate})
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
	if instance.buildQuery(db, condition).First(user).RecordNotFound() {
		return nil
	}
	return user
}

func (instance *UserDAO) First(condition []database.SqlCondition, isWriteDB bool, options ...database.SqlOptions) *entity.User {
	db := SqlDBRead
	if isWriteDB {
		db = SqlDBWrite
	}
	user := &entity.User{}
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
	if instance.buildQuery(db, condition).First(user).RecordNotFound() {
		return nil
	}

	return user
}

func (instance *UserDAO) FirstBySql(sqlQuery string, isWriteDB bool, options ...database.SqlOptions) *entity.User {
	db := SqlDBRead
	if isWriteDB {
		db = SqlDBWrite
	}
	user := &entity.User{}
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
	if db.Where(sqlQuery).First(user).RecordNotFound() {
		return nil
	}

	return user
}

func (instance *UserDAO) Count(condition []database.SqlCondition, isWriteDB bool) int {
	var c int
	db := SqlDBRead

	if isWriteDB {
		db = SqlDBWrite
	}
	instance.buildQuery(db, condition).Model(&entity.User{}).Count(&c)

	return c
}

func (instance *UserDAO) CountBySql(sqlQuery string, isWriteDB bool) int {
	var c int
	db := SqlDBRead

	if isWriteDB {
		db = SqlDBWrite
	}
	db.Where(sqlQuery).Model(&entity.User{}).Count(&c)

	return c
}

func (instance *UserDAO) GetList(conditions []database.SqlCondition, isWriteDB bool, options ...database.SqlOptions) []entity.User {

	db := SqlDBRead
	users := []entity.User{}
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
	instance.buildQuery(db, conditions).Find(&users)
	return users
}

func (instance *UserDAO) GetListBySql(sqlQuery string, isWriteDB bool, options ...database.SqlOptions) []entity.User {

	db := SqlDBRead
	users := []entity.User{}
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
	db.Where(sqlQuery).Find(&users)
	return users
}

func (instance *UserDAO) GetPageList(conditions []database.SqlCondition, start int, limit int, orderBy map[string]string, isWriteDB bool, options ...database.SqlOptions) []entity.User {
	db := SqlDBRead
	if isWriteDB {
		db = SqlDBWrite
	}
	users := []entity.User{}
	for orderKey, ascDes := range orderBy {
		db = db.Order(orderKey + " " + ascDes)
	}
	for _, option := range options {
		if option.SelectField != "" {
			db = db.Select(option.SelectField)
		}
	}
	instance.buildQuery(db, conditions).Offset(start).Limit(limit).Find(&users)
	return users
}

func (instance *UserDAO) GetPageListBySql(sqlQuery string, start int, limit int, orderBy map[string]string, isWriteDB bool, options ...database.SqlOptions) []entity.User {
	db := SqlDBRead
	if isWriteDB {
		db = SqlDBWrite
	}
	users := []entity.User{}
	for orderKey, ascDes := range orderBy {
		db = db.Order(orderKey + " " + ascDes)
	}
	for _, option := range options {
		if option.SelectField != "" {
			db = db.Select(option.SelectField)
		}
	}
	db.Where(sqlQuery).Offset(start).Limit(limit).Find(&users)
	return users
}

func (instance *UserDAO) buildQuery(db *gorm.DB, condition []database.SqlCondition) *gorm.DB {

	for _, where := range condition {
		if where.Predicate == "" {
			where.Predicate = database.SqlEqualPredicate
		}
		db = db.Where(fmt.Sprintf("%v %v ", where.QueryName, where.Predicate), where.Value)
	}

	return db
}
