package repositories

import (
	"gin-gorm-sqlite/models"
	"gin-gorm-sqlite/models/requests"

	"gorm.io/gorm"
)

type NoteRepository struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) *NoteRepository {
	return &NoteRepository{db}
}

func (n *NoteRepository) FindAll() (*[]models.Note, error) {
	notes := []models.Note{}

	result := n.db.Find(&notes)

	if result.Error != nil {
		return nil, result.Error
	}

	return &notes, nil
}

func (n *NoteRepository) FindById(id int) (*models.Note, error) {
	note := models.Note{}

	result := n.db.First(&note, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &note, nil
}

func (n *NoteRepository) Create(note *models.Note) error {
	result := n.db.Create(note)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (n *NoteRepository) UpdateById(id int, updateNoteRequest requests.UpdateNoteRequest) (*models.Note, error) {
	note, err := n.FindById(id)
	if err != nil {
		return nil, err
	}

	n.db.Model(note).Updates(models.Note{
		Title: updateNoteRequest.Title,
		Body:  updateNoteRequest.Body,
	})

	return note, nil
}

func (n *NoteRepository) DeleteById(id int) error {
	note, err := n.FindById(id)
	if err != nil {
		return err
	}

	n.db.Delete(note)

	return nil
}
