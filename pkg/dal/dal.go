package dal

import (
	"wukong/pkg/page"
)

type Client interface {
	List(any) error
	Query(any) error
	Get(any) error
	Create(any) error
	Update(any) error
	Delete(any) error
	Where(any) *dal
	Closed()
}

var _ Client = &dal{}

type QueryLike struct {
	Keys  string
	Value string
}

type dal struct {
	page      *page.Page
	queryLike *QueryLike
	queryMust *map[string]any
	queryType *int
	whereData any
}

type Options struct {
	Page      page.Page
	QueryLike QueryLike
	QueryMust map[string]any
	QueryType int
}

func NewDal(opt *Options) Client {
	d := &dal{
		page:      &opt.Page,
		queryMust: &opt.QueryMust,
		queryLike: &opt.QueryLike,
		queryType: &opt.QueryType,
	}
	return d
}

func (c *dal) Closed() {
	// 避免不当使用造成的数据污染
	var d dal
	c.queryLike = d.queryLike
	c.queryMust = d.queryMust
	c.queryType = d.queryType
	c.whereData = d.whereData
}
