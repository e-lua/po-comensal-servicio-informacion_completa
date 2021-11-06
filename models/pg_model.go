package models

import "time"

type Pg_Comensal struct {
	IdComensal     int       `json:"id"`
	IdCountry      int       `json:"country"`
	Phone          int       `json:"phone"`
	UrlPerfil      string    `json:"urlPerfil"`
	Name           string    `json:"name"`
	LastName       string    `json:"lastName"`
	Password       string    `json:"password"`
	Email          string    `json:"email"`
	IsAvailable    string    `json:"isAvailable"`
	CreatedDate    time.Time `json:"createdDate"`
	UpdatedDate    time.Time `json:"updatedDate"`
	DeletedDate    time.Time `json:"deletedDate"`
	OrdersRejected int       `json:"orderesRejected"`
}
