package general

import "gopkg.in/guregu/null.v4"

type PaginationData struct {
	Offset    int         `json:"-"`
	Page      int         `json:"page"`
	Limit     int         `json:"limit"`
	Sort      string      `json:"sort"`
	OrderBy   null.String `json:"order_by"`
	TotalPage int         `json:"total_page"`
	TotalData int         `json:"total_data"`
	IsGetAll  bool        `json:"is_get_all"`
}

func (pd *PaginationData) SetOffset() {
	pd.Offset = (pd.Page - 1) * pd.Limit
}

type TotalData struct {
	Total int `db:"count"`
}

func GetPagination() PaginationData {
	return PaginationData{
		Page:   1,
		Offset: 0,
		Limit:  10,
		Sort:   "asc",
	}
}

type PaginationRequest struct {
	Page  int `json:"page" validate:"gte=1"`
	Limit int `json:"limit" validate:"gte=1"`
}
