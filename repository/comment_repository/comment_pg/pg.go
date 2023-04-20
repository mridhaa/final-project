package comment_pg

import (
	"database/sql"
	"errors"
	"fmt"
	"h8-movies/entity"
	"h8-movies/pkg/errs"
	"h8-movies/repository/comment_repository"
)

const (
	getCommentByIdQuery = `
		SELECT id, message, user_id, photo_id, createdAt, updatedAt from "comments"
		WHERE id = $1;
	`

	updateCommentByIdQuery = `
		UPDATE "comments"
		SET message = $2
		WHERE id = $1;
	`

	deleteCommentByIdQuery = `
		DELETE from  "comments"
		WHERE id = $1;
	`


	commentAllQuery = `
		Select id, message, user_id, photo_id, createdAt, updatedAt from "comments" where user_id = $1
	`
)

type commentPG struct {
	db *sql.DB
}

func NewCommentPG(db *sql.DB) comment_repository.CommentRepository {
	return &commentPG{
		db: db,
	}
}

func (m *commentPG) UpdateCommentById(payload entity.Comment) errs.MessageErr {

	fmt.Println(payload.Id , "               ", payload.Message)
	_, err := m.db.Exec(updateCommentByIdQuery, payload.Id, payload.Message)

	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}

func (m *commentPG) GetCommentById(commentId int) (*entity.Comment, errs.MessageErr) {
	row := m.db.QueryRow(getCommentByIdQuery, commentId)

	var comment entity.Comment
	

	err := row.Scan(&comment.Id, &comment.Message, &comment.UserId, &comment.PhotoId,  &comment.CreatedAt, &comment.UpdatedAt)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("comment not found")
		}

		return nil, errs.NewInternalServerError("something went wrong in Scan")
	}

	return &comment, nil
}

func (m *commentPG) GetCommentAll(userId int) ([]entity.Comment, errs.MessageErr) {
	rows, _:= m.db.Query(commentAllQuery, userId)

	defer rows.Close()
	
	var photos []entity.Comment

	for rows.Next(){
		var photo entity.Comment
		err := rows.Scan(&photo.Id, &photo.Message, &photo.UserId, &photo.PhotoId, &photo.CreatedAt, &photo.UpdatedAt)

		if err != nil {	
			return nil, errs.NewInternalServerError("something went wrong")
		}
		photos = append(photos, photo)
	}

	return photos, nil
}

func (m *commentPG) DeleteCommentById(commentId int) errs.MessageErr {
	_,err := m.db.Exec(deleteCommentByIdQuery, commentId)


	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}

func (m *commentPG) CreateComment(commentPayload *entity.Comment) (*entity.Comment, errs.MessageErr) {
	createCommentQuery := `
		INSERT INTO "comments"
		(
			message,
			user_id,
			photo_id
		)
		VALUES($1, $2, $3)
		RETURNING id,message, user_id, photo_id;
	`
	
	row := m.db.QueryRow(createCommentQuery, commentPayload.Message, commentPayload.UserId, commentPayload.PhotoId)

	var comment entity.Comment

	err := row.Scan(&comment.Id, &comment.Message, &comment.UserId, &comment.PhotoId)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, errs.NewInternalServerError("something went wrong in Scan")
	}

	return &comment, nil

}



