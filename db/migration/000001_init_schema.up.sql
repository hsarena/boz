CREATE TYPE "wool" AS ENUM (
  'carpet',
  'medium',
  'long',
  'fur',
  'fine'
);

CREATE TABLE "sheeps" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "breed" varchar NOT NULL,
  "wool" wool NOT NULL,
  "color" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);