package users

type Repository interface {
	GetAll() ([]User, error)
	Create(user User) (User, error)
}

type memoryRepository struct {
	data []User
}

func NewRepository() Repository {
	return &memoryRepository{
		data: []User{},
	}
}

func (r *memoryRepository) GetAll() ([]User, error) {
	return r.data, nil
}

func (r *memoryRepository) Create(user User) (User, error) {
	r.data = append(r.data, user)
	return user, nil
}
