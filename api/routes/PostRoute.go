package routes

import (
	"Ecommer/api/controllers"
	"Ecommer/infrastructures"
)

type PostRoute struct {
	_postController controllers.PostController
	_handler        infrastructures.GinRouter
}

func NewPostRoute(postController controllers.PostController,
	handler infrastructures.GinRouter) PostRoute {
	return PostRoute{
		_postController: postController,
		_handler:        handler,
	}
}

func (p PostRoute) Setup() {
	post := p._handler.Gin.Group("/posts")
	{
		post.POST("/", p._postController.Insert)
		post.GET("/", p._postController.Find)
		post.DELETE("/:id", p._postController.Delete)
	}
}
