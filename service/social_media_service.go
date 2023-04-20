package service

import (
	"fmt"
	"h8-movies/dto"
	"h8-movies/entity"
	"h8-movies/pkg/errs"
	"h8-movies/pkg/helpers"
	"h8-movies/repository/social_media_repository"
	"net/http"
)

type SocialMediaService interface {
	CreateSocialMedia(userId int, payload dto.NewSocialMediasRequest) (*dto.NewSocialMediasResponse, errs.MessageErr)
	UpdateSocialMediaById(socialmediId int, socialmediaRequest dto.NewSocialMediasRequest) (*dto.NewSocialMediasResponse, errs.MessageErr)
	GetSocialMediaById(socialmediId int) (*dto.NewSocialMediasResponse, errs.MessageErr)
	DeleteSocialMediaById(socialmediId int) (*dto.NewSocialMediasResponse, errs.MessageErr)
	GetSocialMediaAll(userId int) (*dto.NewSocialMediasResponse, errs.MessageErr)
}

type socialmediaService struct {
	socialmediaRepo social_media_repository.SocialMediaRepository
}

func NewSocialMediaService(socialmediaRepo social_media_repository.SocialMediaRepository) SocialMediaService {
	return &socialmediaService{
		socialmediaRepo: socialmediaRepo,
	}
}

func (m *socialmediaService) UpdateSocialMediaById(socialmediId int, socialmediaRequest dto.NewSocialMediasRequest) (*dto.NewSocialMediasResponse, errs.MessageErr) {

	err := helpers.ValidateStruct(socialmediaRequest)

	if err != nil {
		return nil, err
	}

	payload := entity.SocialMedias{
		Id:       socialmediId,
		Name:    socialmediaRequest.Name,
		SocialMediaUrl: socialmediaRequest.SocialMediaUrl,
	}

	err = m.socialmediaRepo.UpdateSocialMediaById(payload)

	if err != nil {
		return nil, err
	}

	response := dto.NewSocialMediasResponse{
		StatusCode: http.StatusOK,
		Result:     "success",
		Message:    "social media data successfully updated",
	}

	return &response, nil
}

func(m *socialmediaService) GetSocialMediaAll(userId int)(*dto.NewSocialMediasResponse, errs.MessageErr){
	result, err := m.socialmediaRepo.GetSocialMediaAll(userId)

	if err != nil {
		return nil, err
	}

	response := dto.NewSocialMediasResponse{
		StatusCode: http.StatusOK,
		Result:     fmt.Sprintf("%#v", result),
		Message:    "Get SocialMedias All",
	}

	return &response, nil
}


func (m *socialmediaService) GetSocialMediaById(socialmediId int) (*dto.NewSocialMediasResponse, errs.MessageErr) {

	
	result, err := m.socialmediaRepo.GetSocialMediaById(socialmediId)

	if err != nil {
		return nil, err
	}

	response := dto.NewSocialMediasResponse{
		StatusCode: http.StatusOK,
		Result:     fmt.Sprintf("%+v", result),
		Message:    "Get SocialMedias By Id Successfully",
	}

	return &response, nil
}


func (m *socialmediaService) DeleteSocialMediaById(socialmediId int) (*dto.NewSocialMediasResponse, errs.MessageErr) {

	
	err := m.socialmediaRepo.DeleteSocialMediaById(socialmediId)

	if err != nil {
		return nil, err
	}

	response := dto.NewSocialMediasResponse{
		StatusCode: http.StatusOK,
		Result:     "success",
		Message:    "social media data success deleted",
	}

	return &response, nil
}

func (m *socialmediaService) CreateSocialMedia(userId int, payload dto.NewSocialMediasRequest) (*dto.NewSocialMediasResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	
	socialmediaRequest := &entity.SocialMedias{
		Name:    payload.Name,
		SocialMediaUrl: payload.SocialMediaUrl,
		UserId:   userId,
	}

	_, err = m.socialmediaRepo.CreateSocialMedia(socialmediaRequest)

	if err != nil {
		return nil, err
	}

	response := dto.NewSocialMediasResponse{
		StatusCode: http.StatusCreated,
		Result:     "success",
		Message:    "new social media data successfully created",
	}

	return &response, err
}
