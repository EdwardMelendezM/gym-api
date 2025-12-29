package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/lucsky/cuid"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(func() string {
				return cuid.New()
			}).
			Immutable(),

		field.String("fullName").
			Optional().
			Nillable(),

		field.String("firstName").
			Optional().
			Nillable(),

		field.String("lastName").
			Optional().
			Nillable(),

		field.String("email").
			Unique().
			NotEmpty(),

		field.String("password").
			Sensitive(),

		field.Time("created_at").
			Default(time.Now),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}
