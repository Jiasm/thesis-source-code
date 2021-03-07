package entity

type User struct {
	ID int `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	RoleId int `json:"role_id"`
	Status int `json:"status"`
}

