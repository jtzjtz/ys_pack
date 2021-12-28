package user_grpc

import (
	"context"
	"github.com/jtzjtz/ys_pack/entity"
	"github.com/jtzjtz/ys_pack/proto/user_proto"
	"google.golang.org/grpc"
)

type UserService struct {
}

func (t *UserService) CreateUser(ctx context.Context, userEntity *user_proto.User, conn *grpc.ClientConn) (result *user_proto.EntityResult) {
	result = &user_proto.EntityResult{Code: entity.RESULTERRORINT32}
	if userEntity == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件参数错误"
	} else {
		client := user_proto.NewUserServiceClient(conn)
		resultReponse, err := client.CreateUser(ctx, userEntity)
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

func (t *UserService) UpdateUser(ctx context.Context, orderUpdateAndCondition *user_proto.UpdateAndCondition, conn *grpc.ClientConn) (result *user_proto.Result) {
	result = &user_proto.Result{Code: entity.RESULTERRORINT32}
	if orderUpdateAndCondition.Entity == nil || orderUpdateAndCondition.Query == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询query错误"
	} else if orderUpdateAndCondition.Query.EntityQuery == nil && orderUpdateAndCondition.Query.SqlQuery == "" {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询SqlQuery错误"
	} else {
		client := user_proto.NewUserServiceClient(conn)
		resultReponse, err := client.UpdateUser(ctx, orderUpdateAndCondition)
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

func (t *UserService) GetUser(ctx context.Context, query *user_proto.Query, conn *grpc.ClientConn) (result *user_proto.EntityResult) {
	result = &user_proto.EntityResult{Code: entity.RESULTERRORINT32}
	if query == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询query错误"
	} else if query.EntityQuery == nil && query.SqlQuery == "" {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询SqlQuery错误"
	} else {
		client := user_proto.NewUserServiceClient(conn)
		resultReponse, err := client.GetUser(ctx, query)
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

func (t *UserService) GetUserList(ctx context.Context, query *user_proto.Query, conn *grpc.ClientConn) (result *user_proto.ListResult) {
	result = &user_proto.ListResult{Code: entity.RESULTERRORINT32}
	if query == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询query错误"
	} else if query.EntityQuery == nil && query.SqlQuery == "" {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询SqlQuery错误"
	} else {
		client := user_proto.NewUserServiceClient(conn)
		resultReponse, err := client.GetUserList(ctx, query)
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

func (this *UserService) DeleteUser(ctx context.Context, query *user_proto.Query, conn *grpc.ClientConn) (result *user_proto.Result) {
	result = &user_proto.Result{Code: entity.RESULTERRORINT32}
	if query == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询query错误"
	} else if query.EntityQuery == nil && query.SqlQuery == "" {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询SqlQuery错误"
	} else {
		client := user_proto.NewUserServiceClient(conn)
		resultReponse, err := client.DeleteUser(ctx, query)
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

func (t *UserService) GetUserPageList(ctx context.Context, pageQuery *user_proto.PageQuery, conn *grpc.ClientConn) (result *user_proto.PageResult) {
	result = &user_proto.PageResult{Code: entity.RESULTERRORINT32}
	if pageQuery == nil || pageQuery.Query == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询query错误"
	} else if pageQuery.Query.EntityQuery == nil && pageQuery.Query.SqlQuery == "" {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询SqlQuery错误"

	} else {
		client := user_proto.NewUserServiceClient(conn)
		resultReponse, err := client.GetUserPageList(ctx, pageQuery)
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

func (t *UserService) GetUserCount(ctx context.Context, query *user_proto.Query, conn *grpc.ClientConn) (result *user_proto.Result) {
	result = &user_proto.Result{Code: entity.RESULTERRORINT32}
	if query == nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询query错误"
	} else if query.EntityQuery == nil && query.SqlQuery == "" {
		result.Code = entity.RESULTERRORINT32
		result.Msg = "条件查询SqlQuery错误"

	} else {
		client := user_proto.NewUserServiceClient(conn)
		resultReponse, err := client.GetUserCount(ctx, query)
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
