package social_media_repository

import (
	"h8-movies/entity"
	"h8-movies/pkg/errs"
)

type SocialMediaRepository interface {
	CreateSocialMedia(socialmediaPayload *entity.SocialMedias) (*entity.SocialMedias, errs.MessageErr)
	GetSocialMediaById(socialmediaId int) (*entity.SocialMedias, errs.MessageErr)
	GetSocialMediaAll(userId int) ([]entity.SocialMedias, errs.MessageErr)
	DeleteSocialMediaById(socialmediaId int) errs.MessageErr
	UpdateSocialMediaById(payload entity.SocialMedias) errs.MessageErr
}
