package department_grpc

import (
	"context"
	"github.com/jtzjtz/ys_pack/entity"
	"github.com/jtzjtz/ys_pack/proto/department_proto"
	"google.golang.org/grpc"
)

type DepartmentService struct {
}

func (t *DepartmentService) CreateDepartment(ctx context.Context, departmentEntity *department_proto.Department, conn *grpc.ClientConn) (result *department_proto.EntityResult) {
	result = &department_proto.EntityResult{Code: entity.RESULTERRORINT32}
	if departmentEntity == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件参数错误"
	} else {
		client := department_proto.NewDepartmentServiceClient(conn)
		resultReponse, err := client.CreateDepartment(ctx, departmentEntity)
		if err != nil || resultReponse == nil {
			result.Msg = err.Error()
			result.Code = entity.RESULTERRORINT32
		} else {
			result.Code = resultReponse.GetCode()
			result.Msg = resultReponse.GetMsg()
			result.Data = resultReponse.GetData()
		}
	}
	return result
}

func (t *DepartmentService) UpdateDepartment(ctx context.Context, orderUpdateAndCondition *department_proto.UpdateAndCondition, conn *grpc.ClientConn) (result *department_proto.Result) {
	result = &department_proto.Result{Code: entity.RESULTERRORINT32}
	if orderUpdateAndCondition.Entity == nil || orderUpdateAndCondition.Query == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询query错误"
	} else if orderUpdateAndCondition.Query.EntityQuery == nil && orderUpdateAndCondition.Query.SqlQuery == "" {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询SqlQuery错误"
	} else {
		client := department_proto.NewDepartmentServiceClient(conn)
		resultReponse, err := client.UpdateDepartment(ctx, orderUpdateAndCondition)
		if err != nil || resultReponse == nil {
			result.Msg = err.Error()
			result.Code = entity.RESULTERRORINT32
		} else {
			result.Code = resultReponse.GetCode()
			result.Msg = resultReponse.GetMsg()
			result.Data = resultReponse.GetData()
		}
	}
	return result
}

func (t *DepartmentService) GetDepartment(ctx context.Context, query *department_proto.Query, conn *grpc.ClientConn) (result *department_proto.EntityResult) {
	result = &department_proto.EntityResult{Code: entity.RESULTERRORINT32}
	if query == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询query错误"
	} else if query.EntityQuery == nil && query.SqlQuery == "" {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询SqlQuery错误"
	} else {
		client := department_proto.NewDepartmentServiceClient(conn)
		resultReponse, err := client.GetDepartment(ctx, query)
		if err != nil || resultReponse == nil {
			result.Msg = err.Error()
			result.Code = entity.RESULTERRORINT32
		} else {
			result.Code = resultReponse.GetCode()
			result.Msg = resultReponse.GetMsg()
			result.Data = resultReponse.GetData()
		}
	}
	return result
}

func (t *DepartmentService) GetDepartmentList(ctx context.Context, query *department_proto.Query, conn *grpc.ClientConn) (result *department_proto.ListResult) {
	result = &department_proto.ListResult{Code: entity.RESULTERRORINT32}
	if query == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询query错误"
	} else if query.EntityQuery == nil && query.SqlQuery == "" {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询SqlQuery错误"
	} else {
		client := department_proto.NewDepartmentServiceClient(conn)
		resultReponse, err := client.GetDepartmentList(ctx, query)
		if err != nil || resultReponse == nil {
			result.Msg = err.Error()
			result.Code = entity.RESULTERRORINT32
		} else {
			result.Code = resultReponse.GetCode()
			result.Msg = resultReponse.GetMsg()
			result.Data = resultReponse.GetData()
		}
	}
	return result
}

func (this *DepartmentService) DeleteDepartment(ctx context.Context, query *department_proto.Query, conn *grpc.ClientConn) (result *department_proto.Result) {
	result = &department_proto.Result{Code: entity.RESULTERRORINT32}
	if query == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询query错误"
	} else if query.EntityQuery == nil && query.SqlQuery == "" {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询SqlQuery错误"
	} else {
		client := department_proto.NewDepartmentServiceClient(conn)
		resultReponse, err := client.DeleteDepartment(ctx, query)
		if err != nil || resultReponse == nil {
			result.Msg = err.Error()
			result.Code = entity.RESULTERRORINT32
		} else {
			result.Code = resultReponse.GetCode()
			result.Msg = resultReponse.GetMsg()
			result.Data = resultReponse.GetData()
		}
	}
	return result
}

func (t *DepartmentService) GetDepartmentPageList(ctx context.Context, pageQuery *department_proto.PageQuery, conn *grpc.ClientConn) (result *department_proto.PageResult) {
	result = &department_proto.PageResult{Code: entity.RESULTERRORINT32}
	if pageQuery == nil || pageQuery.Query == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询query错误"
	} else if pageQuery.Query.EntityQuery == nil && pageQuery.Query.SqlQuery == "" {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询SqlQuery错误"

	} else {
		client := department_proto.NewDepartmentServiceClient(conn)
		resultReponse, err := client.GetDepartmentPageList(ctx, pageQuery)
		if err != nil || resultReponse == nil {
			result.Msg = err.Error()
			result.Code = entity.RESULTERRORINT32
		} else {
			result.Code = resultReponse.GetCode()
			result.Msg = resultReponse.GetMsg()
			result.Data = resultReponse.GetData()
		}
	}
	return result
}

func (t *DepartmentService) GetDepartmentCount(ctx context.Context, query *department_proto.Query, conn *grpc.ClientConn) (result *department_proto.Result) {
	result = &department_proto.Result{Code: entity.RESULTERRORINT32}
	if query == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询query错误"
	} else if query.EntityQuery == nil && query.SqlQuery == "" {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询SqlQuery错误"

	} else {
		client := department_proto.NewDepartmentServiceClient(conn)
		resultReponse, err := client.GetDepartmentCount(ctx, query)
		if err != nil || resultReponse == nil {
			result.Msg = err.Error()
			result.Code = entity.RESULTERRORINT32
		} else {
			result.Code = resultReponse.GetCode()
			result.Msg = resultReponse.GetMsg()
			result.Data = resultReponse.GetData()
		}
	}
	return result
}
