package app

// type SortDirection string

// const (
// 	ASC  SortDirection = "ASC"
// 	DESC SortDirection = "DESC"
// )

type Query struct {
	Take          int    `form:"take,omitempty"`
	Page          int    `form:"page,omitempty"`
	SortField     string `form:"sortField,omitempty"`
	SortDirection string `form:"sortDirection,omitempty" enums:"ASC,DESC"`
	Search        string `form:"search,omitempty"`
}
