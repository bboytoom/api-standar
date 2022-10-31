package models

import (
	"fmt"
	"time"
)

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name" validate:"required,min=3,max=50"`
	LastName  string    `json:"last_name" validate:"required,min=2,max=50"`
	Email     string    `json:"email" validate:"email,required,min=8,max=70"`
	Password  string    `json:"password,omitempty" validate:"required,min=8,max=20"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type ReturnUser struct {
	Result string `json:"result"`
	Id     int    `json:"id"`
}

type ResponseUser struct {
	Data []User `json:"data"`
}

func RetrieveUsers() *ResponseUser {
	all_users := []User{
		{
			Id:       1,
			Name:     "Test 1",
			LastName: "LastTest 1",
			Email:    "test1@prueba.com",
			Status:   true,
		},
		{
			Id:       2,
			Name:     "Test 2",
			LastName: "LastTest 2",
			Email:    "test2@prueba.com",
			Status:   true,
		},
	}

	return &ResponseUser{
		Data: all_users,
	}
}

func AddOnlyUser(u User) {
	fmt.Println(u)
}
