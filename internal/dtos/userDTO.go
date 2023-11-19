package dtos

type UserDTO struct {
	Name     string  `json:"name"`
	Email    string `json:"email"`
	Password string  `json:"password"`
}

type GetJWTDTO struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type GetJwtOutput struct {
	AccessToken string `json:"acess_token"`
}
