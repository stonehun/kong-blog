package services

import (
	"errors"
	"github.com/iris-contrib/middleware/jwt"
	"kong_blog/configs"
	"kong_blog/repositories"
	"kong_blog/util"
)

type UserService interface {
	Login(name string, password string) (token string, err error)
}

type userService struct {
	UserRepository repositories.UserRepository
}

func NewUserServices() UserService {
	return &userService{
		UserRepository: repositories.NewUserRepository(),
	}
}

//用户登录逻辑
func (u *userService) Login(name string, password string) (token string, err error) {
	//查询用户
	user := u.UserRepository.GetUserByName(name)
	if user.ID == 0 {
		err = errors.New("未找到用户")
		return
	}

	//密码是否正确
	if user.Password != util.MdString(password) {
		err = errors.New("密码不正确")
		return
	}

	//生成token
	var mySecret = []byte(configs.GetConfig().JwtKey)
	tokenOb := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":      user.Name,
		"id":        user.ID,
		"nick_name": user.NickName,
	})

	token, err = tokenOb.SignedString(mySecret)
	return
}
