package models

type User struct {
	Password string `json:"password"`
	Username string `json:"username"`

	ID       string `bson:"_id,omitempty" json:"id"`
}
