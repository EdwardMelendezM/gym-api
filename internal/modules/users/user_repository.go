package users

type Repository interface {
	GetAll() ([]User, error)
	Create(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindById(id string) (User, error)
}
