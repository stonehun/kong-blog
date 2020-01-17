package repositories

import (
	"kong_blog/datasource"
	"kong_blog/models"
)

type UserRepository interface {
	GetUserByName(username string) (user models.User)
	GetUserById(id uint) (user models.User)
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

type userRepository struct{}

//登录
func (n userRepository) GetUserByName(name string) (user models.User) {
	db := datasource.GetDB()
	db.Where("name = ?", name).First(&user)
	return
}

//根据id查询用户
func (n userRepository) GetUserById(id uint) (user models.User) {
	db := datasource.GetDB()
	db.Where("id = ?", id).First(&user)
	return
}
