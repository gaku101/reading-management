CREATE TABLE "notes" (
  "id" bigserial PRIMARY KEY,
  "author" varchar NOT NULL,
  "post_id" bigint NOT NULL,
  "body" text NOT NULL,
  "page" smallint NOT NULL,
  "line" smallint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
ALTER TABLE "notes"
ADD FOREIGN KEY ("author") REFERENCES "users" ("username");
ALTER TABLE "notes"
ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");
ALTER TABLE "posts"
ADD COLUMN "book_author" varchar NOT NULL;
ALTER TABLE "posts"
ADD COLUMN "book_image" varchar NOT NULL;
ALTER TABLE "posts"
ADD COLUMN "book_page" smallint NOT NULL;
ALTER TABLE "posts"
ADD COLUMN "book_page_read" smallint NOT NULL;
ALTER TABLE "posts" DROP COLUMN "body";