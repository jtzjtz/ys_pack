package workflow_controller

import (
	"context"
	"github.com/jtzjtz/kit/convert"
	"github.com/jtzjtz/ys_pack/entity"
	"github.com/jtzjtz/ys_pack/proto/workflow_proto"
	"github.com/jtzjtz/ys_pack/server/service/workflow_service"
	"strconv"
)

type WorkflowController struct {
}

func (instance *WorkflowController) CreateWorkflow(ctx context.Context, req *workflow_proto.Workflow) (result *workflow_proto.EntityResult, err error) {
	result = new(workflow_proto.EntityResult)
	defer func() {
		if err := recover(); err != nil {
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
	workflowEntity := entity.Workflow{}
	err = convert.EntityToEntity(req, &workflowEntity)
	if err != nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = err.Error()
		return result, err
	}
	service := workflow_service.WorkflowService{}
	workflowNew, errCreate := service.CreateWorkflow(workflowEntity)
	if errCreate != nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = errCreate.Error()
		return result, errCreate
	}
	workflowOut := workflow_proto.Workflow{}
	convert.EntityToEntity(workflowNew, &workflowOut)
	result.Data = &workflowOut
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = "create success"
	return result, err
}
func (instance *WorkflowController) UpdateWorkflow(ctx context.Context, req *workflow_proto.UpdateAndCondition) (result *workflow_proto.Result, err error) {
	result = new(workflow_proto.Result)
	defer func() {
		if err := recover(); err != nil {
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
	service := workflow_service.WorkflowService{}
	if req.Entity == nil {
		req.Entity = &workflow_proto.Workflow{}
	}
	entityUpdate := convert.EntityToMapWithEmpty(req.Entity, req.UpdateEmptyFields)
	if err = service.UpdateWorkflow(*req.Query, entityUpdate); err != nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = err.Error()
		return result, err
	}
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = "create success"
	return result, err
}
func (instance *WorkflowController) GetWorkflow(ctx context.Context, req *workflow_proto.Query) (result *workflow_proto.EntityResult, err error) {
	result = new(workflow_proto.EntityResult)
	defer func() {
		if err := recover(); err != nil {
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
	service := workflow_service.WorkflowService{}
	workflowEntity := service.GetWorkflow(*req)
	workflowOut := workflow_proto.Workflow{}
	if workflowEntity == nil {
		result.Data = nil
	} else {
		convert.EntityToEntity(workflowEntity, &workflowOut)
		result.Data = &workflowOut

	}
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = ""
	return result, nil

}
func (instance *WorkflowController) GetWorkflowList(ctx context.Context, req *workflow_proto.Query) (result *workflow_proto.ListResult, err error) {
	result = new(workflow_proto.ListResult)
	defer func() {
		if err := recover(); err != nil {
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
	service := workflow_service.WorkflowService{}
	workflowList := service.GetWorkflowList(*req)
	workflowOut := []*workflow_proto.Workflow{}
	for _, workflowEntity := range workflowList {
		workflowTemp := workflow_proto.Workflow{}
		convert.EntityToEntity(workflowEntity, &workflowTemp)
		workflowOut = append(workflowOut, &workflowTemp)

	}
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = ""
	result.Data = workflowOut
	return result, nil
}
func (instance *WorkflowController) DeleteWorkflow(ctx context.Context, req *workflow_proto.Query) (result *workflow_proto.Result, err error) {
	result = new(workflow_proto.Result)
	defer func() {
		if err := recover(); err != nil {
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
	service := workflow_service.WorkflowService{}
	if err = service.DeleteWorkflow(*req); err != nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = err.Error()
		return result, err
	}
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = "delete success"
	return result, err
}
func (instance *WorkflowController) GetWorkflowPageList(ctx context.Context, req *workflow_proto.PageQuery) (result *workflow_proto.PageResult, err error) {
	result = new(workflow_proto.PageResult)
	defer func() {
		if err := recover(); err != nil {
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
	service := workflow_service.WorkflowService{}
	workflowList := service.GetWorkflowPageList(*req.Query, int(req.Page), int(req.PageNum))
	workflowOut := []*workflow_proto.Workflow{}
	workflowPageData := workflow_proto.WorkflowPageData{}
	for _, workflowEntity := range workflowList {
		workflowTemp := workflow_proto.Workflow{}
		convert.EntityToEntity(workflowEntity, &workflowTemp)
		workflowOut = append(workflowOut, &workflowTemp)

	}
	workflowPageData.List = workflowOut
	workflowPageData.CurrentPage = req.Page
	workflowPageData.Count = int32(service.GetWorkflowCount(*req.Query))
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = ""
	result.Data = &workflowPageData
	return result, nil
}
func (instance *WorkflowController) GetWorkflowCount(ctx context.Context, req *workflow_proto.Query) (result *workflow_proto.Result, err error) {
	result = new(workflow_proto.Result)
	defer func() {
		if err := recover(); err != nil {
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
	service := workflow_service.WorkflowService{}
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = ""
	result.Data = strconv.Itoa(service.GetWorkflowCount(*req))
	return result, nil
}
