package handler

import (
	"h8-movies/dto"
	"h8-movies/entity"
	"h8-movies/pkg/errs"
	"h8-movies/pkg/helpers"
	"h8-movies/service"
	"net/http"

	_ "h8-movies/entity"

	"github.com/gin-gonic/gin"
)

type commentHandler struct {
	commentService service.CommentService
}

func NewCommentHandler(commentService service.CommentService) commentHandler {
	return commentHandler{
		commentService: commentService,
	}
}

// CreateNewMovie godoc
// @Tags movies
// @Description Create New Movie Data
// @ID create-new-movie
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.NewMovieRequest true "request body json"
// @Success 201 {object} dto.NewMovieRequest
// @Router /movies [post]
func (m commentHandler) CreateComment(c *gin.Context) {
	var commentRequest dto.NewCommentRequest

	if err := c.ShouldBindJSON(&commentRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	photoId, err := helpers.GetParamId(c, "photoId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	user := c.MustGet("userData").(entity.User)

	newPhoto, err := m.commentService.CreateComment(user.Id, photoId, commentRequest)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, newPhoto)
}

func (m commentHandler) UpdateCommentById(c *gin.Context) {
	var commentRequest dto.NewCommentRequest

	if err := c.ShouldBindJSON(&commentRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	commentId, err := helpers.GetParamId(c, "commentId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := m.commentService.UpdateCommentById(commentId, commentRequest)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}



func (m commentHandler) GetCommentById(c *gin.Context) {
	
	commentId, err := helpers.GetParamId(c, "commentId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := m.commentService.GetCommentById(commentId)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}


func (m commentHandler) GetCommentAll(c *gin.Context) {

	user := c.MustGet("userData").(entity.User)

	response, err := m.commentService.GetCommentAll(user.Id)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}


func (m commentHandler) DeleteCommentById(c *gin.Context) {
	
	commentId, err := helpers.GetParamId(c, "commentId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := m.commentService.DeleteCommentById(commentId)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}

