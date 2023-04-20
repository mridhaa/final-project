package photo_pg

import (
	"database/sql"
	"errors"
	"fmt"
	"h8-movies/entity"
	"h8-movies/pkg/errs"
	"h8-movies/repository/photo_repository"
)

const (
	getPhotoByIdQuery = `
		SELECT id, title, user_id, photo_url, caption,  createdAt, updatedAt from "photos"
		WHERE id = $1;
	`

	updatePhotoByIdQuery = `
		UPDATE "photos"
		SET title = $2,
		photo_url = $3,
		caption = $4
		WHERE id = $1;
	`

	deletePhotoByIdQuery = `
		DELETE from  "photos"
		WHERE id = $1;
	`


	photoAllQuery = `
		Select id, title, user_id, photo_url, caption,  createdAt, updatedAt from "photos" where user_id = $1
	`
)

type photoPG struct {
	db *sql.DB
}

func NewPhotoPG(db *sql.DB) photo_repository.PhotoRepository {
	return &photoPG{
		db: db,
	}
}

func (m *photoPG) UpdatePhotoById(payload entity.Photo) errs.MessageErr {
	_, err := m.db.Exec(updatePhotoByIdQuery, payload.Id, payload.Title, payload.PhotoUrl, payload.Caption)

	if err != nil {

		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}

func (m *photoPG) GetPhotoById(photoId int) (*entity.Photo, errs.MessageErr) {
	row := m.db.QueryRow(getPhotoByIdQuery, photoId)

	var photo entity.Photo
	

	err := row.Scan(&photo.Id, &photo.Title, &photo.UserId, &photo.PhotoUrl, &photo.Caption,  &photo.CreatedAt, &photo.UpdatedAt)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("photo not found")
		}

		return nil, errs.NewInternalServerError("something went wrong in Scan")
	}

	return &photo, nil
}

func (m *photoPG) GetPhotoAll(userId int) ([]entity.Photo, errs.MessageErr) {
	rows, _:= m.db.Query(photoAllQuery, userId)

	defer rows.Close()
	
	var photos []entity.Photo

	for rows.Next(){
		var photo entity.Photo
		err := rows.Scan(&photo.Id, &photo.Title, &photo.UserId, &photo.PhotoUrl, &photo.Caption,  &photo.CreatedAt, &photo.UpdatedAt)

		if err != nil {	
			return nil, errs.NewInternalServerError("something went wrong")
		}
		photos = append(photos, photo)
	}

	return photos, nil
}

func (m *photoPG) DeletePhotoById(photoId int) errs.MessageErr {
	_,err := m.db.Exec(deletePhotoByIdQuery, photoId)


	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}

func (m *photoPG) CreatePhoto(photoPayload *entity.Photo) (*entity.Photo, errs.MessageErr) {
	createPhotoQuery := `
		INSERT INTO "photos"
		(
			title,
			photo_url,
			caption,
			user_id
		)
		VALUES($1, $2, $3, $4)
		RETURNING id,title, photo_url, caption, user_id;
	`
	row := m.db.QueryRow(createPhotoQuery, photoPayload.Title, photoPayload.PhotoUrl, photoPayload.Caption,  photoPayload.UserId)

	fmt.Println(photoPayload.Title,photoPayload.PhotoUrl, photoPayload.Caption, photoPayload.UserId)

	var photo entity.Photo

	err := row.Scan(&photo.Id, &photo.Title, &photo.PhotoUrl, &photo.Caption , &photo.UserId)

	if err != nil {
	
		fmt.Printf("err: %v\n", err)
		return nil, errs.NewInternalServerError("something went wrong in Scan")
	}

	return &photo, nil

}



