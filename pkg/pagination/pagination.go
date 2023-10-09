package pagination

import (
	"gorm.io/gorm"
)

type Paginator interface {
	Paginate(*gorm.DB) *gorm.DB
}

type Searcher interface {
	Search(*gorm.DB) *gorm.DB
}

type Config struct {
	DefaultPageSize int
	SearchField     string
}

type Option func(*Config)

func WithDefaultPageSize(size int) Option {
	return func(c *Config) {
		c.DefaultPageSize = size
	}
}

func WithSearchField(field string) Option {
	return func(c *Config) {
		c.SearchField = field
	}
}

type DefaultPaginator struct {
	Page     int
	PageSize int
}

func (p *DefaultPaginator) Paginate(db *gorm.DB) *gorm.DB {
	offset := (p.Page - 1) * p.PageSize
	return db.Offset(offset).Limit(p.PageSize)
}

type DefaultSearcher struct {
	Query string
	Field string
}

func (s *DefaultSearcher) Search(db *gorm.DB) *gorm.DB {
	if s.Query != "" {
		db = db.Where(s.Field+" LIKE ?", "%"+s.Query+"%")
	}
	return db
}

func NewPaginationService(currentPage int, pageSize int, query string, opts ...Option) (Paginator, Searcher) {
	config := &Config{
		DefaultPageSize: pageSize,
		SearchField:     "name",
	}

	for _, o := range opts {
		o(config)
	}

	paginator := &DefaultPaginator{
		Page:     currentPage,
		PageSize: pageSize,
	}

	searcher := &DefaultSearcher{
		Query: query,
		Field: config.SearchField,
	}

	return paginator, searcher
}
