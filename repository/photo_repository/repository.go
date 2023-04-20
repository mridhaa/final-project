package photo_repository

import (
	"h8-movies/entity"
	"h8-movies/pkg/errs"
)

type PhotoRepository interface {
	CreatePhoto(photoPayload *entity.Photo) (*entity.Photo, errs.MessageErr)
	GetPhotoById(photoId int) (*entity.Photo, errs.MessageErr)
	GetPhotoAll(userId int) ([]entity.Photo, errs.MessageErr)
	DeletePhotoById(photoId int) errs.MessageErr
	UpdatePhotoById(payload entity.Photo) errs.MessageErr
}
