package middlewares

import "aldesrahim/test-go-dating-apps/models"

type LoggedInUser struct {
	User       *models.User
	IsLoggedIn bool
}

func NewLoggedInUser() *LoggedInUser {
	return &LoggedInUser{User: &models.User{}, IsLoggedIn: false}
}

func (l *LoggedInUser) Reset() {
	l.IsLoggedIn = false
	l.User = &models.User{}
}
