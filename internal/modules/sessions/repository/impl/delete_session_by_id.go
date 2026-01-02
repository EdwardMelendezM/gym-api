package impl

import (
	"context"
)

func (r *sessionEntRepository) DeleteSessionById(id string) error {
	return r.client.Session.
		DeleteOneID(id).
		Exec(context.Background())
}
