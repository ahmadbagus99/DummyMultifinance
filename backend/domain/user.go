package domain

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name" form:"name"`
	Username  string `json:"username" form:"username"`
	Password  string `json:"password" form:"password"`
	Email     string `json:"email" form:"email"`
	RoleID    int    `json:"role_id" form:"role_id"`
	UpdatedAt string `json:"updated_at" form:"updated_at"`
	CreatedAt string `json:"created_at" form:"created_at"`
}

type UserRepository interface {
	CreateUser(user *User) (*User, error)
	GetUserByUsername(username string) (*User, error)
}
