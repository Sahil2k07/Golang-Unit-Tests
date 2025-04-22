package services

import (
	"net/http"

	"github.com/Sahil2k07/Golang-Unit-Tests/interfaces"
	"github.com/labstack/echo"
)

type postService struct {
	ctx  echo.Context
	repo interfaces.IPostRepository
}

func NewPostService(c echo.Context, r interfaces.IPostRepository) interfaces.IPostService {
	return &postService{
		ctx:  c,
		repo: r,
	}
}

func (ps *postService) GetUserPosts() error {
	return ps.ctx.JSON(http.StatusOK, nil)
}
