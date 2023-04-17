package service

import (
	"wukong/pkg/dal"
	"wukong/pkg/page"
	"wukong/server/model"
)

func setQuery(query *model.Query) *dal.QueryOptions {
	opt := &dal.QueryOptions{
		Page: &page.Page{},
	}
	opt.Page.PageNumber = query.PageNumber
	opt.Page.PageSize = query.PageSize

	if len(query.QueryLikeValue) != 0 && len(query.QueryLikeKeys) != 0 {
		opt.QueryType = 1
		opt.QueryLike.Keys = query.QueryLikeKeys
		opt.QueryLike.Value = query.QueryLikeValue
	}

	if len(query.QueryMustMap) != 0 {
		opt.QueryType = 0
		opt.QueryMust = query.QueryMustMap
	}

	return opt
}
