package httpEntity

type UserRequest struct {
	UserName string `json:"user_name"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Status   uint   `json:"status"`
}

type LoginUser struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type UserResponse struct {
	UserName string `json:"user_name"`
	Name     string `json:"name"`
	IsActive *bool  `json:"is_active"`
}

type UserDataResponse struct {
	UserID   uint   `json:"user_id"`
	UserName string `json:"user_name"`
	Name     string `json:"name"`
}
