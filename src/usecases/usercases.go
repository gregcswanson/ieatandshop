package usecases

import (
	"src/domain"
)

type CommentInteractor struct {
	UserRepository  domain.UserRepository
	CommentRepository domain.CommentRepository
}

type UserComment struct {
	Id    string
	Name  string
	Comment string
}

func (interactor *CommentInteractor) CurrentUserName() (string, error) {
	user := interactor.UserRepository.FindCurrent()
	return user.Name, nil
}

func (interactor *CommentInteractor) Comments() ([]UserComment, error) {
	var userComments []UserComment
	user := interactor.UserRepository.FindCurrent()
	comments := interactor.CommentRepository.FindForUserId(user.Id)
	userComments = make([]UserComment, len(comments))
	for i, comment := range comments {
		userComments[i] = UserComment{comment.Id, user.Name, comment.Comment}
	}
	return userComments, nil
}