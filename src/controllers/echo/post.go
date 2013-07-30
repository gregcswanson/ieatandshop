package echoController

import (
  "net/http"
  "time"
  "src/infrastructure"   
  "encoding/json"
  "src/usecases"
)

func Get(w http.ResponseWriter, r *http.Request) {
	defer infrastructure.TimeTrack(time.Now(), "EchoController.Get")
	
	echoInteractor := CreateEchoInteractor(r)
	
	echo, _ := echoInteractor.FindForUser()
	
	infrastructure.SendData(&echo, w)
}

func Post(w http.ResponseWriter, r *http.Request) {
	defer infrastructure.TimeTrack(time.Now(), "EchoController.Post")
	
	echoInteractor := CreateEchoInteractor(r)
	
	decoder := json.NewDecoder(r.Body)
    var postedEcho usecases.Echo   
    err := decoder.Decode(&postedEcho)
    if err != nil {
    	infrastructure.SendError(err, w)
        return
    }
    
    savedEcho, saveError := echoInteractor.Save(postedEcho)
    if saveError != nil {
    	infrastructure.SendError(saveError, w)
  		return
    }
  	infrastructure.SendData(&savedEcho, w)
}