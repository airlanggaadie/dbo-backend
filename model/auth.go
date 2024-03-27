package model

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type LoginReport struct {
	Data  []UserLoginHistoryDetail `json:"data"`
	Total int64                    `json:"total"`
}
