CREATE TABLE IF NOT EXISTS "users" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar NOT NULL UNIQUE,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "accounts" (
  "id" bigserial PRIMARY KEY,
  "type" varchar NOT NULL,
  "currency" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "clients" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar NOT NULL UNIQUE,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "accounts" ("user_id");

CREATE INDEX ON "clients" ("email");

COMMENT ON COLUMN "accounts"."balance" IS 'must be positive';
