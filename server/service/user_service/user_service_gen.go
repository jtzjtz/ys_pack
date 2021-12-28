package user_service

import (
	"github.com/jtzjtz/kit/convert"
	"github.com/jtzjtz/kit/database"
	"github.com/jtzjtz/ys_pack/entity"
	"github.com/jtzjtz/ys_pack/proto/user_proto"
	"github.com/jtzjtz/ys_pack/server/dao"
)

type UserService struct {
}

//创建实体
func (instance *UserService) CreateUser(userEntity entity.User) (entity.User, error) {
	dao := dao.UserDAO{}
	userNew, err := dao.Create(userEntity)
	return userNew, err
}

//更新实体
func (instance *UserService) UpdateUser(query user_proto.Query, userUpdate map[string]interface{}) (err error) {
	dao := dao.UserDAO{}
	if query.SqlQuery != "" {
		err = dao.UpdateBySql(userUpdate, query.SqlQuery)
	} else {
		if query.EntityQuery == nil {
			query.EntityQuery = &user_proto.User{}
		}
		queryEntity := convert.EntityToMapWithEmpty(query.EntityQuery, query.QueryEmptyFields)
		sqlCondition := convert.MapToSqlcondition(queryEntity)
		err = dao.Update(userUpdate, sqlCondition)

	}
	return err

}

//删除实体
func (instance *UserService) DeleteUser(query user_proto.Query) (err error) {
	dao := dao.UserDAO{}
	if query.SqlQuery != "" {
		err = dao.DeleteBySql(query.SqlQuery)

	} else {
		if query.EntityQuery == nil {
			query.EntityQuery = &user_proto.User{}
		}
		queryEntity := convert.EntityToMapWithEmpty(query.EntityQuery, query.QueryEmptyFields)
		sqlCondition := convert.MapToSqlcondition(queryEntity)
		err = dao.Delete(sqlCondition)
	}
	return err
}

//查询实体
func (instance *UserService) GetUser(query user_proto.Query) *entity.User {
	var isWriteDB bool
	isWriteDB = user_proto.DB_WRITE == query.Db
	options := database.SqlOptions{}
	dao := dao.UserDAO{}
	if query.OrderBy != nil {
		options.OrderBy = query.OrderBy
	}
	options.SelectField = query.SelectField
	if query.SqlQuery != "" {
		entity := dao.FirstBySql(query.SqlQuery, isWriteDB, options)
		return entity

	} else {
		if query.EntityQuery == nil {
			query.EntityQuery = &user_proto.User{}
		}
		queryEntity := convert.EntityToMapWithEmpty(query.EntityQuery, query.QueryEmptyFields)
		sqlCondition := convert.MapToSqlcondition(queryEntity)
		entity := dao.First(sqlCondition, isWriteDB, options)
		return entity

	}
}

//查询实体分页列表
func (instance *UserService) GetUserPageList(query user_proto.Query, page int, pageNum int) (list []entity.User) {
	var isWriteDB bool
	dao := dao.UserDAO{}
	isWriteDB = user_proto.DB_WRITE == query.Db
	orderBy := query.OrderBy
	start := 0
	limit := pageNum
	if page == 0 {
		page = 1
	}
	start = (page - 1) * pageNum
	options := database.SqlOptions{}
	options.SelectField = query.SelectField
	if query.SqlQuery != "" {
		list = dao.GetPageListBySql(query.SqlQuery, start, limit, orderBy, isWriteDB, options)
	} else {
		if query.EntityQuery == nil {
			query.EntityQuery = &user_proto.User{}
		}
		queryEntity := convert.EntityToMapWithEmpty(query.EntityQuery, query.QueryEmptyFields)
		sqlCondition := convert.MapToSqlcondition(queryEntity)
		list = dao.GetPageList(sqlCondition, start, limit, orderBy, isWriteDB, options)
	}
	return list

}

//查询实体列表
func (instance *UserService) GetUserList(query user_proto.Query) (list []entity.User) {
	var isWriteDB bool
	dao := dao.UserDAO{}
	isWriteDB = user_proto.DB_WRITE == query.Db
	options := database.SqlOptions{}
	if query.OrderBy != nil {
		options.OrderBy = query.OrderBy
	}
	options.SelectField = query.SelectField
	if query.SqlQuery != "" {
		list = dao.GetListBySql(query.SqlQuery, isWriteDB, options)
	} else {
		if query.EntityQuery == nil {
			query.EntityQuery = &user_proto.User{}
		}
		queryEntity := convert.EntityToMapWithEmpty(query.EntityQuery, query.QueryEmptyFields)
		sqlCondition := convert.MapToSqlcondition(queryEntity)
		list = dao.GetList(sqlCondition, isWriteDB, options)
	}
	return list
}

//查询实体个数
func (instance *UserService) GetUserCount(query user_proto.Query) (count int) {
	var isWriteDB bool
	dao := dao.UserDAO{}
	isWriteDB = user_proto.DB_WRITE == query.Db
	if query.SqlQuery != "" {
		count = dao.CountBySql(query.SqlQuery, isWriteDB)
	} else {
		if query.EntityQuery == nil {
			query.EntityQuery = &user_proto.User{}
		}
		queryEntity := convert.EntityToMapWithEmpty(query.EntityQuery, query.QueryEmptyFields)
		sqlCondition := convert.MapToSqlcondition(queryEntity)
		count = dao.Count(sqlCondition, isWriteDB)
	}
	return count
}
