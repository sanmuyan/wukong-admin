package service

import (
	"wukong/pkg/dal"
	"wukong/server/model"
)

type options struct {
	dal.Options
}

func newDal(opt *options) dal.Client {
	return dal.NewDal(&opt.Options)
}

func (opt *options) setQuery(query *model.Query) {
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
}
