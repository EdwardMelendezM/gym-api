package users

import (
	"context"
	"gym-api/internal/ent"
	"gym-api/internal/ent/user"
	"net/http"

	"gym-api/internal/errors"
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
		return nil, errors.New().
			SetStatus(http.StatusInternalServerError).
			SetLayer("users.repository").
			SetFunction("GetAll").
			SetMessage("failed to get all users").
			SetError(err)
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
		return User{}, errors.New().
			SetStatus(http.StatusInternalServerError).
			SetLayer("users.repository").
			SetFunction("Create").
			SetMessage("failed to insert user").
			SetError(err)
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
		return User{}, errors.New().
			SetStatus(http.StatusInternalServerError).
			SetLayer("users.repository").
			SetFunction("FindByEmail").
			SetMessage("failed to find by email").
			SetError(err)
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

func (r *entRepository) FindById(id string) (User, error) {
	row, err := r.client.User.
		Query().
		Where(user.IDEQ(id)).
		Only(context.Background())

	if err != nil {
		return User{}, errors.New().
			SetStatus(http.StatusInternalServerError).
			SetLayer("users.repository").
			SetFunction("FindById").
			SetMessage("failed to find by email").
			SetError(err)
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
