CREATE TABLE "users" (
  "userid" varchar PRIMARY KEY NOT NULL,
  "username" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
