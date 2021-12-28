package sq_user_blacklist_controller

import (
	"context"
	"github.com/jtzjtz/kit/convert"
	"github.com/jtzjtz/ys_pack/entity"
	"github.com/jtzjtz/ys_pack/proto/sq_user_blacklist_proto"
	"github.com/jtzjtz/ys_pack/server/service/sq_user_blacklist_service"
	"strconv"
)

type SqUserBlacklistController struct {
}

func (instance *SqUserBlacklistController) CreateSqUserBlacklist(ctx context.Context, req *sq_user_blacklist_proto.SqUserBlacklist) (result *sq_user_blacklist_proto.EntityResult, err error) {
	result = new(sq_user_blacklist_proto.EntityResult)
	defer func() {
		if err := recover(); err != nil {
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
	sqUserBlacklistEntity := entity.SqUserBlacklist{}
	err = convert.EntityToEntity(req, &sqUserBlacklistEntity)
	if err != nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = err.Error()
		return result, err
	}
	service := sq_user_blacklist_service.SqUserBlacklistService{}
	sqUserBlacklistNew, errCreate := service.CreateSqUserBlacklist(sqUserBlacklistEntity)
	if errCreate != nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = errCreate.Error()
		return result, errCreate
	}
	sqUserBlacklistOut := sq_user_blacklist_proto.SqUserBlacklist{}
	convert.EntityToEntity(sqUserBlacklistNew, &sqUserBlacklistOut)
	result.Data = &sqUserBlacklistOut
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = "create success"
	return result, err
}
func (instance *SqUserBlacklistController) UpdateSqUserBlacklist(ctx context.Context, req *sq_user_blacklist_proto.UpdateAndCondition) (result *sq_user_blacklist_proto.Result, err error) {
	result = new(sq_user_blacklist_proto.Result)
	defer func() {
		if err := recover(); err != nil {
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
	service := sq_user_blacklist_service.SqUserBlacklistService{}
	if req.Entity == nil {
		req.Entity = &sq_user_blacklist_proto.SqUserBlacklist{}
	}
	entityUpdate := convert.EntityToMapWithEmpty(req.Entity, req.UpdateEmptyFields)
	if err = service.UpdateSqUserBlacklist(*req.Query, entityUpdate); err != nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = err.Error()
		return result, err
	}
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = "create success"
	return result, err
}
func (instance *SqUserBlacklistController) GetSqUserBlacklist(ctx context.Context, req *sq_user_blacklist_proto.Query) (result *sq_user_blacklist_proto.EntityResult, err error) {
	result = new(sq_user_blacklist_proto.EntityResult)
	defer func() {
		if err := recover(); err != nil {
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
	service := sq_user_blacklist_service.SqUserBlacklistService{}
	sqUserBlacklistEntity := service.GetSqUserBlacklist(*req)
	sqUserBlacklistOut := sq_user_blacklist_proto.SqUserBlacklist{}
	if sqUserBlacklistEntity == nil {
		result.Data = nil
	} else {
		convert.EntityToEntity(sqUserBlacklistEntity, &sqUserBlacklistOut)
		result.Data = &sqUserBlacklistOut

	}
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = ""
	return result, nil

}
func (instance *SqUserBlacklistController) GetSqUserBlacklistList(ctx context.Context, req *sq_user_blacklist_proto.Query) (result *sq_user_blacklist_proto.ListResult, err error) {
	result = new(sq_user_blacklist_proto.ListResult)
	defer func() {
		if err := recover(); err != nil {
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
	service := sq_user_blacklist_service.SqUserBlacklistService{}
	sqUserBlacklistList := service.GetSqUserBlacklistList(*req)
	sqUserBlacklistOut := []*sq_user_blacklist_proto.SqUserBlacklist{}
	for _, sqUserBlacklistEntity := range sqUserBlacklistList {
		sqUserBlacklistTemp := sq_user_blacklist_proto.SqUserBlacklist{}
		convert.EntityToEntity(sqUserBlacklistEntity, &sqUserBlacklistTemp)
		sqUserBlacklistOut = append(sqUserBlacklistOut, &sqUserBlacklistTemp)

	}
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = ""
	result.Data = sqUserBlacklistOut
	return result, nil
}
func (instance *SqUserBlacklistController) DeleteSqUserBlacklist(ctx context.Context, req *sq_user_blacklist_proto.Query) (result *sq_user_blacklist_proto.Result, err error) {
	result = new(sq_user_blacklist_proto.Result)
	defer func() {
		if err := recover(); err != nil {
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
	service := sq_user_blacklist_service.SqUserBlacklistService{}
	if err = service.DeleteSqUserBlacklist(*req); err != nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = err.Error()
		return result, err
	}
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = "delete success"
	return result, err
}
func (instance *SqUserBlacklistController) GetSqUserBlacklistPageList(ctx context.Context, req *sq_user_blacklist_proto.PageQuery) (result *sq_user_blacklist_proto.PageResult, err error) {
	result = new(sq_user_blacklist_proto.PageResult)
	defer func() {
		if err := recover(); err != nil {
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
	service := sq_user_blacklist_service.SqUserBlacklistService{}
	sqUserBlacklistList := service.GetSqUserBlacklistPageList(*req.Query, int(req.Page), int(req.PageNum))
	sqUserBlacklistOut := []*sq_user_blacklist_proto.SqUserBlacklist{}
	sqUserBlacklistPageData := sq_user_blacklist_proto.SqUserBlacklistPageData{}
	for _, sqUserBlacklistEntity := range sqUserBlacklistList {
		sqUserBlacklistTemp := sq_user_blacklist_proto.SqUserBlacklist{}
		convert.EntityToEntity(sqUserBlacklistEntity, &sqUserBlacklistTemp)
		sqUserBlacklistOut = append(sqUserBlacklistOut, &sqUserBlacklistTemp)

	}
	sqUserBlacklistPageData.List = sqUserBlacklistOut
	sqUserBlacklistPageData.CurrentPage = req.Page
	sqUserBlacklistPageData.Count = int32(service.GetSqUserBlacklistCount(*req.Query))
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = ""
	result.Data = &sqUserBlacklistPageData
	return result, nil
}
func (instance *SqUserBlacklistController) GetSqUserBlacklistCount(ctx context.Context, req *sq_user_blacklist_proto.Query) (result *sq_user_blacklist_proto.Result, err error) {
	result = new(sq_user_blacklist_proto.Result)
	defer func() {
		if err := recover(); err != nil {
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
	service := sq_user_blacklist_service.SqUserBlacklistService{}
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = ""
	result.Data = strconv.Itoa(service.GetSqUserBlacklistCount(*req))
	return result, nil
}
