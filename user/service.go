package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	CheckEmail(input CheckEmailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Name = input.Name
	user.Occupation = input.Occupation
	user.Email = input.Email
	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	rsUser, err := s.repository.FindByEmail(email)
	if err != nil {
		return rsUser, err
	}
	if rsUser.ID == 0 {
		return rsUser, errors.New("Email or password is wrong")
	}

	errCompare := bcrypt.CompareHashAndPassword([]byte(rsUser.PasswordHash), []byte(password))
	if errCompare != nil {
		return rsUser, errors.New("Email or password is wrong")
	}

	return rsUser, nil
}

func (s *service) CheckEmail(input CheckEmailInput) error {
	email := input.Email

	rsUser, err := s.repository.FindByEmail(email)
	if err != nil {
		return err
	}

	if rsUser.ID != 0 {
		return errors.New("Email is already taken")
	}

	return nil
}
