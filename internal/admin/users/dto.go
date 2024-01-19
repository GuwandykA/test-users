package users

import "time"

type DataDTO struct {
	Users []UserDTO `json:"users"`
	Count int       `json:"count"`
}

type PaginationDTO struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type UserDTO struct {
	Id          int         `json:"Id"`
	Name        string      `json:"name"`
	Surname     string      `json:"surname"`
	Patronymic  string      `json:"patronymic"`
	Age         int         `json:"age"`
	Gender      string      `json:"gender"`
	Probability float32     `json:"probability"`
	Country     interface{} `json:"country"`
	CreatedAt   time.Time   `json:"createdAt"`
}

type ReqUser struct {
	Id         int    `json:"Id"`
	Name       string `json:"name" binding:"required"`
	Surname    string `json:"surname" binding:"required"`
	Patronymic string `json:"patronymic"`
}

type ResUserAge struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
	Age   int    `json:"age"`
}

type ResUserGender struct {
	Id          int     `json:"id"`
	Gender      string  `json:"gender"`
	Count       int     `json:"count"`
	Probability float32 `json:"probability"`
}

type ResUserCountry struct {
	Id      int         `json:"id"`
	Name    string      `json:"name"`
	Count   int         `json:"count"`
	Country interface{} `json:"country"`
}

type ReqIdDTO struct {
	Id int `json:"Id"`
}
