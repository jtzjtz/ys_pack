package hys_city_service

import (
	"github.com/jtzjtz/kit/convert"
	"github.com/jtzjtz/kit/database"
	"github.com/jtzjtz/ys_pack/entity"
	"github.com/jtzjtz/ys_pack/proto/hys_city_proto"
	"github.com/jtzjtz/ys_pack/server/dao"
)

type HysCityService struct {
}

//创建实体
func (instance *HysCityService) CreateHysCity(hysCityEntity entity.HysCity) (entity.HysCity, error) {
	dao := dao.HysCityDAO{}
	hysCityNew, err := dao.Create(hysCityEntity)
	return hysCityNew, err
}

//更新实体
func (instance *HysCityService) UpdateHysCity(query hys_city_proto.Query, hysCityUpdate map[string]interface{}) (err error) {
	dao := dao.HysCityDAO{}
	if query.SqlQuery != "" {
		err = dao.UpdateBySql(hysCityUpdate, query.SqlQuery)
	} else {
		if query.EntityQuery == nil {
			query.EntityQuery = &hys_city_proto.HysCity{}
		}
		queryEntity := convert.EntityToMapWithEmpty(query.EntityQuery, query.QueryEmptyFields)
		sqlCondition := convert.MapToSqlcondition(queryEntity)
		err = dao.Update(hysCityUpdate, sqlCondition)

	}
	return err

}

//删除实体
func (instance *HysCityService) DeleteHysCity(query hys_city_proto.Query) (err error) {
	dao := dao.HysCityDAO{}
	if query.SqlQuery != "" {
		err = dao.DeleteBySql(query.SqlQuery)

	} else {
		if query.EntityQuery == nil {
			query.EntityQuery = &hys_city_proto.HysCity{}
		}
		queryEntity := convert.EntityToMapWithEmpty(query.EntityQuery, query.QueryEmptyFields)
		sqlCondition := convert.MapToSqlcondition(queryEntity)
		err = dao.Delete(sqlCondition)
	}
	return err
}

//查询实体
func (instance *HysCityService) GetHysCity(query hys_city_proto.Query) *entity.HysCity {
	var isWriteDB bool
	isWriteDB = hys_city_proto.DB_WRITE == query.Db
	options := database.SqlOptions{}
	dao := dao.HysCityDAO{}
	if query.OrderBy != nil {
		options.OrderBy = query.OrderBy
	}
	options.SelectField = query.SelectField
	if query.SqlQuery != "" {
		entity := dao.FirstBySql(query.SqlQuery, isWriteDB, options)
		return entity

	} else {
		if query.EntityQuery == nil {
			query.EntityQuery = &hys_city_proto.HysCity{}
		}
		queryEntity := convert.EntityToMapWithEmpty(query.EntityQuery, query.QueryEmptyFields)
		sqlCondition := convert.MapToSqlcondition(queryEntity)
		entity := dao.First(sqlCondition, isWriteDB, options)
		return entity

	}
}

//查询实体分页列表
func (instance *HysCityService) GetHysCityPageList(query hys_city_proto.Query, page int, pageNum int) (list []entity.HysCity) {
	var isWriteDB bool
	dao := dao.HysCityDAO{}
	isWriteDB = hys_city_proto.DB_WRITE == query.Db
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
			query.EntityQuery = &hys_city_proto.HysCity{}
		}
		queryEntity := convert.EntityToMapWithEmpty(query.EntityQuery, query.QueryEmptyFields)
		sqlCondition := convert.MapToSqlcondition(queryEntity)
		list = dao.GetPageList(sqlCondition, start, limit, orderBy, isWriteDB, options)
	}
	return list

}

//查询实体列表
func (instance *HysCityService) GetHysCityList(query hys_city_proto.Query) (list []entity.HysCity) {
	var isWriteDB bool
	dao := dao.HysCityDAO{}
	isWriteDB = hys_city_proto.DB_WRITE == query.Db
	options := database.SqlOptions{}
	if query.OrderBy != nil {
		options.OrderBy = query.OrderBy
	}
	options.SelectField = query.SelectField
	if query.SqlQuery != "" {
		list = dao.GetListBySql(query.SqlQuery, isWriteDB, options)
	} else {
		if query.EntityQuery == nil {
			query.EntityQuery = &hys_city_proto.HysCity{}
		}
		queryEntity := convert.EntityToMapWithEmpty(query.EntityQuery, query.QueryEmptyFields)
		sqlCondition := convert.MapToSqlcondition(queryEntity)
		list = dao.GetList(sqlCondition, isWriteDB, options)
	}
	return list
}

//查询实体个数
func (instance *HysCityService) GetHysCityCount(query hys_city_proto.Query) (count int) {
	var isWriteDB bool
	dao := dao.HysCityDAO{}
	isWriteDB = hys_city_proto.DB_WRITE == query.Db
	if query.SqlQuery != "" {
		count = dao.CountBySql(query.SqlQuery, isWriteDB)
	} else {
		if query.EntityQuery == nil {
			query.EntityQuery = &hys_city_proto.HysCity{}
		}
		queryEntity := convert.EntityToMapWithEmpty(query.EntityQuery, query.QueryEmptyFields)
		sqlCondition := convert.MapToSqlcondition(queryEntity)
		count = dao.Count(sqlCondition, isWriteDB)
	}
	return count
}
