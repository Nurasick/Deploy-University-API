package model

//models needed for responses to authentication requests
type AuthResponse struct {
	Token string `json:"token"`
}

type MeResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Role  int    `json:"role"`
}
