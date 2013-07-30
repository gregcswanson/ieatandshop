package echoController

import (
  "net/http"
  "time"
  "src/usecases"
  "src/interfaces"
  "src/infrastructure"  
  "src/infrastructure/views"
)

func CreateEchoInteractor(r *http.Request) *usecases.EchoInteractor {
	echoInteractor := new(usecases.EchoInteractor)
	echoInteractor.EchoRepository = interfaces.NewEchoRepository(r)
	echoInteractor.EchoLineRepository = interfaces.NewEchoLineRepository(r)
	echoInteractor.UserRepository = interfaces.NewUserRepositiory(r)
  	return echoInteractor
}

func Index(w http.ResponseWriter, r *http.Request) {
	defer infrastructure.TimeTrack(time.Now(), "EchoController.Index")
	_, v := views.StaticLayout("src/controllers/echo/templates/index")
    w.Write(v)
}