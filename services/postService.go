package services

import (
	"net/http"

	"github.com/Sahil2k07/Golang-Unit-Tests/errors"
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
	email := ps.ctx.QueryParam("email")

	posts, err := ps.repo.GetUserPosts(email)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}

	return ps.ctx.JSON(http.StatusOK, posts)
}
