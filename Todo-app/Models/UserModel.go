package Models

import "time"

type Users struct {
	UserId    string    `json:"userid" db:"userid"`
	UserName  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"createdAt"`
	UpdatedAt time.Time `json:"updated_at" db:"updatedAt"`
	Todos     Todos
}
