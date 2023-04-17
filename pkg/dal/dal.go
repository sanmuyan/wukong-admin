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
	Where(any) Client
	SetQuery(*QueryOptions) Client
}

var _ Client = &mysql{}

type QueryLike struct {
	Keys  string
	Value string
}

type QueryOptions struct {
	Page      *page.Page
	QueryLike QueryLike
	QueryMust map[string]any
	QueryType int
}

type mysql struct {
	*QueryOptions
	WhereData any
}

func NewDal() Client {
	return &mysql{}
}

func (c *mysql) SetQuery(query *QueryOptions) Client {
	c.QueryOptions = query
	return c
}
