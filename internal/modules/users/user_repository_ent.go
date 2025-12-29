package users

import (
	"context"
	"gym-api/internal/ent"
	"gym-api/internal/ent/user"
	"gym-api/internal/modules/shared/pagination"
	"net/http"

	"gym-api/internal/errors"
)

var allowedOrderFields = map[string]string{
	"createdAt": user.FieldCreatedAt,
	"email":     user.FieldEmail,
	"firstName": user.FieldFirstName,
}

type entRepository struct {
	client *ent.Client
}

func NewUserEntRepository(client *ent.Client) Repository {
	return &entRepository{client: client}
}

func (r *entRepository) GetAll(p pagination.Params) ([]*ent.User, int, error) {
	ctx := context.Background()

	query := r.client.User.Query()

	// 🔍 SEARCH
	if p.Search != "" && p.SearchBy != "" {
		switch p.SearchBy {
		case "email":
			query = query.Where(user.EmailContainsFold(p.Search))
		case "firstName":
			query = query.Where(user.FirstNameContainsFold(p.Search))
		case "lastName":
			query = query.Where(user.LastNameContainsFold(p.Search))
		}
	}

	// 🔃 ORDER
	field, ok := allowedOrderFields[p.OrderBy]
	if !ok {
		field = user.FieldCreatedAt
	}

	if p.Order == "asc" {
		query = query.Order(ent.Asc(field))
	} else {
		query = query.Order(ent.Desc(field))
	}

	// total
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	// pagination
	users, err := query.
		Limit(p.PageSize).
		Offset(p.Offset()).
		All(ctx)

	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
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
