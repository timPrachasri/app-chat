DROP TABLE IF EXISTS "public"."users" CASCADE;

CREATE TABLE "public"."users" (
    "id" serial NOT NULL,
    "first_name" character varying(64) NOT NULL,
    "last_name" character varying(64) NOT NULL,
    "date_of_birth" date,
    CONSTRAINT "user_id" PRIMARY KEY ("id")
);
