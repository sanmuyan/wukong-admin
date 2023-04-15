package dal

import (
	"github.com/sirupsen/logrus"
	"wukong/pkg/db"
	"wukong/pkg/page"
)

func (c *dal) Where(data any) *dal {
	c.whereData = data
	return c
}

func (c *dal) List(data any) error {
	defer c.Closed()
	if c.whereData == nil {
		err := db.DB.Find(data).Error
		if err != nil {
			return err
		}
	} else {
		err := db.DB.Where(c.whereData).Find(data).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *dal) Query(data any) error {
	defer c.Closed()
	switch *c.queryType {
	case 1:
		// 接口模糊查询
		sql := "CONCAT(" + c.queryLike.Keys + ")" + "LIKE ?"
		err := db.DB.Scopes(page.Paginate(c.page)).Where(sql, "%"+c.queryLike.Value+"%").Find(data).Error
		if err != nil {
			return err
		}
		db.DB.Model(data).Where(sql, "%"+c.queryLike.Value+"%").Count(&c.page.TotalCount)
	default:
		// 默认查询
		err := db.DB.Scopes(page.Paginate(c.page)).Where(*c.queryMust).Find(data).Error
		if err != nil {
			return err
		}
		db.DB.Model(data).Where(*c.queryMust).Count(&c.page.TotalCount)
	}
	return nil
}

func (c *dal) Get(data any) error {
	defer c.Closed()
	err := db.DB.Where(data).First(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *dal) Create(data any) error {
	defer c.Closed()
	if err := db.DB.Create(data).Error; err != nil {
		logrus.Errorln(err)
		return err
	}
	return nil
}

func (c *dal) Update(data any) error {
	defer c.Closed()
	if err := db.DB.Model(data).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func (c *dal) Delete(data any) error {
	defer c.Closed()
	// 加上 where 条件防止批量删除, where 条件 主键同时为空的时候阻止删除
	if err := db.DB.Where(data).Delete(data).Error; err != nil {
		logrus.Errorln(err)
		return err
	}
	return nil
}
