package api

import (
	db "github.com/gaku101/my-portfolio/db/sqlc"
	"github.com/gaku101/my-portfolio/util"
)

func randomPost(author string) db.Post {
	return db.Post{
		ID:           util.RandomInt(1, 100),
		Author:       author,
		Title:        util.RandomString(6),
		BookAuthor:   util.RandomString(6),
		BookImage:    util.RandomString(6),
		BookPage:     int16(util.RandomInt(1, 100)),
		BookPageRead: int16(util.RandomInt(1, 100)),
	}
}
