ALTER TABLE "comments"
ADD COLUMN "author" varchar NOT NULL;
ALTER TABLE "comments"
ADD FOREIGN KEY ("author") REFERENCES "users" ("username");