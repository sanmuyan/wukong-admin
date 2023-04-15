package page

import (
	"gorm.io/gorm"
)

type Page struct {
	PageNumber int   `json:"page_number,omitempty"`
	PageSize   int   `json:"page_size,omitempty"`
	TotalCount int64 `json:"total_count,omitempty"`
}

func Paginate(page *Page) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page.PageNumber == 0 {
			page.PageNumber = 1
		}
		switch {
		case page.PageSize > 100:
			page.PageSize = 100
		case page.PageSize <= 0:
			page.PageSize = 10
		}
		offset := (page.PageNumber - 1) * page.PageSize
		return db.Offset(offset).Limit(page.PageSize)
	}
}
