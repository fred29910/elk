package ov

type UserQuery struct {
	Page     int    `form:"page" json:"page"`
	PageSize int    `form:"page_size" json:"page_size"`
	Username string `form:"username" json:"username"`
	Email    string `form:"email" json:"email"`
	SortBy   string `form:"sort_by" json:"sort_by"`
	SortDesc bool   `form:"sort_desc" json:"sort_desc"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type SystemQuery struct {
	Page     int     `form:"page" json:"page"`
	PageSize int     `form:"page_size" json:"page_size"`
	Type     *string `form:"type" json:"type"`
	Status   *uint8  `form:"status" json:"status"`
	SortBy   *string `form:"sort_by" json:"sort_by"`
	SortDesc *bool   `form:"sort_desc" json:"sort_desc"`
}

type CreateSystem struct {
	Config string `json:"config"`
	Type   string `json:"type"`
	Status uint8  `json:"status"`
	Remark string `json:"remark"`
}
