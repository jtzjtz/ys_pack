package workflow_grpc

import (
	"context"
	"github.com/jtzjtz/ys_pack/entity"
	"github.com/jtzjtz/ys_pack/proto/workflow_proto"
	"google.golang.org/grpc"
)

type WorkflowService struct {
}

func (t *WorkflowService) CreateWorkflow(ctx context.Context, workflowEntity *workflow_proto.Workflow, conn *grpc.ClientConn) (result *workflow_proto.EntityResult) {
	result = &workflow_proto.EntityResult{Code: entity.RESULTERRORINT32}
	if workflowEntity == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件参数错误"
	} else {
		client := workflow_proto.NewWorkflowServiceClient(conn)
		resultReponse, err := client.CreateWorkflow(ctx, workflowEntity)
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

func (t *WorkflowService) UpdateWorkflow(ctx context.Context, orderUpdateAndCondition *workflow_proto.UpdateAndCondition, conn *grpc.ClientConn) (result *workflow_proto.Result) {
	result = &workflow_proto.Result{Code: entity.RESULTERRORINT32}
	if orderUpdateAndCondition.Entity == nil || orderUpdateAndCondition.Query == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询query错误"
	} else if orderUpdateAndCondition.Query.EntityQuery == nil && orderUpdateAndCondition.Query.SqlQuery == "" {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询SqlQuery错误"
	} else {
		client := workflow_proto.NewWorkflowServiceClient(conn)
		resultReponse, err := client.UpdateWorkflow(ctx, orderUpdateAndCondition)
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

func (t *WorkflowService) GetWorkflow(ctx context.Context, query *workflow_proto.Query, conn *grpc.ClientConn) (result *workflow_proto.EntityResult) {
	result = &workflow_proto.EntityResult{Code: entity.RESULTERRORINT32}
	if query == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询query错误"
	} else if query.EntityQuery == nil && query.SqlQuery == "" {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询SqlQuery错误"
	} else {
		client := workflow_proto.NewWorkflowServiceClient(conn)
		resultReponse, err := client.GetWorkflow(ctx, query)
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

func (t *WorkflowService) GetWorkflowList(ctx context.Context, query *workflow_proto.Query, conn *grpc.ClientConn) (result *workflow_proto.ListResult) {
	result = &workflow_proto.ListResult{Code: entity.RESULTERRORINT32}
	if query == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询query错误"
	} else if query.EntityQuery == nil && query.SqlQuery == "" {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询SqlQuery错误"
	} else {
		client := workflow_proto.NewWorkflowServiceClient(conn)
		resultReponse, err := client.GetWorkflowList(ctx, query)
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

func (this *WorkflowService) DeleteWorkflow(ctx context.Context, query *workflow_proto.Query, conn *grpc.ClientConn) (result *workflow_proto.Result) {
	result = &workflow_proto.Result{Code: entity.RESULTERRORINT32}
	if query == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询query错误"
	} else if query.EntityQuery == nil && query.SqlQuery == "" {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询SqlQuery错误"
	} else {
		client := workflow_proto.NewWorkflowServiceClient(conn)
		resultReponse, err := client.DeleteWorkflow(ctx, query)
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

func (t *WorkflowService) GetWorkflowPageList(ctx context.Context, pageQuery *workflow_proto.PageQuery, conn *grpc.ClientConn) (result *workflow_proto.PageResult) {
	result = &workflow_proto.PageResult{Code: entity.RESULTERRORINT32}
	if pageQuery == nil || pageQuery.Query == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询query错误"
	} else if pageQuery.Query.EntityQuery == nil && pageQuery.Query.SqlQuery == "" {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询SqlQuery错误"

	} else {
		client := workflow_proto.NewWorkflowServiceClient(conn)
		resultReponse, err := client.GetWorkflowPageList(ctx, pageQuery)
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

func (t *WorkflowService) GetWorkflowCount(ctx context.Context, query *workflow_proto.Query, conn *grpc.ClientConn) (result *workflow_proto.Result) {
	result = &workflow_proto.Result{Code: entity.RESULTERRORINT32}
	if query == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询query错误"
	} else if query.EntityQuery == nil && query.SqlQuery == "" {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询SqlQuery错误"

	} else {
		client := workflow_proto.NewWorkflowServiceClient(conn)
		resultReponse, err := client.GetWorkflowCount(ctx, query)
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
