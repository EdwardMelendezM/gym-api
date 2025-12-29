package users

import (
	"context"
	"gym-api/internal/ent"
	"gym-api/internal/ent/user"
)

type entRepository struct {
	client *ent.Client
}

func NewUserEntRepository(client *ent.Client) Repository {
	return &entRepository{client: client}
}

func (r *entRepository) GetAll() ([]User, error) {
	rows, err := r.client.User.
		Query().
		All(context.Background())
	if err != nil {
		return nil, err
	}

	users := make([]User, 0, len(rows))

	for _, u := range rows {
		users = append(users, User{
			ID:        u.ID,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			FullName:  u.FullName,
			Email:     u.Email,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		})
	}

	return users, nil
}

func (r *entRepository) Create(user User) (User, error) {
	row, err := r.client.User.
		Create().
		SetNillableFirstName(user.FirstName).
		SetNillableLastName(user.LastName).
		SetNillableFullName(user.FullName).
		SetEmail(user.Email).
		SetPassword(user.Password).
		Save(context.Background())

	if err != nil {
		return User{}, err
	}

	return User{
		ID:        row.ID,
		FirstName: row.FirstName,
		LastName:  row.LastName,
		FullName:  row.FullName,
		Email:     row.Email,
		Password:  row.Password,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}, nil
}

func (r *entRepository) FindByEmail(email string) (User, error) {
	row, err := r.client.User.
		Query().
		Where(user.EmailEQ(email)).
		Only(context.Background())

	if err != nil {
		return User{}, err
	}

	return User{
		ID:        row.ID,
		FirstName: row.FirstName,
		LastName:  row.LastName,
		FullName:  row.FullName,
		Email:     row.Email,
		Password:  row.Password,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}, nil
}
