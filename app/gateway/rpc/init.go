package rpc

import (
	"gin-gorm-memo/v2/idl/pb"
	"go-micro.dev/v4"
)

var (
	UserService pb.UserService
)

func InitRPC() {
	userMicroService := micro.NewService(micro.Name("userService.client"))
	userService := pb.NewUserService("rpcUserService", userMicroService.Client())
	UserService = userService
}
