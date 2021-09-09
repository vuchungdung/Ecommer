package controllers

import (
	services "Ecommer/api/services/interfaces"
	"Ecommer/models"
	"Ecommer/utilities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	_postSerivce services.IPostService
}

func NewPostController(postService services.IPostService) PostController {
	return PostController{
		_postSerivce: postService,
	}
}

func (p PostController) Insert(ctx *gin.Context) {
	var post models.Post
	ctx.ShouldBindJSON(&post)
	if post.Title == "" {
		utilities.Error(ctx, http.StatusBadRequest, post)
		return
	}
	if post.Body == "" {
		utilities.Error(ctx, http.StatusBadRequest, "Body is required")
	}
	err := p._postSerivce.Insert(post)
	if err != nil {
		utilities.Error(ctx, http.StatusBadGateway, "Failed to create post")
		return
	}
	utilities.Success(ctx, http.StatusCreated, "Successfully")
}

func (p PostController) Find(ctx *gin.Context) {
	keyword := ctx.Query("keyword")

	data, err := p._postSerivce.Find(keyword)

	if err != nil {
		utilities.Error(ctx, http.StatusBadRequest, "Faild to find")
		return
	}

	response := make([]map[string]interface{}, 0)

	for _, n := range *data {
		resp := n.ResponseMap()
		response = append(response, resp)
	}

	ctx.JSON(http.StatusOK, &utilities.Response{
		Success: true,
		Message: "Success",
		Data: map[string]interface{}{
			"datas": response,
		},
	})
}

func (p PostController) Delete(ctx *gin.Context) {
	idParam := ctx.Query("id")

	if idParam == "" {
		utilities.Error(ctx, http.StatusBadRequest, "Id not nil")
		return
	}

	id, _ := strconv.ParseInt(idParam, 10, 64)

	err := p._postSerivce.Delete(id)

	if err != nil {
		utilities.Error(ctx, http.StatusBadRequest, "Faild to delete")
		return
	}

	response := &utilities.Response{
		Success: true,
		Message: "Successfully",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, response)
}
