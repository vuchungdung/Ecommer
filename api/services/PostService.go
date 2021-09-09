package services

import (
	repositories "Ecommer/api/repositories/interfaces"
	services "Ecommer/api/services/interfaces"
	"Ecommer/models"
)

type PostService struct {
	_postRepository repositories.IPostRepository
}

func NewPostService(postRepository repositories.IPostRepository) services.IPostService {
	return PostService{
		_postRepository: postRepository,
	}
}

func (p PostService) Insert(post models.Post) error {
	return p._postRepository.Insert(post)
}

func (p PostService) Find(keyword string) (*[]models.Post, error) {
	return p._postRepository.Find(keyword)
}

func (p PostService) Delete(id int64) error {
	return p._postRepository.Delete(id)
}
