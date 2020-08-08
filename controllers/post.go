package controllers

import (
	"net/http"

	"github.com/angelmendozacap/mvc/database"
	"github.com/angelmendozacap/mvc/utils"
)

// GetPosts retrieve all the posts
func GetPosts(w http.ResponseWriter, r *http.Request) {
	posts := database.Posts
	utils.ExecuteTemplate(w, "index.html", posts)
}
