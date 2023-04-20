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

type socialmediaHandler struct {
	socialmediaService service.SocialMediaService
}

func NewSocialMediaHandler(socialmediaService service.SocialMediaService) socialmediaHandler {
	return socialmediaHandler{
		socialmediaService: socialmediaService,
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
func (m socialmediaHandler) CreateSocialMedia(c *gin.Context) {
	var socialmediaRequest dto.NewSocialMediasRequest

	if err := c.ShouldBindJSON(&socialmediaRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	user := c.MustGet("userData").(entity.User)

	newSocialMedia, err := m.socialmediaService.CreateSocialMedia(user.Id, socialmediaRequest)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, newSocialMedia)
}

func (m socialmediaHandler) UpdateSocialMediaById(c *gin.Context) {
	var socialmediaRequest dto.NewSocialMediasRequest

	if err := c.ShouldBindJSON(&socialmediaRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	socialmediaId, err := helpers.GetParamId(c, "socialmediaId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := m.socialmediaService.UpdateSocialMediaById(socialmediaId, socialmediaRequest)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}



func (m socialmediaHandler) GetSocialMediaById(c *gin.Context) {
	
	socialmediaId, err := helpers.GetParamId(c, "socialmediaId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := m.socialmediaService.GetSocialMediaById(socialmediaId)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}


func (m socialmediaHandler) GetSocialMediaAll(c *gin.Context) {

	user := c.MustGet("userData").(entity.User)

	response, err := m.socialmediaService.GetSocialMediaAll(user.Id)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}


func (m socialmediaHandler) DeleteSocialMediaById(c *gin.Context) {
	
	socialmediaId, err := helpers.GetParamId(c, "socialmediaId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := m.socialmediaService.DeleteSocialMediaById(socialmediaId)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}

