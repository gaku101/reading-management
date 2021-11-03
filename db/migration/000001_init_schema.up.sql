CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "profile" varchar NOT NULL,
  "image" varchar NOT NULL,
  "points" bigint NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "last_logined_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "posts" (
  "id" bigserial PRIMARY KEY,
  "author" varchar NOT NULL,
  "title" varchar NOT NULL,
  "book_author" varchar NOT NULL,
  "book_image" varchar NOT NULL,
  "book_page" smallint NOT NULL,
  "book_page_read" smallint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE TABLE "notes" (
  "id" bigserial PRIMARY KEY,
  "author" varchar NOT NULL,
  "post_id" bigint NOT NULL,
  "body" text NOT NULL,
  "page" smallint NOT NULL,
  "line" smallint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "comments" (
  "id" bigserial PRIMARY KEY,
  "author" varchar NOT NULL,
  "post_id" bigint NOT NULL,
  "body" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "post_category" (
  "id" bigserial PRIMARY KEY,
  "post_id" bigint NOT NULL,
  "category_id" bigint NOT NULL
);

CREATE TABLE "category" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "post_favorites" (
  "id" bigserial PRIMARY KEY,
  "post_id" bigint NOT NULL,
  "user_id" bigint NOT NULL
);

CREATE TABLE "follow" (
  "id" bigserial PRIMARY KEY,
  "following_id" bigint NOT NULL,
  "follower_id" bigint NOT NULL
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "from_user_id" bigint NOT NULL,
  "to_user_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "user_badge" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "badge_id" bigint NOT NULL
);

CREATE TABLE "badge" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL
);

ALTER TABLE "posts" ADD FOREIGN KEY ("author") REFERENCES "users" ("username");

ALTER TABLE "notes" ADD FOREIGN KEY ("author") REFERENCES "users" ("username");

ALTER TABLE "notes" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("author") REFERENCES "users" ("username");

ALTER TABLE "comments" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

ALTER TABLE "post_category" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

ALTER TABLE "post_category" ADD FOREIGN KEY ("category_id") REFERENCES "category" ("id");

ALTER TABLE "post_favorites" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

ALTER TABLE "post_favorites" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "follow" ADD FOREIGN KEY ("following_id") REFERENCES "users" ("id");

ALTER TABLE "follow" ADD FOREIGN KEY ("follower_id") REFERENCES "users" ("id");

ALTER TABLE "entries" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_user_id") REFERENCES "users" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_user_id") REFERENCES "users" ("id");

ALTER TABLE "user_badge" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_badge" ADD FOREIGN KEY ("badge_id") REFERENCES "badge" ("id");

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "posts" ("author");

CREATE INDEX ON "entries" ("user_id");

CREATE INDEX ON "transfers" ("from_user_id");

CREATE INDEX ON "transfers" ("to_user_id");

CREATE INDEX ON "transfers" ("from_user_id", "to_user_id");
