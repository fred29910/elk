package ov

type Token struct {
	Token        string `json:"token,omitempty"`
	TokenType    string `json:"token_type,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	ExpiresIn    int    `json:"expires_in,omitempty"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Captcha  string `json:"captcha" binding:"required"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
	Token        string `json:"token" binding:"required"`
}

// result warp
type Result struct {
	Data  interface{} `json:"data"`
	Code  int         `json:"code,omitempty"`
	Error string      `json:"msg,omitempty"`
}

// page warp
type PageInfo struct {
	Data  any    `json:"data"`
	Total int64  `json:"total"`
	Code  int    `json:"code,omitempty"`
	Error string `json:"msg,omitempty"`
}

// Register 注册
type RegisterReq struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegisterResp 注册响应
type RegisterResp struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
