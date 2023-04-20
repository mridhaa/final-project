package dto

type NewUserRequest struct {
	Username    string `json:"username" valid:"required~username cannot be empty"`
	Email    string `json:"email" valid:"required~email cannot be empty, email~format must email "`
	Password string `json:"password" valid:"stringlength(6|100)~password must minimal 6 character,required~password cannot be empty "`
	Age int `json:"age" valid:"required~age cannot be empty"`
}

type NewUserRequestLogin struct {
	Email    string `json:"email" valid:"required~email cannot be empty, email~format must email "`
	Password string `json:"password" valid:"stringlength(6|100)~password must minimal 6 character,required~password cannot be empty "`
}

type NewUserResponse struct {
	Result     string `json:"result"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type LoginResponse struct {
	Result     string        `json:"result"`
	StatusCode int           `json:"statusCode"`
	Message    string        `json:"message"`
	Data       TokenResponse `json:"data"`
}
