package impl

import (
	"context"
	"gym-api/internal/ent"
	"gym-api/internal/ent/user"
	"gym-api/internal/utils/pagination"
)

func (r *entRepository) GetUsersPaginated(ctx context.Context, p pagination.Params) ([]*ent.User, error) {
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
	// pagination
	users, err := query.
		Limit(p.PageSize).
		Offset(p.Offset()).
		All(ctx)

	if err != nil {
		return nil, err
	}

	return users, nil
}
