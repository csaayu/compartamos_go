package model

import (
	"time"
)

type Users struct {
	Id              int    `gorm:"type:int;primary_key"`
	Dni             int    `gorm:"type:int"`
	Name            string `gorm:"type:varchar(255)"`
	LastName        string `gorm:"type:varchar(255)"`
	FechaNacimiento time.Time
	Sexo            string `gorm:"type:varchar(10)"`
	Ciudad          string `gorm:"type:varchar(255)"`
}
