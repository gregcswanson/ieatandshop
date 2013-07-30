package interfaces

import (
	"src/domain"
	"net/http"
	"appengine"
	"appengine/datastore"
)

type EchoLineRepository BaseRepository

func NewEchoLineRepository(request *http.Request) *EchoLineRepository {
	echoLineRepository := new(EchoLineRepository)
	echoLineRepository.request = request
	return echoLineRepository
}

func (repository *EchoLineRepository) Store(item domain.EchoLine) (domain.EchoLine, error) {
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
		key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "EchoLines", nil), &item)
    	if err != nil {
        	return item, err
    	} else {
    		item.ID = key.Encode()
    	}
	}
	return item, nil
	
}

func (repository *EchoLineRepository) FindByEchoID(echoId string) ([]domain.EchoLine, error) {
	var echoLines []domain.EchoLine
	
	c := appengine.NewContext(repository.request)
    q := datastore.NewQuery("EchoLines").Filter("EchoID =", echoId)
	
	keys, err := q.GetAll(c, &echoLines)
    if err != nil {
    	return echoLines, err
    } else {
    	// loop through and add the keys as ID
    	for i := 0; i < len(keys); i++ {
    		echoLines[i].ID = keys[i].Encode()
    	}
    }
  	return echoLines, nil
}

func (repository *EchoLineRepository) Delete(id string) error {
	c := appengine.NewContext(repository.request)
	key , err := datastore.DecodeKey(id)
	if err != nil {
		return err
	}
	err = datastore.Delete(c, key)
    return err
}
