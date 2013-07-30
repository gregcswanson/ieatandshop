package interfaces

import (
	"src/domain"
  "net/http"
  "appengine"
  "appengine/user"
  	//"strconv"
)

type BaseRepository struct {
	request  *http.Request
}

type UserRepositiory BaseRepository
type CommentRepositiory BaseRepository

func NewUserRepositiory(request *http.Request) *UserRepositiory {
	userRepository := new(UserRepositiory)
	userRepository.request = request
	return userRepository
}

func NewCommentRepositiory(request *http.Request) *CommentRepositiory {
	commentRepositiory := new(CommentRepositiory)
	commentRepositiory.request = request
	return commentRepositiory
}

func (repository *UserRepositiory) FindCurrent() domain.User {
  // is the current user logged in
  c := appengine.NewContext(repository.request)
  u := user.Current(c)
  var user domain.User
  user = domain.User{}
  if u == nil {
    user.Id = "0"
    user.Name = ""
    user.Nickname = ""
    user.Email = ""
    user.IsLoggedIn = false
  } else {
    user.Id = u.ID
    user.Name = ""
    user.Nickname = ""
    user.Email = u.Email
    user.IsLoggedIn = true
  }  
  return user
}

func (repository *UserRepositiory) Store(user domain.User) domain.User {
  var u = domain.User{}
  return u
}

func (repository *UserRepositiory) LoginUrl()  (string, error){
  c := appengine.NewContext(repository.request)
  u := user.Current(c)
    if u == nil {
        url, err := user.LoginURL(c, repository.request.URL.String())
        return url, err
    }
  return "/", nil 
}

func (repository *UserRepositiory) LogoutUrl()  (string, error){
  c := appengine.NewContext(repository.request)
  u := user.Current(c)
  if u != nil {
    url, err := user.LogoutURL(c, repository.request.URL.String())
    return url, err
  }
  return "/", nil 
}


func (repository *CommentRepositiory) Store(comment domain.Comment) {
	// save the comment using the storage context
}

func (repository *CommentRepositiory) FindForUserId(userid string) []domain.Comment {
  var comments []domain.Comment
	comments = make([]domain.Comment, 5)
	//for i, _ := range comments {
	//is := strconv.itoa(i)
    //comments[i] = domain.Comment{is, userid, "Comment"}
	//}
	return comments
}