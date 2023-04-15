package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
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
		PageNumber:     pageNumber,
		PageSize:       pageSize,
		QueryLikeValue: c.Query("query"),
		QueryLikeKeys:  likeKeys,
		QueryMustMap:   make(map[string]any),
	}

	for _, key := range mustKeys {
		value := c.Query(key)
		if len(value) > 0 {
			query.QueryMustMap[key] = value
		}
	}
	return query
}

func keysToUserToken(keys map[string]any) (userToken *model.Token) {
	keysJson, _ := json.Marshal(keys["userToken"])
	_ = json.Unmarshal(keysJson, &userToken)
	return userToken
}
