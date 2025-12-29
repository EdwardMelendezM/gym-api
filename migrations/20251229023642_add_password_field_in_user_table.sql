-- Modify "users" table
ALTER TABLE "public"."users" ALTER COLUMN "full_name" DROP NOT NULL, ALTER COLUMN "first_name" DROP NOT NULL, ALTER COLUMN "last_name" DROP NOT NULL, ADD COLUMN "password" character varying NOT NULL;
