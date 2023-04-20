package comment_repository

import (
	"h8-movies/entity"
	"h8-movies/pkg/errs"
)

type CommentRepository interface {
	CreateComment(commentPayload *entity.Comment) (*entity.Comment, errs.MessageErr)
	GetCommentById(commentId int) (*entity.Comment, errs.MessageErr)
	GetCommentAll(userId int) ([]entity.Comment, errs.MessageErr)
	DeleteCommentById(commentId int) errs.MessageErr
	UpdateCommentById(payload entity.Comment) errs.MessageErr
}
