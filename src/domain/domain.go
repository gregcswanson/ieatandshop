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

type EchoRepository interface {
	Store(item Echo) (Echo, error)
	FindForUser(userid string) (Echo, error)
}

type EchoLineRepository interface {
	Store(item EchoLine) (EchoLine, error)
	Delete(id string) error
	FindByEchoID(echoId string) ([]EchoLine, error)
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

type Echo struct {
	ID string `datastore:"-"`
	UserID string
	Title string
}

type EchoLine struct {
	ID string `datastore:"-"`
	EchoID string
	Name string
}