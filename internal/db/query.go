package db

type QueryParams struct {
	Fields   []string
	Filters  map[string]interface{}
	Page     int
	PageSize int
}

const (
	RangeQueryLt  = "lt"
	RangeQueryGt  = "gt"
	RangeQueryGte = "gte"
	RangeQuerylte = "gte"
)
