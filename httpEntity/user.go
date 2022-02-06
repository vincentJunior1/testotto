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
