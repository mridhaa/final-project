package dto

type NewPhotoRequest struct {
	Title    string `json:"title" valid:"required~title cannot be empty" example:"Jelangkung"`
	PhotoUrl string `json:"photo_url" valid:"required~photo url url cannot be empty" example:"http://imageurl.com"`
	Caption string `json:"caption"`

}

type NewPhotoResponse struct {
	Result     string `json:"result"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}
