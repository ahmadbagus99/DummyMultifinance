package models

import "time"

type Users struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	RoleID    string    `json:"role_id"` // 'System Administrator' or 'Employee'
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
