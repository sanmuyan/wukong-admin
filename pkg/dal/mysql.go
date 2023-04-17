package dal

import (
	"github.com/sirupsen/logrus"
	"wukong/pkg/db"
	"wukong/pkg/page"
)

func (c *mysql) Where(data any) Client {
	c.WhereData = data
	return c
}

func (c *mysql) List(data any) error {
	if c.WhereData == nil {
		err := db.DB.Find(data).Error
		if err != nil {
			return err
		}
	} else {
		err := db.DB.Where(c.WhereData).Find(data).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *mysql) Query(data any) error {
	switch c.QueryType {
	case 1:
		// 接口模糊查询
		sql := "CONCAT(" + c.QueryLike.Keys + ")" + "LIKE ?"
		err := db.DB.Scopes(page.Paginate(c.Page)).Where(sql, "%"+c.QueryLike.Value+"%").Find(data).Error
		if err != nil {
			return err
		}
		db.DB.Model(data).Where(sql, "%"+c.QueryLike.Value+"%").Count(&c.Page.TotalCount)
	default:
		// 默认查询
		err := db.DB.Scopes(page.Paginate(c.Page)).Where(c.QueryMust).Find(data).Error
		if err != nil {
			return err
		}
		db.DB.Model(data).Where(c.QueryMust).Count(&c.Page.TotalCount)
	}
	return nil
}

func (c *mysql) Get(data any) error {
	err := db.DB.Where(data).First(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *mysql) Create(data any) error {
	if err := db.DB.Create(data).Error; err != nil {
		logrus.Errorln(err)
		return err
	}
	return nil
}

func (c *mysql) Update(data any) error {
	if err := db.DB.Model(data).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func (c *mysql) Delete(data any) error {
	// 加上 where 条件防止批量删除, where 条件 主键同时为空的时候阻止删除
	if err := db.DB.Where(data).Delete(data).Error; err != nil {
		logrus.Errorln(err)
		return err
	}
	return nil
}
