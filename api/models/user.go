package models

type UserAuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUserById struct {
	Id string `json:"id"`
}

type UserAuthResponse struct {
	Success string `json:"success"`
	UserId  string `json:"user_id"`
}
