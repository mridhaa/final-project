package dto

type NewCommentRequest struct {
	Message string `json:"message" valid:"required~ message cannot be empty" example:"http://imageurl.com"`

}

type NewCommentResponse struct {
	Result     string `json:"result"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}
