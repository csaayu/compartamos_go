package response

import "time"

type UsersResponse struct {
	Id              int       `json:"id"`
	Dni             int       `json: "dni"`
	Name            string    `json:name`
	LastName        string    `json:lastName`
	FechaNacimiento time.Time `json:fechaNacimiento`
	Sexo            string    `json:sexo`
	Ciudad          string    `json:ciudad`
}
