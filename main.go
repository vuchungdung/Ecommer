package main

import (
	"Ecommer/api/controllers"
	"Ecommer/api/repositories"
	"Ecommer/api/routes"
	"Ecommer/api/services"
	"Ecommer/infrastructures"
	"Ecommer/models"
)

func init() {
	infrastructures.LoadEnv()
}

func main() {
	router := infrastructures.NewGinRouter()
	db := infrastructures.NewDatabase()
	postRepository := repositories.NewPostRepository(db)
	postSerivce := services.NewPostService(postRepository)
	postController := controllers.NewPostController(postSerivce)
	postRoute := routes.NewPostRoute(postController, router)
	postRoute.Setup()
	db.DB.AutoMigrate(&models.Post{})

	router.Gin.Run(":8000")
}
