-- Modify "users" table
ALTER TABLE "public"."users" ADD COLUMN "full_name" character varying NOT NULL, ADD COLUMN "first_name" character varying NOT NULL, ADD COLUMN "last_name" character varying NOT NULL;
