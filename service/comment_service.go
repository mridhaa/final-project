package service

import (
	"fmt"
	"h8-movies/dto"
	"h8-movies/entity"
	"h8-movies/pkg/errs"
	"h8-movies/pkg/helpers"
	"h8-movies/repository/comment_repository"
	"net/http"
)

type CommentService interface {
	CreateComment(userId, photoId int, payload dto.NewCommentRequest) (*dto.NewCommentResponse, errs.MessageErr)
	UpdateCommentById(commentId int, commentRequest dto.NewCommentRequest) (*dto.NewCommentResponse, errs.MessageErr)
	GetCommentById(commentId int) (*dto.NewCommentResponse, errs.MessageErr)
	DeleteCommentById(commentId int) (*dto.NewCommentResponse, errs.MessageErr)
	GetCommentAll(userId int) (*dto.NewCommentResponse, errs.MessageErr)
}

type commentService struct {
	commentRepo comment_repository.CommentRepository
}

func NewCommentService(commentRepo comment_repository.CommentRepository) CommentService {
	return &commentService{
		commentRepo: commentRepo,
	}
}

func (m *commentService) UpdateCommentById(commentId int, commentRequest dto.NewCommentRequest) (*dto.NewCommentResponse, errs.MessageErr) {

	err := helpers.ValidateStruct(commentRequest)

	if err != nil {
		return nil, err
	}

	payload := entity.Comment{
		Id:       commentId,
		Message:    commentRequest.Message,
	}

	err = m.commentRepo.UpdateCommentById(payload)

	if err != nil {
		return nil, err
	}

	response := dto.NewCommentResponse{
		StatusCode: http.StatusOK,
		Result:     "success",
		Message:    "comment data successfully updated",
	}

	return &response, nil
}

func(m *commentService) GetCommentAll(userId int)(*dto.NewCommentResponse, errs.MessageErr){
	result, err := m.commentRepo.GetCommentAll(userId)

	if err != nil {
		return nil, err
	}

	response := dto.NewCommentResponse{
		StatusCode: http.StatusOK,
		Result:     fmt.Sprintf("%#v", result),
		Message:    "Get Comment All",
	}

	return &response, nil
}


func (m *commentService) GetCommentById(commentId int) (*dto.NewCommentResponse, errs.MessageErr) {

	
	result, err := m.commentRepo.GetCommentById(commentId)

	if err != nil {
		return nil, err
	}

	response := dto.NewCommentResponse{
		StatusCode: http.StatusOK,
		Result:     fmt.Sprintf("%+v", result),
		Message:    "Get Comment By Id Successfully",
	}

	return &response, nil
}


func (m *commentService) DeleteCommentById(commentId int) (*dto.NewCommentResponse, errs.MessageErr) {

	
	err := m.commentRepo.DeleteCommentById(commentId)

	if err != nil {
		return nil, err
	}

	response := dto.NewCommentResponse{
		StatusCode: http.StatusOK,
		Result:     "success",
		Message:    "comment data success deleted",
	}

	return &response, nil
}

func (m *commentService) CreateComment(userId, photoId int, payload dto.NewCommentRequest) (*dto.NewCommentResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	
	commentRequest := &entity.Comment{
		Message:    payload.Message,
		UserId:   userId,
		PhotoId: photoId,
	}

	_, err = m.commentRepo.CreateComment(commentRequest)

	if err != nil {
		return nil, err
	}

	response := dto.NewCommentResponse{
		StatusCode: http.StatusCreated,
		Result:     "success",
		Message:    "new comment data successfully created",
	}

	return &response, err
}
