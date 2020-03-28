package models

import "github.com/jinzhu/gorm"

//Tecnico es un modelo para la entidad TECNICOS
//de nuestra base de datos SCA y para poder realizar
//las consulta DML.
type Tecnico struct {
	gorm.Model
	ID       string `json:"id" gorm:"not null"`
	Nombre   string `json:"nombre" gorm:"not null"`
	Apaterno string `json:"apaterno" gorm:"not null"`
	Amaterno string `json:"amaterno" gorm:"not null"`
	Usuario  string `json:"usuario"`
	Email    string `json:"email"`
	Celular  uint16 `json:"celular"`
	Pin      string `json:"pin,omitempty"`
}
