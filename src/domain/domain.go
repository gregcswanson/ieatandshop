package domain

/*
Define the interfaces for the repositories [SPLIT TO DIFFERENT FILE]
Using Verbs Store Find

*/

type UserRepository interface {
	FindCurrent() User
	Store(user User) User
  LoginUrl() (string, error)
  LogoutUrl() (string, error)
}

type CommentRepository interface {
	Store(item Comment) 
	FindForUserId(userid string) []Comment
}

type User struct {
	Id   string
	Name string
  	Email string
  	Nickname string
  	IsLoggedIn bool
}

type Comment struct {
	Id string
  UserId string
  Comment string
}
