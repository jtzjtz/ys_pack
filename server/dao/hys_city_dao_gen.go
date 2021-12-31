package dao

import (
	"fmt"
	"github.com/jtzjtz/kit/database"
	"github.com/jtzjtz/ys_pack/entity"
	"gorm.io/gorm"
)

type HysCityDAO struct {
}

func (instance *HysCityDAO) Create(hyscity entity.HysCity) (entity.HysCity, error) {

	db := SqlDBWrite
	err := db.Create(&hyscity).Error
	return hyscity, err
}

//Update(map[string]interface{}{"name": "hello", "age": 18, "actived": false})
func (instance *HysCityDAO) Update(data map[string]interface{}, conditions []database.SqlCondition) error {

	db := SqlDBWrite
	return instance.buildQuery(db, conditions).Model(&entity.HysCity{}).Updates(data).Error

}

//Update(map[string]interface{}{"name": "hello", "age": 18, "actived": false})
func (instance *HysCityDAO) UpdateBySql(data map[string]interface{}, sqlQuery string) error {

	db := SqlDBWrite
	return db.Where(sqlQuery).Model(&entity.HysCity{}).Updates(data).Error

}
func (instance *HysCityDAO) Save(hyscity entity.HysCity) error {
	db := SqlDBWrite
	return db.Save(&hyscity).Error
}

func (instance *HysCityDAO) Delete(conditions []database.SqlCondition) error {
	db := SqlDBWrite
	return instance.buildQuery(db, conditions).Model(&entity.HysCity{}).Delete(entity.HysCity{}).Error
}

func (instance *HysCityDAO) DeleteBySql(sqlQuery string) error {
	db := SqlDBWrite
	return db.Where(sqlQuery).Model(&entity.HysCity{}).Delete(entity.HysCity{}).Error
}

func (instance *HysCityDAO) GetByID(id int, isWriteDB bool, options ...database.SqlOptions) *entity.HysCity {
	db := SqlDBRead
	if isWriteDB {
		db = SqlDBWrite
	}
	hyscity := &entity.HysCity{}
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
	if instance.buildQuery(db, condition).First(hyscity).RowsAffected == 0 {
		return nil
	}
	return hyscity
}

func (instance *HysCityDAO) First(condition []database.SqlCondition, isWriteDB bool, options ...database.SqlOptions) *entity.HysCity {
	db := SqlDBRead
	if isWriteDB {
		db = SqlDBWrite
	}
	hyscity := &entity.HysCity{}
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
	if instance.buildQuery(db, condition).First(hyscity).RowsAffected == 0 {
		return nil
	}

	return hyscity
}

func (instance *HysCityDAO) FirstBySql(sqlQuery string, isWriteDB bool, options ...database.SqlOptions) *entity.HysCity {
	db := SqlDBRead
	if isWriteDB {
		db = SqlDBWrite
	}
	hyscity := &entity.HysCity{}
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
	if db.Where(sqlQuery).First(hyscity).RowsAffected == 0 {
		return nil
	}

	return hyscity
}

func (instance *HysCityDAO) Count(condition []database.SqlCondition, isWriteDB bool) int {
	var c int64
	db := SqlDBRead

	if isWriteDB {
		db = SqlDBWrite
	}
	instance.buildQuery(db, condition).Model(&entity.HysCity{}).Count(&c)

	return int(c)
}

func (instance *HysCityDAO) CountBySql(sqlQuery string, isWriteDB bool) int {
	var c int64
	db := SqlDBRead

	if isWriteDB {
		db = SqlDBWrite
	}
	db.Where(sqlQuery).Model(&entity.HysCity{}).Count(&c)

	return int(c)
}

func (instance *HysCityDAO) GetList(conditions []database.SqlCondition, isWriteDB bool, options ...database.SqlOptions) []entity.HysCity {

	db := SqlDBRead
	hyscitys := []entity.HysCity{}
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
	instance.buildQuery(db, conditions).Find(&hyscitys)
	return hyscitys
}

func (instance *HysCityDAO) GetListBySql(sqlQuery string, isWriteDB bool, options ...database.SqlOptions) []entity.HysCity {

	db := SqlDBRead
	hyscitys := []entity.HysCity{}
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
	db.Where(sqlQuery).Find(&hyscitys)
	return hyscitys
}

func (instance *HysCityDAO) GetPageList(conditions []database.SqlCondition, start int, limit int, orderBy map[string]string, isWriteDB bool, options ...database.SqlOptions) []entity.HysCity {
	db := SqlDBRead
	if isWriteDB {
		db = SqlDBWrite
	}
	hyscitys := []entity.HysCity{}
	for orderKey, ascDes := range orderBy {
		db = db.Order(orderKey + " " + ascDes)
	}
	for _, option := range options {
		if option.SelectField != "" {
			db = db.Select(option.SelectField)
		}
	}
	instance.buildQuery(db, conditions).Offset(start).Limit(limit).Find(&hyscitys)
	return hyscitys
}

func (instance *HysCityDAO) GetPageListBySql(sqlQuery string, start int, limit int, orderBy map[string]string, isWriteDB bool, options ...database.SqlOptions) []entity.HysCity {
	db := SqlDBRead
	if isWriteDB {
		db = SqlDBWrite
	}
	hyscitys := []entity.HysCity{}
	for orderKey, ascDes := range orderBy {
		db = db.Order(orderKey + " " + ascDes)
	}
	for _, option := range options {
		if option.SelectField != "" {
			db = db.Select(option.SelectField)
		}
	}
	db.Where(sqlQuery).Offset(start).Limit(limit).Find(&hyscitys)
	return hyscitys
}

func (instance *HysCityDAO) buildQuery(db *gorm.DB, condition []database.SqlCondition) *gorm.DB {

	for _, where := range condition {
		if where.Predicate == "" {
			where.Predicate = database.SqlEqualPredicate
		}
		db = db.Where(fmt.Sprintf("%v %v ", where.QueryName, where.Predicate), where.Value)
	}

	return db
}
