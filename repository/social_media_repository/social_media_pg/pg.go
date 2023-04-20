package social_media_pg

import (
	"database/sql"
	"errors"
	"fmt"
	"h8-movies/entity"
	"h8-movies/pkg/errs"
	"h8-movies/repository/social_media_repository"
)

const (
	getSMByIdQuery = `
		SELECT id, name, social_media_url, user_id, createdAt, updatedAt from "socialmedias"
		WHERE id = $1;
	`

	updateSMByIdQuery = `
		UPDATE "socialmedias"
		SET name = $2,
		social_media_url = $3
		WHERE id = $1;
	`

	deleteSMByIdQuery = `
		DELETE from  "socialmedias"
		WHERE id = $1;
	`


	smAllQuery = `
		Select id, name, social_media_url, user_id, createdAt, updatedAt from "socialmedias" where user_id = $1
	`
)

type socialmediaPG struct {
	db *sql.DB
}

func NewSocialMediaPG(db *sql.DB) social_media_repository.SocialMediaRepository {
	return &socialmediaPG{
		db: db,
	}
}

func (m *socialmediaPG) UpdateSocialMediaById(payload entity.SocialMedias) errs.MessageErr {
	_, err := m.db.Exec(updateSMByIdQuery, payload.Id, payload.Name, payload.SocialMediaUrl)

	if err != nil {

		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}

func (m *socialmediaPG) GetSocialMediaById(socialmediaId int) (*entity.SocialMedias, errs.MessageErr) {
	row := m.db.QueryRow(getSMByIdQuery, socialmediaId)

	var socialmedia entity.SocialMedias
	

	err := row.Scan(&socialmedia.Id, &socialmedia.Name, &socialmedia.SocialMediaUrl, &socialmedia.UserId, &socialmedia.CreatedAt, &socialmedia.UpdatedAt)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("socialmedia not found")
		}

		return nil, errs.NewInternalServerError("something went wrong in Scan")
	}

	return &socialmedia, nil
}

func (m *socialmediaPG) GetSocialMediaAll(userId int) ([]entity.SocialMedias, errs.MessageErr) {
	rows, _:= m.db.Query(smAllQuery, userId)

	defer rows.Close()
	
	var socialmedias []entity.SocialMedias

	for rows.Next(){
		var socialmedia entity.SocialMedias
		err := rows.Scan(&socialmedia.Id, &socialmedia.Name, &socialmedia.SocialMediaUrl, &socialmedia.UserId,  &socialmedia.CreatedAt, &socialmedia.UpdatedAt)

		if err != nil {	
			return nil, errs.NewInternalServerError("something went wrong")
		}
		socialmedias = append(socialmedias, socialmedia)
	}

	return socialmedias, nil
}

func (m *socialmediaPG) DeleteSocialMediaById(socialmediaId int) errs.MessageErr {
	_,err := m.db.Exec(deleteSMByIdQuery, socialmediaId)


	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}

func (m *socialmediaPG) CreateSocialMedia(photoPayload *entity.SocialMedias) (*entity.SocialMedias, errs.MessageErr) {
	createSosialMeduaQuery := `
		INSERT INTO "socialmedias"
		(
			name,
			social_media_url,
			user_id
		)
		VALUES($1, $2, $3)
		RETURNING id,name, social_media_url, user_id;
	`
	row := m.db.QueryRow(createSosialMeduaQuery, photoPayload.Name, photoPayload.SocialMediaUrl,  photoPayload.UserId)



	var socialmedia entity.SocialMedias

	err := row.Scan(&socialmedia.Id, &socialmedia.Name, &socialmedia.SocialMediaUrl, &socialmedia.UserId)

	if err != nil {
	
		fmt.Printf("err: %v\n", err)
		return nil, errs.NewInternalServerError("something went wrong in Scan")
	}

	return &socialmedia, nil

}



