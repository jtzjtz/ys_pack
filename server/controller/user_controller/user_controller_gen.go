package user_controller

import (
	"context"
	"github.com/jtzjtz/kit/convert"
	"github.com/jtzjtz/ys_pack/entity"
	"github.com/jtzjtz/ys_pack/proto/user_proto"
	"github.com/jtzjtz/ys_pack/server/service/user_service"
	"strconv"
)

type UserController struct {
}

func (instance *UserController) CreateUser(ctx context.Context, req *user_proto.User) (result *user_proto.EntityResult, err error) {
	result = new(user_proto.EntityResult)
	defer func() {
		if err := recover(); err != nil {
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
	userEntity := entity.User{}
	err = convert.EntityToEntity(req, &userEntity)
	if err != nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = err.Error()
		return result, err
	}
	service := user_service.UserService{}
	userNew, errCreate := service.CreateUser(userEntity)
	if errCreate != nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = errCreate.Error()
		return result, errCreate
	}
	userOut := user_proto.User{}
	convert.EntityToEntity(userNew, &userOut)
	result.Data = &userOut
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = "create success"
	return result, err
}
func (instance *UserController) UpdateUser(ctx context.Context, req *user_proto.UpdateAndCondition) (result *user_proto.Result, err error) {
	result = new(user_proto.Result)
	defer func() {
		if err := recover(); err != nil {
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
	service := user_service.UserService{}
	if req.Entity == nil {
		req.Entity = &user_proto.User{}
	}
	entityUpdate := convert.EntityToMapWithEmpty(req.Entity, req.UpdateEmptyFields)
	if err = service.UpdateUser(*req.Query, entityUpdate); err != nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = err.Error()
		return result, err
	}
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = "create success"
	return result, err
}
func (instance *UserController) GetUser(ctx context.Context, req *user_proto.Query) (result *user_proto.EntityResult, err error) {
	result = new(user_proto.EntityResult)
	defer func() {
		if err := recover(); err != nil {
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
	service := user_service.UserService{}
	userEntity := service.GetUser(*req)
	userOut := user_proto.User{}
	if userEntity == nil {
		result.Data = nil
	} else {
		convert.EntityToEntity(userEntity, &userOut)
		result.Data = &userOut

	}
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = ""
	return result, nil

}
func (instance *UserController) GetUserList(ctx context.Context, req *user_proto.Query) (result *user_proto.ListResult, err error) {
	result = new(user_proto.ListResult)
	defer func() {
		if err := recover(); err != nil {
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
	service := user_service.UserService{}
	userList := service.GetUserList(*req)
	userOut := []*user_proto.User{}
	for _, userEntity := range userList {
		userTemp := user_proto.User{}
		convert.EntityToEntity(userEntity, &userTemp)
		userOut = append(userOut, &userTemp)

	}
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = ""
	result.Data = userOut
	return result, nil
}
func (instance *UserController) DeleteUser(ctx context.Context, req *user_proto.Query) (result *user_proto.Result, err error) {
	result = new(user_proto.Result)
	defer func() {
		if err := recover(); err != nil {
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
	service := user_service.UserService{}
	if err = service.DeleteUser(*req); err != nil {
		result.Code = entity.RESULTERRORINT32
		result.Msg = err.Error()
		return result, err
	}
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = "delete success"
	return result, err
}
func (instance *UserController) GetUserPageList(ctx context.Context, req *user_proto.PageQuery) (result *user_proto.PageResult, err error) {
	result = new(user_proto.PageResult)
	defer func() {
		if err := recover(); err != nil {
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
	service := user_service.UserService{}
	userList := service.GetUserPageList(*req.Query, int(req.Page), int(req.PageNum))
	userOut := []*user_proto.User{}
	userPageData := user_proto.UserPageData{}
	for _, userEntity := range userList {
		userTemp := user_proto.User{}
		convert.EntityToEntity(userEntity, &userTemp)
		userOut = append(userOut, &userTemp)

	}
	userPageData.List = userOut
	userPageData.CurrentPage = req.Page
	userPageData.Count = int32(service.GetUserCount(*req.Query))
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = ""
	result.Data = &userPageData
	return result, nil
}
func (instance *UserController) GetUserCount(ctx context.Context, req *user_proto.Query) (result *user_proto.Result, err error) {
	result = new(user_proto.Result)
	defer func() {
		if err := recover(); err != nil {
			result.Code = entity.RESULTERRORINT32
			result.Msg = "参数异常！"
		}
	}()
	service := user_service.UserService{}
	result.Code = entity.RESULTSUCCESSINT32
	result.Msg = ""
	result.Data = strconv.Itoa(service.GetUserCount(*req))
	return result, nil
}
