package utils

type Pagination struct {
	Page  int `query:"page" validate:"gte=1"`
	Limit int `query:"limit" validate:"gte=1"`
}

const (
	defaultPage  int = 1
	defaultLimit int = 100
)

func NewPagination() Pagination {
	return Pagination{
		Page:  defaultPage,
		Limit: defaultLimit,
	}
}

func (p *Pagination) GetOffset() int {
	page := p.Page
	limit := p.Limit

	if page < 1 {
		page = defaultPage
	}
	if limit < 1 {
		limit = defaultLimit
	}

	return (page - 1) * limit
}
