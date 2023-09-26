package auth

import (
	"fmt"
)

type IService interface {
	SignUp(name, email, password string) (*User, error)
	Login(email, password string) (*User, error)
}

func NewService(repo IRepo) IService {
	return &service{repo: repo}
}

type service struct {
	repo IRepo
}

// Login implements IService.
func (s *service) Login(email string, password string) (*User, error) {
	existingUser := s.repo.getUser(email)
	if existingUser == nil {
		return nil, fmt.Errorf("user doesnt exist")
	}
	if existingUser.password != password {
		return nil, fmt.Errorf("wrong password for user email %v", email)
	}
	return existingUser, nil
}

// SignUp implements IService.
func (s *service) SignUp(name string, email string, password string) (*User, error) {
	existingUser := s.repo.getUser(email)
	if existingUser != nil {
		return nil, fmt.Errorf("user already registered for email %v, please login", email)
	}
	newUser, err := s.repo.createUser(name, email, password)
	if err != nil {
		return nil, fmt.Errorf("error signing up due to err: %v", err)
	}
	return newUser, nil
}
