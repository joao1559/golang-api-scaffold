package usecases

import (
	"encoding/json"

	"github.com/joao1559/golang-api-scaffold/interfaces"
	"github.com/joao1559/golang-api-scaffold/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type noteUseCase struct {
	noteRepo interfaces.NoteRepository
}

//NewNoteUseCase will create new an NoteUseCase object representation of usecase.NoteUseCase interface
func NewNoteUseCase(h interfaces.NoteRepository) interfaces.NoteUseCase {
	return &noteUseCase{
		noteRepo: h,
	}
}

func (n *noteUseCase) Get() ([]*models.Note, error) {
	res, err := n.noteRepo.Get()
	return res, err
}

func (n *noteUseCase) Insert(body []byte) error {
	var noteBody models.Note

	err := json.Unmarshal(body, &noteBody)
	if err != nil {
		return err
	}

	noteBody.ID = primitive.NewObjectID()
	err = n.noteRepo.Insert(noteBody)
	return err
}
