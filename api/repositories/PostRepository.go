package repositories

import (
	repositories "Ecommer/api/repositories/interfaces"
	"Ecommer/infrastructures"
	"Ecommer/models"
)

type PostRepository struct {
	DB infrastructures.Database
}

func NewPostRepository(db infrastructures.Database) repositories.IPostRepository {
	return PostRepository{
		DB: db,
	}
}

func (p PostRepository) Insert(post models.Post) error {
	return p.DB.DB.Create(&post).Error
}

func (p PostRepository) Find(keyword string) (*[]models.Post, error) {
	var posts []models.Post

	queryBuilder := p.DB.DB.Order("created_at desc").Model(&models.Post{})

	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuilder = queryBuilder.Where(p.DB.DB.Where("post.body Like ? ", queryKeyword))
	}
	err := queryBuilder.Find(&posts).Error

	return &posts, err
}

func (p PostRepository) Delete(id int64) error {
	var post models.Post
	post.Id = id
	return p.DB.DB.Delete(&post).Error
}

func (p PostRepository) Update(post models.Post) error {
	return p.DB.DB.Save(&post).Error
}
