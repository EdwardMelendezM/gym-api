package impl

import (
	"context"
)

func (r *entRepository) GetTotalUsers(ctx context.Context) (int, error) {
	query := r.client.User.Query()

	total, err := query.Count(ctx)
	if err != nil {
		return 0, err
	}
	return total, nil
}
