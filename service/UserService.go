package service

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"grpc-demo/global"
	model "grpc-demo/model/system"
	"grpc-demo/proto"
	"grpc-demo/tools"
	"log"
)

type userService struct {
}

func NewUserService() *userService {
	return &userService{}
}

func (u *userService) Login(ctx context.Context, request *proto.LoginRequest) (*proto.LoginResponse, error) {
	username, password := request.Username, request.Password
	log.Printf("username: %s\t password: %s", request.Username, request.Password)
	if len(username) == 0 || len(password) == 0 {
		return nil, errors.New("username and password not empty")
	}
	var user model.User
	result := global.DB.Where("username = ?", username).Find(&user)
	if result.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}
	if ok := tools.BcryptCheck(password, user.Password); !ok {
		return nil, errors.New("password is error")
	}

	token, err := tools.GenerateToken(username, uint64(user.ID))
	if err != nil {
		log.Println(err)
		return nil, errors.New("generate token error")
	}
	// 1.校验合法
	return &proto.LoginResponse{
		Username: user.Username,
		ImgUrl:   "http://localhost",
		Token:    token,
	}, nil
}

func (u *userService) Register(ctx context.Context, request *proto.RegisterRequest) (response *proto.RegisterResponse, err error) {
	username := request.GetUsername()
	if len(username) == 0 {
		return nil, errors.New("username is empty")
	}
	if isExistByUsername(username) {
		return nil, errors.New(fmt.Sprintf("%s already exists", username))
	}
	var bcryptPassword string
	if bcryptPassword, err = checkAndBcryptPassword(request.GetPassword(), request.GetPassword2()); err != nil {
		return nil, err
	}
	user := &model.User{
		Username: username,
		Password: bcryptPassword,
	}
	if err = global.DB.Save(&user).Error; err != nil {
		log.Println(err)
		return nil, errors.New("register failed")
	}
	response = &proto.RegisterResponse{
		Username: username,
		Id:       uint64(user.ID),
	}
	return response, nil
}

func isExistByUsername(username string) bool {
	err := global.DB.Where("username = ?", username).First(&model.User{}).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

func checkAndBcryptPassword(password, password2 string) (string, error) {
	if len(password) == 0 || len(password2) == 0 {
		return "", errors.New("password is empty")
	}
	if password != password2 {
		return "", errors.New("twice password is not equals")
	}
	return tools.BcryptHash(password), nil
}
