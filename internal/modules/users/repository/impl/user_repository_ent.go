package impl

import (
	"gym-api/internal/ent"
	"gym-api/internal/ent/user"
	"gym-api/internal/modules/users/repository"
)

var allowedOrderFields = map[string]string{
	"createdAt": user.FieldCreatedAt,
	"email":     user.FieldEmail,
	"firstName": user.FieldFirstName,
}

type entRepository struct {
	client *ent.Client
}

func NewUserEntRepository(client *ent.Client) repository.Repository {
	return &entRepository{client: client}
}
