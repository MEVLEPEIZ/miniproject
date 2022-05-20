package models

import (
"gorm.io/gorm")

type User struct {
	gorm.Model
	ID int `json:"id"`
	Nama string `json:"nama"`
	Email string `json:"email"`
	Password string `json:"password"`
	Kondisi_kulit string `json:"kondisi_kulit"`
	
}
type UsersResponse struct {
	ID int `json:"id"`
	Nama string `json:"nama"`
	Email string `json:"email"`
	Password string `json:"password"`
	Kondisi_kulit string `json:"kondisi_kulit"`
	Token  string `json:"token" form:"token"`
}