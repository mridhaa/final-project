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

type photoHandler struct {
	photoService service.PhotoService
}

func NewPhotoHandler(photoService service.PhotoService) photoHandler {
	return photoHandler{
		photoService: photoService,
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
func (m photoHandler) CreatePhoto(c *gin.Context) {
	var photoRequest dto.NewPhotoRequest

	if err := c.ShouldBindJSON(&photoRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	user := c.MustGet("userData").(entity.User)

	newPhoto, err := m.photoService.CreatePhoto(user.Id, photoRequest)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, newPhoto)
}

func (m photoHandler) UpdatePhotoById(c *gin.Context) {
	var photoRequest dto.NewPhotoRequest

	if err := c.ShouldBindJSON(&photoRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	photoId, err := helpers.GetParamId(c, "photoId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := m.photoService.UpdatePhotoById(photoId, photoRequest)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}



func (m photoHandler) GetPhotoById(c *gin.Context) {
	
	photoId, err := helpers.GetParamId(c, "photoId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := m.photoService.GetPhotoById(photoId)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}


func (m photoHandler) GetPhotoAll(c *gin.Context) {

	user := c.MustGet("userData").(entity.User)

	response, err := m.photoService.GetPhotoAll(user.Id)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}


func (m photoHandler) DeletePhotoById(c *gin.Context) {
	
	photoId, err := helpers.GetParamId(c, "photoId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := m.photoService.DeletePhotoById(photoId)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}

