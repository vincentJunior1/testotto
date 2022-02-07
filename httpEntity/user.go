package httpEntity

type UserRequest struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	RoleID   uint   `json:"role_id"`
	Password string `json:"password"`
	Status   uint   `json:"status"`
}

type LoginUser struct {
	Email    string `json:"Email"`
	Password string `json:"password"`
}

type UserResponse struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	IsActive *bool  `json:"is_active"`
}

type UserDataResponse struct {
	UserID   uint   `json:"user_id"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	RoleName string `json:"role_name"`
}
