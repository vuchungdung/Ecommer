package interfaces

import "Ecommer/models"

type IPostService interface {
	Insert(post models.Post) error
	Find(keyword string) (*[]models.Post, error)
	Delete(id int64) error
}
