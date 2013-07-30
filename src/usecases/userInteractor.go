package usecases

import (
	"src/domain"
)

type UserInteractor struct {
	UserRepository  domain.UserRepository
}

func (interactor *UserInteractor) IsLoggedIn() bool {
  user := interactor.UserRepository.FindCurrent()
  return user.IsLoggedIn
}

func (interactor *UserInteractor) LoginUrl() (string, error) {
	url, err := interactor.UserRepository.LoginUrl()
	return url, err
}

func (interactor *UserInteractor) LogoutUrl() (string, error) {
	url, err := interactor.UserRepository.LogoutUrl()
	return url, err
}