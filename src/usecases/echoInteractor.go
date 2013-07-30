package usecases

import (
	"src/domain"
	"errors"
)

type Echo struct {
	ID string
	Title string
	Lines []EchoLine
}

type EchoLine struct {
	ID string
	Name string
}

type EchoInteractor struct {
	EchoRepository domain.EchoRepository
	EchoLineRepository domain.EchoLineRepository
	UserRepository domain.UserRepository
}

func (interactor *EchoInteractor) FindForUser() (Echo, error) {
	user := interactor.UserRepository.FindCurrent()
	domainEcho, err := interactor.EchoRepository.FindForUser(user.Id)
	if err != nil {
		return Echo{}, err
	}
	echo := Echo{domainEcho.ID, domainEcho.Title, nil}
	
	if echo.ID != "" {
		lines, _ := interactor.EchoLineRepository.FindByEchoID(echo.ID)
		echo.Lines = make([]EchoLine, len(lines))
		for i := 0; i < len(lines); i++ {
			echo.Lines[i] = EchoLine{lines[i].ID, lines[i].Name}
		}
	} else {
		echo.Lines = make([]EchoLine, 1)
  	 	echo.Lines[0] = EchoLine{"1", "Add a new line"}
	}
	
	return echo, nil
}

func (interactor *EchoInteractor) Save(echo Echo) (Echo, error) {
	// validate 
	if echo.Title == "" {
		err := errors.New("Title is required")
		return echo, err
	}
	
	
	// find the current user
	user := interactor.UserRepository.FindCurrent()
	
	entity := domain.Echo{}
	entity.ID = echo.ID
	entity.UserID = user.Id
	entity.Title = echo.Title
	// save the header
	entity, err := interactor.EchoRepository.Store(entity)
	if err == nil {
		echo.ID = entity.ID
		
		// get the current list of lines, use a map to indicate the item was updated
		lineIDs := make(map[string]bool)
		currentLines, errCurrentLines:= interactor.EchoLineRepository.FindByEchoID(echo.ID)
		if errCurrentLines != nil {
			return echo, errCurrentLines
		}
		for i := 0; i < len(currentLines); i++ {
			lineIDs[currentLines[i].ID] = false
		}
		
		// save the lines
		for i := 0; i < len(echo.Lines); i++ {
			if echo.Lines[i].ID == "" {
				// add
				entityLine := domain.EchoLine{}
				entityLine.EchoID = echo.ID
				entityLine.Name = echo.Lines[i].Name
				entityLine, errLines := interactor.EchoLineRepository.Store(entityLine)
				if errLines == nil {
					echo.Lines[i].ID = entityLine.ID				}
			} else {
				// edit
				entityLine := domain.EchoLine{}
				entityLine.ID = echo.Lines[i].ID
				entityLine.EchoID = echo.ID
				entityLine.Name = echo.Lines[i].Name
				entityLine, _ = interactor.EchoLineRepository.Store(entityLine)
				lineIDs[echo.Lines[i].ID] = true // mark the row as updated
			}
		}
		
		// delete any that haven't been updated
		for key, value := range lineIDs {
			if value == false {
				err = interactor.EchoLineRepository.Delete(key)
				if err != nil {
					return echo, err
				}
			}
		}
		
	}
	
	return echo, err
}