package pagination

type Params struct {
	PageSize int `form:"pageSize"`
	Page     int `form:"page"`

	SearchBy string `form:"searchBy"`
	Search   string `form:"search"`

	OrderBy string `form:"orderBy"`
	Order   string `form:"order"`
}

func (p *Params) Normalize() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.PageSize <= 0 {
		p.PageSize = 10
	}

	if p.PageSize > 100 {
		p.PageSize = 100
	}

	if p.Order != "asc" && p.Order != "desc" {
		p.Order = "desc"
	}
}

func (p Params) Offset() int {
	return (p.Page - 1) * p.PageSize
}
