package handler

import (
	"context"
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"goShop/srv_user/global"
	"goShop/srv_user/model"
	"goShop/srv_user/proto"
	pb "goShop/srv_user/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserServer struct {
	pb.UnimplementedUserServer
}

func ModelToResponse(user model.User) proto.UserInfoResponse {
	userInfoResponse := proto.UserInfoResponse{
		Id:       user.ID,
		PassWord: user.Password,
		Mobile:   user.Mobile,
		NickName: user.NickName,
		Gender:   user.Gender,
		Role:     int32(user.Role),
	}

	if user.Birthday != nil {
		userInfoResponse.BirthDay = uint64(user.Birthday.Unix())
	}

	return userInfoResponse
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize < 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (s *UserServer) CreateUser(ctx context.Context, req *proto.CreateUserInfo) (*proto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.Where(&model.User{Mobile: req.Mobile}).First(&user)
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "user already exists")
	}

	options := password.Options{
		SaltLen:      16,
		Iterations:   100,
		KeyLen:       32,
		HashFunction: sha512.New,
	}

	salt, enPwd := password.Encode(req.PassWord, &options)
	user.Password = fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, enPwd)

	result = global.DB.Create(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}

	response := ModelToResponse(user)

	return &response, nil
}
