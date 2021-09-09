package interfaces

import (
	"Ecommer/models"
)

type IPostRepository interface {
	Insert(post models.Post) error
	Find(keyword string) (*[]models.Post, error)
	Delete(id int64) error
	Update(post models.Post) error
}
