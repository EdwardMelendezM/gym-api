package impl

import (
	"context"
	"math"

	"gym-api/internal/modules/users/models"
	"gym-api/internal/utils/pagination"
)

func (s *service) ListUsers(ctx context.Context, p pagination.Params) (*pagination.Result[models.UserResponse], error) {
	users, total, err := s.repo.GetAll(ctx, p)
	if err != nil {
		return nil, err
	}

	out := make([]models.UserResponse, 0, len(users))
	for _, u := range users {
		out = append(out, models.UserResponse{
			ID:        u.ID,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
		})
	}

	totalPages := int(math.Ceil(float64(total) / float64(p.PageSize)))

	return &pagination.Result[models.UserResponse]{
		Data: out,
		Meta: pagination.Meta{
			Page:       p.Page,
			PageSize:   p.PageSize,
			Total:      total,
			TotalPages: totalPages,
		},
	}, nil
}
