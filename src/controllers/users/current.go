package usersController

import (
  "fmt"
  "net/http"
  "time"
  "src/usecases"
  "src/interfaces"
  "src/infrastructure"  
)

func Current(w http.ResponseWriter, r *http.Request) {
  defer infrastructure.TimeTrack(time.Now(), "usersController.Current")
   // dependancy injection
  commentInteractor := new(usecases.CommentInteractor)
	commentInteractor.UserRepository = interfaces.NewUserRepositiory(r)
	commentInteractor.CommentRepository = interfaces.NewCommentRepositiory(r)
  // get the data
  userName, _ := commentInteractor.CurrentUserName()
  message := fmt.Sprintf("Hello, %s!",userName)
  infrastructure.SendData(message, w)
}