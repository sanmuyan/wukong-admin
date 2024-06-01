package service

import (
	"sort"
	"wukong/pkg/db"
	"wukong/pkg/page"
	"wukong/server/model"
)

func setQuery(q *model.Query) *model.Query {
	if len(q.QueryLikeValue) != 0 && len(q.QueryLikeKeys) != 0 {
		q.QueryType = 1
	}
	if len(q.QueryMustMap) != 0 {
		q.QueryType = 0
	}
	return q
}

func queryData(query *model.Query, data model.List) error {
	var err error
	q := setQuery(query)
	defer func() {
		data.GetPage().SetPage(q.Page)
	}()
	switch query.QueryType {
	case 1:
		// 接口模糊查询
		sql := "CONCAT(" + q.QueryLikeKeys + ")" + "LIKE ?"
		err = db.DB.Scopes(page.Paginate(query.Page)).Where(sql, "%"+query.QueryLikeValue+"%").Find(data.GetData()).Error
		if err != nil {
			return err
		}
		err = db.DB.Model(data.GetData()).Where(sql, "%"+q.QueryLikeValue+"%").Count(&q.Page.TotalCount).Error
		if err != nil {
			return err
		}
	default:
		// 默认查询
		err = db.DB.Scopes(page.Paginate(q.Page)).Where(q.QueryMustMap).Find(data.GetData()).Error
		if err != nil {
			return err
		}
		err = db.DB.Model(data.GetData()).Where(q.QueryMustMap).Count(&q.Page.TotalCount).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func getMaxAccessLevel(roles []model.Role) int {
	var accessLevels []int
	var accessLevel int
	for _, role := range roles {
		accessLevel = role.AccessLevel
		accessLevels = append(accessLevels, accessLevel)
	}
	if len(accessLevels) != 0 {
		sort.Ints(accessLevels)
		accessLevel = accessLevels[len(accessLevels)-1]
	}
	return accessLevel
}
