package repositories

import (
	"github.com/go-pg/pg"
	"github.com/trilobit/go-chat/src/models"
)

type (
	History interface {
		Add(msg models.Message) error
		Load() ([]models.Message, error)
	}

	historyRepository struct {
		db *pg.DB
	}
)

func NewHistory(db *pg.DB) History {
	return &historyRepository{db: db}
}

func (h *historyRepository) Add(msg models.Message) error {
	if _, err := h.db.Model(&msg).Insert(&msg); err != nil {
		return err
	}

	return nil
}

func (h *historyRepository) Load() ([]models.Message, error) {
	var messages []models.Message
	err := h.db.Model(&messages).Order("id ASC").Limit(200).Select()
	if err != nil {
		return nil, err
	}

	return messages, nil
}
