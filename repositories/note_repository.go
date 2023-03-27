package repositories

import (
	"gin-gorm-sqlite/models"

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
