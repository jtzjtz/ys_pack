package sq_user_blacklist_grpc

import (
	"context"
	"github.com/jtzjtz/ys_pack/entity"
	"github.com/jtzjtz/ys_pack/proto/sq_user_blacklist_proto"
	"google.golang.org/grpc"
)

type SqUserBlacklistService struct {
}

func (t *SqUserBlacklistService) CreateSqUserBlacklist(ctx context.Context, sqUserBlacklistEntity *sq_user_blacklist_proto.SqUserBlacklist, conn *grpc.ClientConn) (result *sq_user_blacklist_proto.EntityResult) {
	result = &sq_user_blacklist_proto.EntityResult{Code: entity.RESULTERRORINT32}
	if sqUserBlacklistEntity == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件参数错误"
	} else {
		client := sq_user_blacklist_proto.NewSqUserBlacklistServiceClient(conn)
		resultReponse, err := client.CreateSqUserBlacklist(ctx, sqUserBlacklistEntity)
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

func (t *SqUserBlacklistService) UpdateSqUserBlacklist(ctx context.Context, orderUpdateAndCondition *sq_user_blacklist_proto.UpdateAndCondition, conn *grpc.ClientConn) (result *sq_user_blacklist_proto.Result) {
	result = &sq_user_blacklist_proto.Result{Code: entity.RESULTERRORINT32}
	if orderUpdateAndCondition.Entity == nil || orderUpdateAndCondition.Query == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询query错误"
	} else if orderUpdateAndCondition.Query.EntityQuery == nil && orderUpdateAndCondition.Query.SqlQuery == "" {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询SqlQuery错误"
	} else {
		client := sq_user_blacklist_proto.NewSqUserBlacklistServiceClient(conn)
		resultReponse, err := client.UpdateSqUserBlacklist(ctx, orderUpdateAndCondition)
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

func (t *SqUserBlacklistService) GetSqUserBlacklist(ctx context.Context, query *sq_user_blacklist_proto.Query, conn *grpc.ClientConn) (result *sq_user_blacklist_proto.EntityResult) {
	result = &sq_user_blacklist_proto.EntityResult{Code: entity.RESULTERRORINT32}
	if query == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询query错误"
	} else if query.EntityQuery == nil && query.SqlQuery == "" {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询SqlQuery错误"
	} else {
		client := sq_user_blacklist_proto.NewSqUserBlacklistServiceClient(conn)
		resultReponse, err := client.GetSqUserBlacklist(ctx, query)
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

func (t *SqUserBlacklistService) GetSqUserBlacklistList(ctx context.Context, query *sq_user_blacklist_proto.Query, conn *grpc.ClientConn) (result *sq_user_blacklist_proto.ListResult) {
	result = &sq_user_blacklist_proto.ListResult{Code: entity.RESULTERRORINT32}
	if query == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询query错误"
	} else if query.EntityQuery == nil && query.SqlQuery == "" {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询SqlQuery错误"
	} else {
		client := sq_user_blacklist_proto.NewSqUserBlacklistServiceClient(conn)
		resultReponse, err := client.GetSqUserBlacklistList(ctx, query)
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

func (this *SqUserBlacklistService) DeleteSqUserBlacklist(ctx context.Context, query *sq_user_blacklist_proto.Query, conn *grpc.ClientConn) (result *sq_user_blacklist_proto.Result) {
	result = &sq_user_blacklist_proto.Result{Code: entity.RESULTERRORINT32}
	if query == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询query错误"
	} else if query.EntityQuery == nil && query.SqlQuery == "" {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询SqlQuery错误"
	} else {
		client := sq_user_blacklist_proto.NewSqUserBlacklistServiceClient(conn)
		resultReponse, err := client.DeleteSqUserBlacklist(ctx, query)
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

func (t *SqUserBlacklistService) GetSqUserBlacklistPageList(ctx context.Context, pageQuery *sq_user_blacklist_proto.PageQuery, conn *grpc.ClientConn) (result *sq_user_blacklist_proto.PageResult) {
	result = &sq_user_blacklist_proto.PageResult{Code: entity.RESULTERRORINT32}
	if pageQuery == nil || pageQuery.Query == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询query错误"
	} else if pageQuery.Query.EntityQuery == nil && pageQuery.Query.SqlQuery == "" {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询SqlQuery错误"

	} else {
		client := sq_user_blacklist_proto.NewSqUserBlacklistServiceClient(conn)
		resultReponse, err := client.GetSqUserBlacklistPageList(ctx, pageQuery)
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

func (t *SqUserBlacklistService) GetSqUserBlacklistCount(ctx context.Context, query *sq_user_blacklist_proto.Query, conn *grpc.ClientConn) (result *sq_user_blacklist_proto.Result) {
	result = &sq_user_blacklist_proto.Result{Code: entity.RESULTERRORINT32}
	if query == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询query错误"
	} else if query.EntityQuery == nil && query.SqlQuery == "" {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询SqlQuery错误"

	} else {
		client := sq_user_blacklist_proto.NewSqUserBlacklistServiceClient(conn)
		resultReponse, err := client.GetSqUserBlacklistCount(ctx, query)
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
