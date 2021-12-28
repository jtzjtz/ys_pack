package workflow_service

import (
	"github.com/jtzjtz/kit/convert"
	"github.com/jtzjtz/kit/database"
	"github.com/jtzjtz/ys_pack/entity"
	"github.com/jtzjtz/ys_pack/proto/workflow_proto"
	"github.com/jtzjtz/ys_pack/server/dao"
)

type WorkflowService struct {
}

//创建实体
func (instance *WorkflowService) CreateWorkflow(workflowEntity entity.Workflow) (entity.Workflow, error) {
	dao := dao.WorkflowDAO{}
	workflowNew, err := dao.Create(workflowEntity)
	return workflowNew, err
}

//更新实体
func (instance *WorkflowService) UpdateWorkflow(query workflow_proto.Query, workflowUpdate map[string]interface{}) (err error) {
	dao := dao.WorkflowDAO{}
	if query.SqlQuery != "" {
		err = dao.UpdateBySql(workflowUpdate, query.SqlQuery)
	} else {
		if query.EntityQuery == nil {
			query.EntityQuery = &workflow_proto.Workflow{}
		}
		queryEntity := convert.EntityToMapWithEmpty(query.EntityQuery, query.QueryEmptyFields)
		sqlCondition := convert.MapToSqlcondition(queryEntity)
		err = dao.Update(workflowUpdate, sqlCondition)

	}
	return err

}

//删除实体
func (instance *WorkflowService) DeleteWorkflow(query workflow_proto.Query) (err error) {
	dao := dao.WorkflowDAO{}
	if query.SqlQuery != "" {
		err = dao.DeleteBySql(query.SqlQuery)

	} else {
		if query.EntityQuery == nil {
			query.EntityQuery = &workflow_proto.Workflow{}
		}
		queryEntity := convert.EntityToMapWithEmpty(query.EntityQuery, query.QueryEmptyFields)
		sqlCondition := convert.MapToSqlcondition(queryEntity)
		err = dao.Delete(sqlCondition)
	}
	return err
}

//查询实体
func (instance *WorkflowService) GetWorkflow(query workflow_proto.Query) *entity.Workflow {
	var isWriteDB bool
	isWriteDB = workflow_proto.DB_WRITE == query.Db
	options := database.SqlOptions{}
	dao := dao.WorkflowDAO{}
	if query.OrderBy != nil {
		options.OrderBy = query.OrderBy
	}
	options.SelectField = query.SelectField
	if query.SqlQuery != "" {
		entity := dao.FirstBySql(query.SqlQuery, isWriteDB, options)
		return entity

	} else {
		if query.EntityQuery == nil {
			query.EntityQuery = &workflow_proto.Workflow{}
		}
		queryEntity := convert.EntityToMapWithEmpty(query.EntityQuery, query.QueryEmptyFields)
		sqlCondition := convert.MapToSqlcondition(queryEntity)
		entity := dao.First(sqlCondition, isWriteDB, options)
		return entity

	}
}

//查询实体分页列表
func (instance *WorkflowService) GetWorkflowPageList(query workflow_proto.Query, page int, pageNum int) (list []entity.Workflow) {
	var isWriteDB bool
	dao := dao.WorkflowDAO{}
	isWriteDB = workflow_proto.DB_WRITE == query.Db
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
			query.EntityQuery = &workflow_proto.Workflow{}
		}
		queryEntity := convert.EntityToMapWithEmpty(query.EntityQuery, query.QueryEmptyFields)
		sqlCondition := convert.MapToSqlcondition(queryEntity)
		list = dao.GetPageList(sqlCondition, start, limit, orderBy, isWriteDB, options)
	}
	return list

}

//查询实体列表
func (instance *WorkflowService) GetWorkflowList(query workflow_proto.Query) (list []entity.Workflow) {
	var isWriteDB bool
	dao := dao.WorkflowDAO{}
	isWriteDB = workflow_proto.DB_WRITE == query.Db
	options := database.SqlOptions{}
	if query.OrderBy != nil {
		options.OrderBy = query.OrderBy
	}
	options.SelectField = query.SelectField
	if query.SqlQuery != "" {
		list = dao.GetListBySql(query.SqlQuery, isWriteDB, options)
	} else {
		if query.EntityQuery == nil {
			query.EntityQuery = &workflow_proto.Workflow{}
		}
		queryEntity := convert.EntityToMapWithEmpty(query.EntityQuery, query.QueryEmptyFields)
		sqlCondition := convert.MapToSqlcondition(queryEntity)
		list = dao.GetList(sqlCondition, isWriteDB, options)
	}
	return list
}

//查询实体个数
func (instance *WorkflowService) GetWorkflowCount(query workflow_proto.Query) (count int) {
	var isWriteDB bool
	dao := dao.WorkflowDAO{}
	isWriteDB = workflow_proto.DB_WRITE == query.Db
	if query.SqlQuery != "" {
		count = dao.CountBySql(query.SqlQuery, isWriteDB)
	} else {
		if query.EntityQuery == nil {
			query.EntityQuery = &workflow_proto.Workflow{}
		}
		queryEntity := convert.EntityToMapWithEmpty(query.EntityQuery, query.QueryEmptyFields)
		sqlCondition := convert.MapToSqlcondition(queryEntity)
		count = dao.Count(sqlCondition, isWriteDB)
	}
	return count
}
