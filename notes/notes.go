package notes

import (
	"fmt"

	"github.com/ondbyte/aknox_demo/auth"
)

type IService interface {
	getAll(userEmail string) ([]*Note, error)
	create(userEmail string, note string) (*Note, error)
	delete(userEmail string, noteId string) (*Note, error)
}
type service struct {
	repo        IRepo
	authService auth.IService
}

// create implements IService.
func (s *service) create(userEmail string, note string) (*Note, error) {
	if note == "" {
		return nil, fmt.Errorf("note cannot be empty")
	}
	n, err := s.repo.add(userEmail, &Note{Id: newId(), Note: note})
	if err != nil {
		return nil, fmt.Errorf("error while creating new note for user email %v, err: %v", userEmail, err)
	}
	return n, nil
}

// delete implements IService.
func (s *service) delete(userEmail string, noteId string) (*Note, error) {
	n, err := s.repo.remove(userEmail, noteId)
	if err != nil {
		return nil, fmt.Errorf("error while deleting note %v for user email %v, err: %v", noteId, userEmail, err)
	}
	return n, nil
}

// getAll implements IService.
func (s *service) getAll(userEmail string) ([]*Note, error) {
	all, err := s.repo.getAll(userEmail)
	if err != nil {
		return nil, fmt.Errorf("error while getting all notes for user email %v, err: %v", userEmail, err)
	}
	return all, nil
}

func NewService(repo IRepo) IService {
	return &service{repo: repo}
}
