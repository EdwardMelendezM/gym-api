package users

import "time"

type User struct {
	ID        string
	FullName  *string
	FirstName *string
	LastName  *string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
