package auto

import "github.com/playground/api/models"

var users = []models.User{
	models.User{Nickname: "Aaa", Email: "aaa@gmail.com", Password: "aaa1234"},
	models.User{Nickname: "Bbb", Email: "bbb@gmail.com", Password: "bbb1234"},
	models.User{Nickname: "Ccc", Email: "ccc@gmail.com", Password: "ccc1234"},
	models.User{Nickname: "Ddd", Email: "ddd@gmail.com", Password: "ddd1234"},
	models.User{Nickname: "Eee", Email: "eee@gmail.com", Password: "eee1234"},
}

var posts = []models.Post{
	models.Post{Title: "Title", Content: "Hello World"},
	models.Post{Title: "Blackhole", Content: "The whole galaxys are have the blackhole in the galaxy"},
	models.Post{Title: "Quantum physical", Content: "Quantum physical is new way of science"},
	models.Post{Title: "Mathermatical", Content: "Mathermatic is basic of everything"},
	models.Post{Title: "Computer programing", Content: "kids are have to learn programing language"},
}
