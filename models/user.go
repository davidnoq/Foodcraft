package models

type User struct {
	Password string `json:"password"`
	Username string `json:"username"`

	Recipe []struct {
		ID int `bson:"id"`
	}
}
