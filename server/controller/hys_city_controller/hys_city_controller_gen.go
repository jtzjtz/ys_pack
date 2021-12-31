package hys_city_controller

import (
	"context"
	"github.com/jtzjtz/kit/convert"
	"github.com/jtzjtz/ys_pack/entity"
	"github.com/jtzjtz/ys_pack/proto/hys_city_proto"
	"github.com/jtzjtz/ys_pack/server/service/hys_city_service"
	"strconv"
)

type HysCityController struct {
}

func (instance *HysCityController) CreateHysCity(ctx context.Context, req *hys_city_proto.HysCity) (result *hys_city_proto.EntityResult, err error) {
	result = new(hys_city_proto.EntityResult)
	defer func() {
		if err := recover(); err != nil {
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
	hysCityEntity := entity.HysCity{}
	err = convert.EntityToEntity(req, &hysCityEntity)
	if err != nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = err.Error()
		return result, err
	}
	service := hys_city_service.HysCityService{}
	hysCityNew, errCreate := service.CreateHysCity(hysCityEntity)
	if errCreate != nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = errCreate.Error()
		return result, errCreate
	}
	hysCityOut := hys_city_proto.HysCity{}
	convert.EntityToEntity(hysCityNew, &hysCityOut)
	result.Data = &hysCityOut
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = "create success"
	return result, err
}
func (instance *HysCityController) UpdateHysCity(ctx context.Context, req *hys_city_proto.UpdateAndCondition) (result *hys_city_proto.Result, err error) {
	result = new(hys_city_proto.Result)
	defer func() {
		if err := recover(); err != nil {
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
	service := hys_city_service.HysCityService{}
	if req.Entity == nil {
		req.Entity = &hys_city_proto.HysCity{}
	}
	entityUpdate := convert.EntityToMapWithEmpty(req.Entity, req.UpdateEmptyFields)
	if err = service.UpdateHysCity(*req.Query, entityUpdate); err != nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = err.Error()
		return result, err
	}
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = "create success"
	return result, err
}
func (instance *HysCityController) GetHysCity(ctx context.Context, req *hys_city_proto.Query) (result *hys_city_proto.EntityResult, err error) {
	result = new(hys_city_proto.EntityResult)
	defer func() {
		if err := recover(); err != nil {
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
	service := hys_city_service.HysCityService{}
	hysCityEntity := service.GetHysCity(*req)
	hysCityOut := hys_city_proto.HysCity{}
	if hysCityEntity == nil {
		result.Data = nil
	} else {
		convert.EntityToEntity(hysCityEntity, &hysCityOut)
		result.Data = &hysCityOut

	}
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = ""
	return result, nil

}
func (instance *HysCityController) GetHysCityList(ctx context.Context, req *hys_city_proto.Query) (result *hys_city_proto.ListResult, err error) {
	result = new(hys_city_proto.ListResult)
	defer func() {
		if err := recover(); err != nil {
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
	service := hys_city_service.HysCityService{}
	hysCityList := service.GetHysCityList(*req)
	hysCityOut := []*hys_city_proto.HysCity{}
	for _, hysCityEntity := range hysCityList {
		hysCityTemp := hys_city_proto.HysCity{}
		convert.EntityToEntity(hysCityEntity, &hysCityTemp)
		hysCityOut = append(hysCityOut, &hysCityTemp)

	}
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = ""
	result.Data = hysCityOut
	return result, nil
}
func (instance *HysCityController) DeleteHysCity(ctx context.Context, req *hys_city_proto.Query) (result *hys_city_proto.Result, err error) {
	result = new(hys_city_proto.Result)
	defer func() {
		if err := recover(); err != nil {
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
	service := hys_city_service.HysCityService{}
	if err = service.DeleteHysCity(*req); err != nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = err.Error()
		return result, err
	}
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = "delete success"
	return result, err
}
func (instance *HysCityController) GetHysCityPageList(ctx context.Context, req *hys_city_proto.PageQuery) (result *hys_city_proto.PageResult, err error) {
	result = new(hys_city_proto.PageResult)
	defer func() {
		if err := recover(); err != nil {
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
	service := hys_city_service.HysCityService{}
	hysCityList := service.GetHysCityPageList(*req.Query, int(req.Page), int(req.PageNum))
	hysCityOut := []*hys_city_proto.HysCity{}
	hysCityPageData := hys_city_proto.HysCityPageData{}
	for _, hysCityEntity := range hysCityList {
		hysCityTemp := hys_city_proto.HysCity{}
		convert.EntityToEntity(hysCityEntity, &hysCityTemp)
		hysCityOut = append(hysCityOut, &hysCityTemp)

	}
	hysCityPageData.List = hysCityOut
	hysCityPageData.CurrentPage = req.Page
	hysCityPageData.Count = int32(service.GetHysCityCount(*req.Query))
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = ""
	result.Data = &hysCityPageData
	return result, nil
}
func (instance *HysCityController) GetHysCityCount(ctx context.Context, req *hys_city_proto.Query) (result *hys_city_proto.Result, err error) {
	result = new(hys_city_proto.Result)
	defer func() {
		if err := recover(); err != nil {
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
	service := hys_city_service.HysCityService{}
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = ""
	result.Data = strconv.Itoa(service.GetHysCityCount(*req))
	return result, nil
}
