package domain

type EchoRepository interface {
	Store(item Echo) (Echo, error)
	FindForUser(userid string) (Echo, error)
}

type EchoLineRepository interface {
	Store(item EchoLine) (EchoLine, error)
	Delete(id string) error
	FindByEchoID(echoId string) ([]EchoLine, error)
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