package impl

import (
	"context"
)

func (r *sessionEntRepository) Delete(id string) error {
	return r.client.Session.
		DeleteOneID(id).
		Exec(context.Background())
}
