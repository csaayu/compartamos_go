package repository

import "compartamos_go/model"

type UsersRepository interface {
	Save(users model.Users)
	Update(users model.Users)
	Delete(userId int)
	FindById(usersId int) (users model.Users, err error)
	FindAll() []model.Users
}
