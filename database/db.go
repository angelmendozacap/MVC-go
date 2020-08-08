package database

import "github.com/angelmendozacap/mvc/models"

// Posts are the posts store
var Posts = map[uint]models.Post{
	1: {
		ID:      1,
		Title:   "Title 1",
		Content: "lorem ipsum",
	},
	2: {
		ID:      2,
		Title:   "Title 2",
		Content: "lorem ipsum",
	},
}
