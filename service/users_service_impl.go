package service

import (
	"compartamos_go/data/request"
	"compartamos_go/data/response"
	"compartamos_go/helper"
	"compartamos_go/model"
	"compartamos_go/repository"

	"github.com/go-playground/validator/v10"
)

type UsersServiceImpl struct {
	UsersRepository repository.UsersRepository
	Validate        *validator.Validate
}

func NewUsersServiceImpl(userRepository repository.UsersRepository, validate *validator.Validate) UsersService {
	return &UsersServiceImpl{
		UsersRepository: userRepository,
		Validate:        validate,
	}
}

// Create implements UsersService.
func (t *UsersServiceImpl) Create(users request.CreateUsersRequest) {
	err := t.Validate.Struct(users)
	helper.ErrorSync(err)
	userModel := model.Users{
		Dni:             users.Dni,
		Name:            users.Name,
		LastName:        users.LastName,
		FechaNacimiento: users.FechaNacimiento,
		Sexo:            users.Sexo,
		Ciudad:          users.Ciudad,
	}
	t.UsersRepository.Save(userModel)
}

// Delete implements UsersService.
func (t *UsersServiceImpl) Delete(usersId int) {
	t.UsersRepository.Delete(usersId)

}

// FindAll implements UsersService.
func (t *UsersServiceImpl) FindAll() []response.UsersResponse {
	result := t.UsersRepository.FindAll()

	var userS []response.UsersResponse

	for _, value := range result {
		userO := response.UsersResponse{
			Id:              value.Id,
			Dni:             value.Dni,
			Name:            value.Name,
			LastName:        value.LastName,
			FechaNacimiento: value.FechaNacimiento,
			Sexo:            value.Sexo,
			Ciudad:          value.Ciudad,
		}
		userS = append(userS, userO)
	}

	return userS
}

// FindById implements UsersService.
func (t *UsersServiceImpl) FindById(usersId int) response.UsersResponse {

	userData, err := t.UsersRepository.FindById(usersId)

	helper.ErrorSync(err)

	userResponse := response.UsersResponse{
		Id:              userData.Id,
		Dni:             userData.Dni,
		Name:            userData.Name,
		LastName:        userData.LastName,
		FechaNacimiento: userData.FechaNacimiento,
		Sexo:            userData.Sexo,
		Ciudad:          userData.Ciudad,
	}

	return userResponse

}

// Update implements UsersService.
func (t *UsersServiceImpl) Update(users request.UpdateUsersRequest) {

	userData, err := t.UsersRepository.FindById(users.Id)
	helper.ErrorSync(err)
	userData.Dni = users.Dni
	userData.Name = users.Name
	userData.LastName = users.LastName
	userData.FechaNacimiento = users.FechaNacimiento
	userData.Sexo = users.Sexo
	userData.Ciudad = users.Ciudad

	t.UsersRepository.Update(userData)

}
