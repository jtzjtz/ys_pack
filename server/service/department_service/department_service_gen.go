package department_service

import (
	"github.com/jtzjtz/kit/convert"
	"github.com/jtzjtz/kit/database"
	"github.com/jtzjtz/ys_pack/entity"
	"github.com/jtzjtz/ys_pack/proto/department_proto"
	"github.com/jtzjtz/ys_pack/server/dao"
)

type DepartmentService struct {
}

//创建实体
func (instance *DepartmentService) CreateDepartment(departmentEntity entity.Department) (entity.Department, error) {
	dao := dao.DepartmentDAO{}
	departmentNew, err := dao.Create(departmentEntity)
	return departmentNew, err
}

//更新实体
func (instance *DepartmentService) UpdateDepartment(query department_proto.Query, departmentUpdate map[string]interface{}) (err error) {
	dao := dao.DepartmentDAO{}
	if query.SqlQuery != "" {
		err = dao.UpdateBySql(departmentUpdate, query.SqlQuery)
	} else {
		if query.EntityQuery == nil {
			query.EntityQuery = &department_proto.Department{}
		}
		queryEntity := convert.EntityToMapWithEmpty(query.EntityQuery, query.QueryEmptyFields)
		sqlCondition := convert.MapToSqlcondition(queryEntity)
		err = dao.Update(departmentUpdate, sqlCondition)

	}
	return err

}

//删除实体
func (instance *DepartmentService) DeleteDepartment(query department_proto.Query) (err error) {
	dao := dao.DepartmentDAO{}
	if query.SqlQuery != "" {
		err = dao.DeleteBySql(query.SqlQuery)

	} else {
		if query.EntityQuery == nil {
			query.EntityQuery = &department_proto.Department{}
		}
		queryEntity := convert.EntityToMapWithEmpty(query.EntityQuery, query.QueryEmptyFields)
		sqlCondition := convert.MapToSqlcondition(queryEntity)
		err = dao.Delete(sqlCondition)
	}
	return err
}

//查询实体
func (instance *DepartmentService) GetDepartment(query department_proto.Query) *entity.Department {
	var isWriteDB bool
	isWriteDB = department_proto.DB_WRITE == query.Db
	options := database.SqlOptions{}
	dao := dao.DepartmentDAO{}
	if query.OrderBy != nil {
		options.OrderBy = query.OrderBy
	}
	options.SelectField = query.SelectField
	if query.SqlQuery != "" {
		entity := dao.FirstBySql(query.SqlQuery, isWriteDB, options)
		return entity

	} else {
		if query.EntityQuery == nil {
			query.EntityQuery = &department_proto.Department{}
		}
		queryEntity := convert.EntityToMapWithEmpty(query.EntityQuery, query.QueryEmptyFields)
		sqlCondition := convert.MapToSqlcondition(queryEntity)
		entity := dao.First(sqlCondition, isWriteDB, options)
		return entity

	}
}

//查询实体分页列表
func (instance *DepartmentService) GetDepartmentPageList(query department_proto.Query, page int, pageNum int) (list []entity.Department) {
	var isWriteDB bool
	dao := dao.DepartmentDAO{}
	isWriteDB = department_proto.DB_WRITE == query.Db
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
			query.EntityQuery = &department_proto.Department{}
		}
		queryEntity := convert.EntityToMapWithEmpty(query.EntityQuery, query.QueryEmptyFields)
		sqlCondition := convert.MapToSqlcondition(queryEntity)
		list = dao.GetPageList(sqlCondition, start, limit, orderBy, isWriteDB, options)
	}
	return list

}

//查询实体列表
func (instance *DepartmentService) GetDepartmentList(query department_proto.Query) (list []entity.Department) {
	var isWriteDB bool
	dao := dao.DepartmentDAO{}
	isWriteDB = department_proto.DB_WRITE == query.Db
	options := database.SqlOptions{}
	if query.OrderBy != nil {
		options.OrderBy = query.OrderBy
	}
	options.SelectField = query.SelectField
	if query.SqlQuery != "" {
		list = dao.GetListBySql(query.SqlQuery, isWriteDB, options)
	} else {
		if query.EntityQuery == nil {
			query.EntityQuery = &department_proto.Department{}
		}
		queryEntity := convert.EntityToMapWithEmpty(query.EntityQuery, query.QueryEmptyFields)
		sqlCondition := convert.MapToSqlcondition(queryEntity)
		list = dao.GetList(sqlCondition, isWriteDB, options)
	}
	return list
}

//查询实体个数
func (instance *DepartmentService) GetDepartmentCount(query department_proto.Query) (count int) {
	var isWriteDB bool
	dao := dao.DepartmentDAO{}
	isWriteDB = department_proto.DB_WRITE == query.Db
	if query.SqlQuery != "" {
		count = dao.CountBySql(query.SqlQuery, isWriteDB)
	} else {
		if query.EntityQuery == nil {
			query.EntityQuery = &department_proto.Department{}
		}
		queryEntity := convert.EntityToMapWithEmpty(query.EntityQuery, query.QueryEmptyFields)
		sqlCondition := convert.MapToSqlcondition(queryEntity)
		count = dao.Count(sqlCondition, isWriteDB)
	}
	return count
}
