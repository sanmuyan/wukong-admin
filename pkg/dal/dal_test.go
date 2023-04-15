package dal

import (
	"fmt"
	"testing"
	"wukong/server/model"
)

func TestGet(t *testing.T) {
	var opt Options
	var data []model.User
	d := NewDal(&opt)
	err := d.List(&data)
	if err != nil {
		fmt.Println(err)
	}
}
