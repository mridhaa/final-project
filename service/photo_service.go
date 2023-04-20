package service

import (
	"fmt"
	"h8-movies/dto"
	"h8-movies/entity"
	"h8-movies/pkg/errs"
	"h8-movies/pkg/helpers"
	"h8-movies/repository/photo_repository"
	"net/http"
)

type PhotoService interface {
	CreatePhoto(userId int, payload dto.NewPhotoRequest) (*dto.NewPhotoResponse, errs.MessageErr)
	UpdatePhotoById(photoId int, photoRequest dto.NewPhotoRequest) (*dto.NewPhotoResponse, errs.MessageErr)
	GetPhotoById(photoId int) (*dto.NewPhotoResponse, errs.MessageErr)
	DeletePhotoById(photoId int) (*dto.NewPhotoResponse, errs.MessageErr)
	GetPhotoAll(userId int) (*dto.NewPhotoResponse, errs.MessageErr)
}

type photoService struct {
	photoRepo photo_repository.PhotoRepository
}

func NewPhotoService(photoRepo photo_repository.PhotoRepository) PhotoService {
	return &photoService{
		photoRepo: photoRepo,
	}
}

func (m *photoService) UpdatePhotoById(photoId int, photoRequest dto.NewPhotoRequest) (*dto.NewPhotoResponse, errs.MessageErr) {

	err := helpers.ValidateStruct(photoRequest)

	if err != nil {
		return nil, err
	}

	payload := entity.Photo{
		Id:       photoId,
		Title:    photoRequest.Title,
		PhotoUrl: photoRequest.PhotoUrl,
	}

	err = m.photoRepo.UpdatePhotoById(payload)

	if err != nil {
		return nil, err
	}

	response := dto.NewPhotoResponse{
		StatusCode: http.StatusOK,
		Result:     "success",
		Message:    "photo data successfully updated",
	}

	return &response, nil
}

func(m *photoService) GetPhotoAll(userId int)(*dto.NewPhotoResponse, errs.MessageErr){
	result, err := m.photoRepo.GetPhotoAll(userId)

	if err != nil {
		return nil, err
	}

	response := dto.NewPhotoResponse{
		StatusCode: http.StatusOK,
		Result:     fmt.Sprintf("%#v", result),
		Message:    "Get Photo All",
	}

	return &response, nil
}


func (m *photoService) GetPhotoById(photoId int) (*dto.NewPhotoResponse, errs.MessageErr) {

	
	result, err := m.photoRepo.GetPhotoById(photoId)

	if err != nil {
		return nil, err
	}

	response := dto.NewPhotoResponse{
		StatusCode: http.StatusOK,
		Result:     fmt.Sprintf("%+v", result),
		Message:    "Get Photo By Id Successfully",
	}

	return &response, nil
}


func (m *photoService) DeletePhotoById(photoId int) (*dto.NewPhotoResponse, errs.MessageErr) {

	
	err := m.photoRepo.DeletePhotoById(photoId)

	if err != nil {
		return nil, err
	}

	response := dto.NewPhotoResponse{
		StatusCode: http.StatusOK,
		Result:     "success",
		Message:    "photo data success deleted",
	}

	return &response, nil
}

func (m *photoService) CreatePhoto(userId int, payload dto.NewPhotoRequest) (*dto.NewPhotoResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	
	photoRequest := &entity.Photo{
		Title:    payload.Title,
		PhotoUrl: payload.PhotoUrl,
		Caption: payload.Caption,
		UserId:   userId,
	}

	_, err = m.photoRepo.CreatePhoto(photoRequest)

	if err != nil {
		return nil, err
	}

	response := dto.NewPhotoResponse{
		StatusCode: http.StatusCreated,
		Result:     "success",
		Message:    "new photo data successfully created",
	}

	return &response, err
}
