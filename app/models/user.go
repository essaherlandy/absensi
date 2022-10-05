package models

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name" validate:"required,min=3,max=25"`
	Email    string `json:"email" gorm:"unique" validate:"required,email"`
	Divisi   string `json:"divisi" validate:"required"`
	Alamat   string `json:"alamat" validate:"required"`
	Password []byte `json:"-"`
}
