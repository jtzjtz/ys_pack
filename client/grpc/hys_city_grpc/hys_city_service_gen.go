package hys_city_grpc

import (
	"context"
	"github.com/jtzjtz/ys_pack/entity"
	"github.com/jtzjtz/ys_pack/proto/hys_city_proto"
	"google.golang.org/grpc"
)

type HysCityService struct {
}

func (t *HysCityService) CreateHysCity(ctx context.Context, hysCityEntity *hys_city_proto.HysCity, conn *grpc.ClientConn) (result *hys_city_proto.EntityResult) {
	result = &hys_city_proto.EntityResult{Code: entity.RESULTERRORINT32}
	if hysCityEntity == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件参数错误"
	} else {
		client := hys_city_proto.NewHysCityServiceClient(conn)
		resultReponse, err := client.CreateHysCity(ctx, hysCityEntity)
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

func (t *HysCityService) UpdateHysCity(ctx context.Context, orderUpdateAndCondition *hys_city_proto.UpdateAndCondition, conn *grpc.ClientConn) (result *hys_city_proto.Result) {
	result = &hys_city_proto.Result{Code: entity.RESULTERRORINT32}
	if orderUpdateAndCondition.Entity == nil || orderUpdateAndCondition.Query == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询query错误"
	} else if orderUpdateAndCondition.Query.EntityQuery == nil && orderUpdateAndCondition.Query.SqlQuery == "" {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询SqlQuery错误"
	} else {
		client := hys_city_proto.NewHysCityServiceClient(conn)
		resultReponse, err := client.UpdateHysCity(ctx, orderUpdateAndCondition)
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

func (t *HysCityService) GetHysCity(ctx context.Context, query *hys_city_proto.Query, conn *grpc.ClientConn) (result *hys_city_proto.EntityResult) {
	result = &hys_city_proto.EntityResult{Code: entity.RESULTERRORINT32}
	if query == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询query错误"
	} else if query.EntityQuery == nil && query.SqlQuery == "" {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询SqlQuery错误"
	} else {
		client := hys_city_proto.NewHysCityServiceClient(conn)
		resultReponse, err := client.GetHysCity(ctx, query)
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

func (t *HysCityService) GetHysCityList(ctx context.Context, query *hys_city_proto.Query, conn *grpc.ClientConn) (result *hys_city_proto.ListResult) {
	result = &hys_city_proto.ListResult{Code: entity.RESULTERRORINT32}
	if query == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询query错误"
	} else if query.EntityQuery == nil && query.SqlQuery == "" {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询SqlQuery错误"
	} else {
		client := hys_city_proto.NewHysCityServiceClient(conn)
		resultReponse, err := client.GetHysCityList(ctx, query)
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

func (this *HysCityService) DeleteHysCity(ctx context.Context, query *hys_city_proto.Query, conn *grpc.ClientConn) (result *hys_city_proto.Result) {
	result = &hys_city_proto.Result{Code: entity.RESULTERRORINT32}
	if query == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询query错误"
	} else if query.EntityQuery == nil && query.SqlQuery == "" {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询SqlQuery错误"
	} else {
		client := hys_city_proto.NewHysCityServiceClient(conn)
		resultReponse, err := client.DeleteHysCity(ctx, query)
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

func (t *HysCityService) GetHysCityPageList(ctx context.Context, pageQuery *hys_city_proto.PageQuery, conn *grpc.ClientConn) (result *hys_city_proto.PageResult) {
	result = &hys_city_proto.PageResult{Code: entity.RESULTERRORINT32}
	if pageQuery == nil || pageQuery.Query == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询query错误"
	} else if pageQuery.Query.EntityQuery == nil && pageQuery.Query.SqlQuery == "" {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询SqlQuery错误"

	} else {
		client := hys_city_proto.NewHysCityServiceClient(conn)
		resultReponse, err := client.GetHysCityPageList(ctx, pageQuery)
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

func (t *HysCityService) GetHysCityCount(ctx context.Context, query *hys_city_proto.Query, conn *grpc.ClientConn) (result *hys_city_proto.Result) {
	result = &hys_city_proto.Result{Code: entity.RESULTERRORINT32}
	if query == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询query错误"
	} else if query.EntityQuery == nil && query.SqlQuery == "" {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询SqlQuery错误"

	} else {
		client := hys_city_proto.NewHysCityServiceClient(conn)
		resultReponse, err := client.GetHysCityCount(ctx, query)
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
