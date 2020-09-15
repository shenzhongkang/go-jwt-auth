package auto

import "gome/api/models"

var users = []models.User{
	models.User{
		Nickname: "Jhon Doe",
		Email: "jhondoe@email.com",
		Password: "123456",
	},
}

var posts = []models.Post{
	models.Post{
		Title: "Title",
		Content: "<h1>123</h1>",
	},
}