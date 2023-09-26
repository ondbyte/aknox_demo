package notes

import (
	"fmt"
)

type IRepo interface {
	getAll(userEmail string) ([]*Note, error)
	add(userEmail string, note *Note) (*Note, error)
	remove(userEmail string, noteId string) (*Note, error)
}

type repo struct {
	notes map[string]map[string]*Note
}

func NewRepo() IRepo {
	return &repo{notes: make(map[string]map[string]*Note)}
}

// add implements IRepo.
func (r *repo) add(userEmail string, note *Note) (*Note, error) {
	existingNotes, ok := r.notes[userEmail]
	if !ok {
		existingNotes = map[string]*Note{}
	}
	existingNotes[note.Id] = note
	r.notes[userEmail] = existingNotes
	return note, nil
}

// getAll implements IRepo.
func (r *repo) getAll(userEmail string) ([]*Note, error) {
	all := []*Note{}
	for _, note := range r.notes[userEmail] {
		if note != nil {
			all = append(all, note)
		}
	}
	return all, nil
}

// remove implements IRepo.
func (r *repo) remove(userEmail string, noteId string) (*Note, error) {
	existingNotes, ok := r.notes[userEmail]
	if !ok {
		return nil, fmt.Errorf("user has no notes")
	}
	existingNote, ok := existingNotes[noteId]
	if !ok {
		return nil, fmt.Errorf("user has no note with id %v", noteId)
	}
	existingNotes[noteId] = nil
	return existingNote, nil
}
