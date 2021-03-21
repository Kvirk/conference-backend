package models

import (
	"net/http"

	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	Id        int    `json:"id"`
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"firstName" validate:"required`
	LastName  string `json:"lastName"  validate:"required`
	Password  string `json:"password"  validate:"required,min=8,max=128"`
	CreatedAt string `json:"createdAt"`
}

type UserResponse struct {
	Id        int    `json:"id"`
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"firstName" validate:"required`
	LastName  string `json:"lastName"  validate:"required`
	CreatedAt string `json:"createdAt"`
}

func (i *User) Bind(r *http.Request) error {
	v := validator.New()

	err := v.Struct(i)

	if err != nil {
		return err
	}

	return nil
}

func (i *User) ConvertToResponse() (UserResponse, error) {
	userReponse := UserResponse{}
	userReponse.Id = i.Id
	userReponse.Email = i.Email
	userReponse.FirstName = i.FirstName
	userReponse.LastName = i.LastName
	userReponse.CreatedAt = i.CreatedAt

	v := validator.New()

	err := v.Struct(userReponse)

	if err != nil {
		return userReponse, err
	}

	return userReponse, nil
}

func (*UserResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
