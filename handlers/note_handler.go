package handlers

import (
	"gin-gorm-sqlite/models"
	"gin-gorm-sqlite/models/requests"
	"gin-gorm-sqlite/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NoteHandler struct {
	noteRepository *repositories.NoteRepository
}

func NewNoteHandler(noteRepository *repositories.NoteRepository) *NoteHandler {
	return &NoteHandler{noteRepository}
}

func (n *NoteHandler) GetAll(c *gin.Context) {
	notes, _ := n.noteRepository.FindAll()

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    notes,
	})
}

func (n *NoteHandler) GetById(c *gin.Context) {
	idString := c.Param("id")

	id, _ := strconv.Atoi(idString)

	note, err := n.noteRepository.FindById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to get note",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    note,
	})
}

func (n *NoteHandler) Create(c *gin.Context) {
	var createNoteRequest requests.CreateNoteRequest

	c.Bind(&createNoteRequest)

	note := models.Note{
		Title: createNoteRequest.Title,
		Body:  createNoteRequest.Body,
	}

	err := n.noteRepository.Create(&note)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to create note",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "ok",
		"data":    note,
	})
}

func (n *NoteHandler) UpdateById(c *gin.Context) {
	var updateNoteRequest requests.UpdateNoteRequest

	c.Bind(&updateNoteRequest)

	idString := c.Param("id")

	id, _ := strconv.Atoi(idString)

	note, err := n.noteRepository.UpdateById(id, updateNoteRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to update note",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "note updated",
		"data":    note,
	})
}

func (n *NoteHandler) DeleteById(c *gin.Context) {
	idString := c.Param("id")

	id, _ := strconv.Atoi(idString)

	err := n.noteRepository.DeleteById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to delete note",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    nil,
	})
}
