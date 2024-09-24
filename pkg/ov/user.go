package ov

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Status   uint8  `json:"status"`
}

type UpdateUserOV struct {
	ID       uint    `json:"id"`
	Username *string `json:"username"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
	Status   *uint8  `json:"status"`
}
