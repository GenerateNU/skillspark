package utils

type Pagination struct {
	Page  int `query:"page" minimum:"1" default:"1" doc:"Page number for pagination"`
	Limit int `query:"limit" minimum:"1" maximum:"100" default:"100" doc:"Number of items per page"`
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
