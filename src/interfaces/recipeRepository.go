package interfaces

import (
	"src/domain"
	"net/http"
	"appengine"
	"appengine/datastore"
)

type RecipeRepository BaseRepository

func NewRecipeRepository(request *http.Request) *RecipeRepository {
	recipeRepository := new(RecipeRepository)
	recipeRepository.request = request
	return recipeRepository
}

func (repository *RecipeRepository) Store(item domain.Recipe) (domain.Recipe, error) {
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
		key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Recipes", nil), &item)
    	if err != nil {
        	return item, err
    	} else {
    		item.ID = key.Encode()
    	}
	}
	return item, nil
}

func (repository *RecipeRepository) FindForUser(userid string) ([]domain.Recipe, error) {
	var recipes []domain.Recipe
	
	c := appengine.NewContext(repository.request)
  q := datastore.NewQuery("Recipes").Filter("UserID =", userid).Limit(1)
	
	keys, err := q.GetAll(c, &recipes)
  if err != nil {
    return recipes, err
  } else {
    // loop through and add the keys as ID
    for i := 0; i < len(keys); i++ {
      recipes[i].ID = keys[i].Encode()
    }
  }
  return recipes, nil
}
