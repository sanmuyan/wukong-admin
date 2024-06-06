package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
	"wukong/pkg/page"
	"wukong/server/model"
)

func getQuery(c *gin.Context, likeKeys string, mustKeys []string) *model.Query {
	pageNumber, err := strconv.Atoi(c.Query("page_number"))
	if err != nil {
		pageNumber = 1
	}
	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		pageSize = 10
	}

	query := &model.Query{
		QueryLikeValue: c.Query("query"),
		QueryLikeKeys:  likeKeys,
		QueryMustMap:   make(map[string]any),
		Page: &page.Page{
			PageNumber: pageNumber,
			PageSize:   pageSize,
		},
	}

	for _, key := range mustKeys {
		value := c.Query(key)
		if len(value) > 0 {
			query.QueryMustMap[key] = value
		}
	}
	return query
}

func keysToUserToken(c *gin.Context) (userToken *model.Token) {
	_userToken, ok := c.Keys["userToken"]
	if !ok {
		logrus.Errorf("请求缺少键值 userToken keys: %+v", c.Keys)
		return userToken
	}
	return _userToken.(*model.Token)
}

func isMustQuery(c *gin.Context, params ...string) bool {
	if len(params) == 0 {
		return false
	}
	for _, param := range params {
		if c.Query(param) == "" {
			return false
		}
	}
	return true
}
