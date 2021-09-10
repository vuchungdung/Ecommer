package models

import "time"

type Post struct {
	Id        int64     `gorm:"primary_key; auto_increment" json:"id"`
	Title     string    `gorm:"size:200" json:"title"`
	Body      string    `gorm:"size:1000" json:"body"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (post *Post) TableName() string {
	return "post"
}

func (post *Post) ResponseMap() map[string]interface{} {
	response := make(map[string]interface{})
	response["id"] = post.Id
	response["title"] = post.Title
	response["body"] = post.Body
	response["created_at"] = post.CreatedAt
	response["updated_at"] = post.UpdatedAt
	return response
}
