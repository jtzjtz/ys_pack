package department_controller

import (
	"context"
	"github.com/jtzjtz/kit/convert"
	"github.com/jtzjtz/ys_pack/entity"
	"github.com/jtzjtz/ys_pack/proto/department_proto"
	"github.com/jtzjtz/ys_pack/server/service/department_service"
	"strconv"
)

type DepartmentController struct {
}

func (instance *DepartmentController) CreateDepartment(ctx context.Context, req *department_proto.Department) (result *department_proto.EntityResult, err error) {
	result = new(department_proto.EntityResult)
	defer func() {
		if err := recover(); err != nil {
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
	departmentEntity := entity.Department{}
	err = convert.EntityToEntity(req, &departmentEntity)
	if err != nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = err.Error()
		return result, err
	}
	service := department_service.DepartmentService{}
	departmentNew, errCreate := service.CreateDepartment(departmentEntity)
	if errCreate != nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = errCreate.Error()
		return result, errCreate
	}
	departmentOut := department_proto.Department{}
	convert.EntityToEntity(departmentNew, &departmentOut)
	result.Data = &departmentOut
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = "create success"
	return result, err
}
func (instance *DepartmentController) UpdateDepartment(ctx context.Context, req *department_proto.UpdateAndCondition) (result *department_proto.Result, err error) {
	result = new(department_proto.Result)
	defer func() {
		if err := recover(); err != nil {
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
	service := department_service.DepartmentService{}
	if req.Entity == nil {
		req.Entity = &department_proto.Department{}
	}
	entityUpdate := convert.EntityToMapWithEmpty(req.Entity, req.UpdateEmptyFields)
	if err = service.UpdateDepartment(*req.Query, entityUpdate); err != nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = err.Error()
		return result, err
	}
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = "create success"
	return result, err
}
func (instance *DepartmentController) GetDepartment(ctx context.Context, req *department_proto.Query) (result *department_proto.EntityResult, err error) {
	result = new(department_proto.EntityResult)
	defer func() {
		if err := recover(); err != nil {
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
	service := department_service.DepartmentService{}
	departmentEntity := service.GetDepartment(*req)
	departmentOut := department_proto.Department{}
	if departmentEntity == nil {
		result.Data = nil
	} else {
		convert.EntityToEntity(departmentEntity, &departmentOut)
		result.Data = &departmentOut

	}
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = ""
	return result, nil

}
func (instance *DepartmentController) GetDepartmentList(ctx context.Context, req *department_proto.Query) (result *department_proto.ListResult, err error) {
	result = new(department_proto.ListResult)
	defer func() {
		if err := recover(); err != nil {
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
	service := department_service.DepartmentService{}
	departmentList := service.GetDepartmentList(*req)
	departmentOut := []*department_proto.Department{}
	for _, departmentEntity := range departmentList {
		departmentTemp := department_proto.Department{}
		convert.EntityToEntity(departmentEntity, &departmentTemp)
		departmentOut = append(departmentOut, &departmentTemp)

	}
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = ""
	result.Data = departmentOut
	return result, nil
}
func (instance *DepartmentController) DeleteDepartment(ctx context.Context, req *department_proto.Query) (result *department_proto.Result, err error) {
	result = new(department_proto.Result)
	defer func() {
		if err := recover(); err != nil {
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
	service := department_service.DepartmentService{}
	if err = service.DeleteDepartment(*req); err != nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = err.Error()
		return result, err
	}
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = "delete success"
	return result, err
}
func (instance *DepartmentController) GetDepartmentPageList(ctx context.Context, req *department_proto.PageQuery) (result *department_proto.PageResult, err error) {
	result = new(department_proto.PageResult)
	defer func() {
		if err := recover(); err != nil {
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
	service := department_service.DepartmentService{}
	departmentList := service.GetDepartmentPageList(*req.Query, int(req.Page), int(req.PageNum))
	departmentOut := []*department_proto.Department{}
	departmentPageData := department_proto.DepartmentPageData{}
	for _, departmentEntity := range departmentList {
		departmentTemp := department_proto.Department{}
		convert.EntityToEntity(departmentEntity, &departmentTemp)
		departmentOut = append(departmentOut, &departmentTemp)

	}
	departmentPageData.List = departmentOut
	departmentPageData.CurrentPage = req.Page
	departmentPageData.Count = int32(service.GetDepartmentCount(*req.Query))
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = ""
	result.Data = &departmentPageData
	return result, nil
}
func (instance *DepartmentController) GetDepartmentCount(ctx context.Context, req *department_proto.Query) (result *department_proto.Result, err error) {
	result = new(department_proto.Result)
	defer func() {
		if err := recover(); err != nil {
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
	service := department_service.DepartmentService{}
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = ""
	result.Data = strconv.Itoa(service.GetDepartmentCount(*req))
	return result, nil
}
