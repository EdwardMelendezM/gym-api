package impl

import (
	"context"
	"math"

	"golang.org/x/sync/errgroup"

	"gym-api/internal/ent"
	"gym-api/internal/modules/users/models"
	"gym-api/internal/utils/pagination"
)

func (s *service) GetUsersPaginated(ctx context.Context, p pagination.Params) (*pagination.Result[models.UserResponse], error) {

	// declare variables
	var (
		users []*ent.User
		total int
	)

	// use goroutines
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		var err error
		users, err = s.repo.GetUsersPaginated(ctx, p)
		return err
	})

	g.Go(func() error {
		var err error
		total, err = s.repo.GetTotalUsers(ctx)
		return err
	})

	if err := g.Wait(); err != nil {
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
