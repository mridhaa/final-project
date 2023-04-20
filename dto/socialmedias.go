package dto

type NewSocialMediasRequest struct {
	Name    string `json:"name" valid:"required~name cannot be empty" example:"Jelangkung"`
	SocialMediaUrl string `json:"social_media_url" valid:"required~social media url cannot be empty" example:"http://imageurl.com"`
}

type NewSocialMediasResponse struct {
	Result     string `json:"result"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}
