package impl

import (
	"context"
	"gym-api/internal/ent"
	"gym-api/internal/ent/user"
	"gym-api/internal/modules/shared/pagination"
)

func (r *entRepository) GetAll(ctx context.Context, p pagination.Params) ([]*ent.User, int, error) {
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
