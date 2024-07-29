package ov

type Position struct {
	Start int `json:"start"`
	End   int `json:"end"`
}
type ParserReq struct {
	Content string   `json:"content"`
	Pos     Position `json:"pos"`
}
