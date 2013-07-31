package domain

type RecipeRepository interface {
	Store(item Recipe) (Recipe, error)
	FindForUser(userid string) ([]Recipe, error)
}

type RecipeMethodRepository interface {
	Store(item RecipeMethod) (Recipe, error)
	FindByRecipeID(recipeid string) ([]RecipeMethod, error)
}

type Recipe struct {
	ID string `datastore:"-"`
	UserID string
  IsPublic bool
	Title string
  Description string
  Tips string
  Feeds int
  Calories int
  Fat int
  PrepMinutes int
  CookingMinutes int
}

type RecipeMethod struct {
	ID string `datastore:"-"`
  Sort int
	RecipeID string
	Description string
}