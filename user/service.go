package user

import (
	"errors"
	"os"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	CheckEmail(input CheckEmailInput) error
	UpdateAvatar(ID int, fileLocation string) (User, error)
	GetUserById(ID int) (User, error)
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
		return rsUser, errors.New("email or password is wrong")
	}

	errCompare := bcrypt.CompareHashAndPassword([]byte(rsUser.PasswordHash), []byte(password))
	if errCompare != nil {
		return rsUser, errors.New("email or password is wrong")
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
		return errors.New("email is already taken")
	}

	return nil
}

func (s *service) UpdateAvatar(ID int, fileLocation string) (User, error) {
	rsUser, err := s.repository.FindByID(ID)
	if err != nil {
		return rsUser, err
	}
	if rsUser.ID == 0 {
		return rsUser, errors.New("User not found")
	}

	existFile := rsUser.AvatarFileName

	rsUser.AvatarFileName = fileLocation
	rsUpdate, err := s.repository.Update(rsUser)
	if err != nil {
		return User{}, err
	}

	if len(existFile) > 0 {
		os.Remove(existFile)
	}

	return rsUpdate, nil
}

func (s *service) GetUserById(ID int) (User, error) {
	rsUser, err := s.repository.FindByID(ID)
	if err != nil {
		return rsUser, err
	}
	if rsUser.ID == 0 {
		return rsUser, errors.New("User not found")
	}

	return rsUser, nil
}
