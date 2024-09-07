package models

type CurierAuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetCurierById struct {
	Id string `json:"id"`
}

type CurierAuthResponse struct {
	Success  string `json:"success"`
	CurierId string `json:"curier_id"`
}
