package repository

import (
	"compartamos_go/data/request"
	"compartamos_go/helper"
	"compartamos_go/model"
	"errors"

	"gorm.io/gorm"
)

type UsersRepositoryImpl struct {
	Db *gorm.DB
}

func NewUsersRePositoryImpl(Db *gorm.DB) UsersRepository {
	return &UsersRepositoryImpl{Db: Db}
}

func (t *UsersRepositoryImpl) Delete(userId int) {
	var users model.Users
	result := t.Db.Where("id = ?", userId).Delete(&users)
	helper.ErrorSync(result.Error)
}

func (t *UsersRepositoryImpl) FindAll() []model.Users {
	var users []model.Users
	result := t.Db.Find(&users)
	helper.ErrorSync(result.Error)
	return users
}

func (t *UsersRepositoryImpl) FindById(usersId int) (users model.Users, err error) {
	var user model.Users
	result := t.Db.Find(&user, usersId)
	if result != nil {
		return user, nil
	} else {
		return user, errors.New("usuario no encontrados")
	}
}

func (t *UsersRepositoryImpl) Save(users model.Users) {
	result := t.Db.Create(&users)
	helper.ErrorSync(result.Error)
}

func (t *UsersRepositoryImpl) Update(users model.Users) {
	var updateUsers = request.UpdateUsersRequest{
		Id:              users.Id,
		Dni:             users.Dni,
		Name:            users.Name,
		LastName:        users.LastName,
		FechaNacimiento: users.FechaNacimiento,
		Sexo:            users.Sexo,
		Ciudad:          users.Ciudad,
	}

	result := t.Db.Model(&users).Updates(updateUsers)

	helper.ErrorSync(result.Error)
}
