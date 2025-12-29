-- Create "users" table
CREATE TABLE "public"."users" (
  "id" character varying NOT NULL,
  "name" character varying NOT NULL,
  "email" character varying NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  PRIMARY KEY ("id")
);
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX "users_email_key" ON "public"."users" ("email");
