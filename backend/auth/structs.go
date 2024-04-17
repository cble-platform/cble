package auth

type UserLoginInput struct {
	Username string `form:"username" json:"username" binding:"required" example:"admin"`
	Password string `form:"password" json:"password" binding:"required" example:"password123"`
}
