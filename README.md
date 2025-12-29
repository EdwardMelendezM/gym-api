# Migrations and generate types
## Install
```sh
curl -sSf https://atlasgo.sh | sh
```

## First migration
- Create a new migration
```sh
atlas migrate diff init \
  --to "ent://internal/ent/schema" \
  --dev-url "postgres://postgres:postgres@localhost:5666/postgres?sslmode=disable"
```

- Apply init migration
```sh
atlas migrate apply \
  --url "postgres://postgres:postgres@localhost:5666/gym-db-postgres?sslmode=disable"
```

## Next migrations
- Create a new migration
```sh
atlas migrate diff name_of_migration \
  --to "ent://internal/ent/schema" \
  --dev-url "postgres://postgres:postgres@localhost:5666/postgres?sslmode=disable"
```

- Apply migration
```sh
atlas migrate apply \
  --url "postgres://postgres:postgres@localhost:5666/gym-db-postgres?sslmode=disable"
```

- Generate types
```sh
go run entgo.io/ent/cmd/ent generate ./internal/ent/schema
```