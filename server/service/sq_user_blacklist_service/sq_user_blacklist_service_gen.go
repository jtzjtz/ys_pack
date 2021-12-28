package sq_user_blacklist_service

import (
	"github.com/jtzjtz/kit/convert"
	"github.com/jtzjtz/kit/database"
	"github.com/jtzjtz/ys_pack/entity"
	"github.com/jtzjtz/ys_pack/proto/sq_user_blacklist_proto"
	"github.com/jtzjtz/ys_pack/server/dao"
)

type SqUserBlacklistService struct {
}

//创建实体
func (instance *SqUserBlacklistService) CreateSqUserBlacklist(sqUserBlacklistEntity entity.SqUserBlacklist) (entity.SqUserBlacklist, error) {
	dao := dao.SqUserBlacklistDAO{}
	sqUserBlacklistNew, err := dao.Create(sqUserBlacklistEntity)
	return sqUserBlacklistNew, err
}

//更新实体
func (instance *SqUserBlacklistService) UpdateSqUserBlacklist(query sq_user_blacklist_proto.Query, sqUserBlacklistUpdate map[string]interface{}) (err error) {
	dao := dao.SqUserBlacklistDAO{}
	if query.SqlQuery != "" {
		err = dao.UpdateBySql(sqUserBlacklistUpdate, query.SqlQuery)
	} else {
		if query.EntityQuery == nil {
			query.EntityQuery = &sq_user_blacklist_proto.SqUserBlacklist{}
		}
		queryEntity := convert.EntityToMapWithEmpty(query.EntityQuery, query.QueryEmptyFields)
		sqlCondition := convert.MapToSqlcondition(queryEntity)
		err = dao.Update(sqUserBlacklistUpdate, sqlCondition)

	}
	return err

}

//删除实体
func (instance *SqUserBlacklistService) DeleteSqUserBlacklist(query sq_user_blacklist_proto.Query) (err error) {
	dao := dao.SqUserBlacklistDAO{}
	if query.SqlQuery != "" {
		err = dao.DeleteBySql(query.SqlQuery)

	} else {
		if query.EntityQuery == nil {
			query.EntityQuery = &sq_user_blacklist_proto.SqUserBlacklist{}
		}
		queryEntity := convert.EntityToMapWithEmpty(query.EntityQuery, query.QueryEmptyFields)
		sqlCondition := convert.MapToSqlcondition(queryEntity)
		err = dao.Delete(sqlCondition)
	}
	return err
}

//查询实体
func (instance *SqUserBlacklistService) GetSqUserBlacklist(query sq_user_blacklist_proto.Query) *entity.SqUserBlacklist {
	var isWriteDB bool
	isWriteDB = sq_user_blacklist_proto.DB_WRITE == query.Db
	options := database.SqlOptions{}
	dao := dao.SqUserBlacklistDAO{}
	if query.OrderBy != nil {
		options.OrderBy = query.OrderBy
	}
	options.SelectField = query.SelectField
	if query.SqlQuery != "" {
		entity := dao.FirstBySql(query.SqlQuery, isWriteDB, options)
		return entity

	} else {
		if query.EntityQuery == nil {
			query.EntityQuery = &sq_user_blacklist_proto.SqUserBlacklist{}
		}
		queryEntity := convert.EntityToMapWithEmpty(query.EntityQuery, query.QueryEmptyFields)
		sqlCondition := convert.MapToSqlcondition(queryEntity)
		entity := dao.First(sqlCondition, isWriteDB, options)
		return entity

	}
}

//查询实体分页列表
func (instance *SqUserBlacklistService) GetSqUserBlacklistPageList(query sq_user_blacklist_proto.Query, page int, pageNum int) (list []entity.SqUserBlacklist) {
	var isWriteDB bool
	dao := dao.SqUserBlacklistDAO{}
	isWriteDB = sq_user_blacklist_proto.DB_WRITE == query.Db
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
			query.EntityQuery = &sq_user_blacklist_proto.SqUserBlacklist{}
		}
		queryEntity := convert.EntityToMapWithEmpty(query.EntityQuery, query.QueryEmptyFields)
		sqlCondition := convert.MapToSqlcondition(queryEntity)
		list = dao.GetPageList(sqlCondition, start, limit, orderBy, isWriteDB, options)
	}
	return list

}

//查询实体列表
func (instance *SqUserBlacklistService) GetSqUserBlacklistList(query sq_user_blacklist_proto.Query) (list []entity.SqUserBlacklist) {
	var isWriteDB bool
	dao := dao.SqUserBlacklistDAO{}
	isWriteDB = sq_user_blacklist_proto.DB_WRITE == query.Db
	options := database.SqlOptions{}
	if query.OrderBy != nil {
		options.OrderBy = query.OrderBy
	}
	options.SelectField = query.SelectField
	if query.SqlQuery != "" {
		list = dao.GetListBySql(query.SqlQuery, isWriteDB, options)
	} else {
		if query.EntityQuery == nil {
			query.EntityQuery = &sq_user_blacklist_proto.SqUserBlacklist{}
		}
		queryEntity := convert.EntityToMapWithEmpty(query.EntityQuery, query.QueryEmptyFields)
		sqlCondition := convert.MapToSqlcondition(queryEntity)
		list = dao.GetList(sqlCondition, isWriteDB, options)
	}
	return list
}

//查询实体个数
func (instance *SqUserBlacklistService) GetSqUserBlacklistCount(query sq_user_blacklist_proto.Query) (count int) {
	var isWriteDB bool
	dao := dao.SqUserBlacklistDAO{}
	isWriteDB = sq_user_blacklist_proto.DB_WRITE == query.Db
	if query.SqlQuery != "" {
		count = dao.CountBySql(query.SqlQuery, isWriteDB)
	} else {
		if query.EntityQuery == nil {
			query.EntityQuery = &sq_user_blacklist_proto.SqUserBlacklist{}
		}
		queryEntity := convert.EntityToMapWithEmpty(query.EntityQuery, query.QueryEmptyFields)
		sqlCondition := convert.MapToSqlcondition(queryEntity)
		count = dao.Count(sqlCondition, isWriteDB)
	}
	return count
}
