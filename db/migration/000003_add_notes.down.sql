ALTER TABLE IF EXISTS "notes" DROP CONSTRAINT IF EXISTS "notes_author_fkey";
ALTER TABLE IF EXISTS "notes" DROP CONSTRAINT IF EXISTS "notes_post_id_fkey";
DROP TABLE IF EXISTS notes;
ALTER TABLE "posts" DROP COLUMN "book_author";
ALTER TABLE "posts" DROP COLUMN "book_image";
ALTER TABLE "posts" DROP COLUMN "book_page";
ALTER TABLE "posts" DROP COLUMN "book_page_read";
ALTER TABLE "posts"
ADD COLUMN "body" text NOT NULL