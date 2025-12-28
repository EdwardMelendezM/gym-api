package users

import (
	"context"

	"gym-api/internal/ent"
)

type entRepository struct {
	client *ent.Client
}

func NewEntRepository(client *ent.Client) Repository {
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
			Name:      u.Name,
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
		SetName(user.Name).
		SetEmail(user.Email).
		Save(context.Background())

	if err != nil {
		return User{}, err
	}

	return User{
		ID:        row.ID,
		Name:      row.Name,
		Email:     row.Email,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}, nil
}
