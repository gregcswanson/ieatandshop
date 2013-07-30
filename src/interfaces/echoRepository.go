package interfaces

import (
	"src/domain"
	"net/http"
	"appengine"
	"appengine/datastore"
)

type EchoRepository BaseRepository

func NewEchoRepository(request *http.Request) *EchoRepository {
	echoRepository := new(EchoRepository)
	echoRepository.request = request
	return echoRepository
}

func (repository *EchoRepository) Store(item domain.Echo) (domain.Echo, error) {
	// upsert operation
	c := appengine.NewContext(repository.request)
	if item.ID != "" {
		// update
		key , err := datastore.DecodeKey(item.ID)
		if err != nil {
			return item, err
		}
		_, err = datastore.Put(c, key, &item)
    	if err != nil {
			return item, err
		}
	} else {
		// new
		key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Echos", nil), &item)
    	if err != nil {
        	return item, err
    	} else {
    		item.ID = key.Encode()
    	}
	}
	return item, nil
	
}

func (repository *EchoRepository) FindForUser(userid string) (domain.Echo, error) {
	var echos []domain.Echo
	var echo domain.Echo
	
	c := appengine.NewContext(repository.request)
    q := datastore.NewQuery("Echos").Filter("UserID =", userid).Limit(1)
	
	keys, err := q.GetAll(c, &echos)
    if err != nil {
    	return domain.Echo{}, err
    } else {
    
    	if len(echos) == 1 {
    		echo.ID = keys[0].Encode()
    		echo.Title = echos[0].Title
    	} else {
    		echo.Title = "- new user -"
    	}
    }
  	echo.UserID = userid
  	return echo, nil
}
