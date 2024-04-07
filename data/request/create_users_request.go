package request

import "time"

type CreateUsersRequest struct {
	Dni             int       `validate:"required" json:"dni"`
	Name            string    `validate:"required" json:"name"`
	LastName        string    `validate:"required" json:"lastName"`
	FechaNacimiento time.Time `validate:"required" json:"fechaNacimiento"`
	Sexo            string    `validate:"required" json:"sexo"`
	Ciudad          string    `validate:"required" json:"ciudad"`
}
