// @Author xiaozhaofu 2023/2/23 14:44:00
package repository

import (
	"context"
	"strconv"

	"github.com/gtkit/goerr"

	"my_gin/internal/model"
)

func (r *reposit) Ping() (*model.AssistantMember, goerr.Error) {
	var m *model.AssistantMember
	err := r.mdb().Where("mobile_phone = ?", "15605388820").First(&m).Error
	if err != nil {
		return nil, goerr.New(err, goerr.MysqlServer(), "mysql 查询失败")
	}
	err = m.CreateOrUpdate(r.mdb(), "id", []string{"nick_name", "mobile_phone"})
	if err != nil {
		return nil, goerr.New(err, goerr.MysqlServer(), "mysql 更新失败")
	}

	// 用到多个redis 库的情况
	r.rdbClient(0).Set(context.Background(), "member_id", strconv.Itoa(int(m.ID)), 0)
	// r.rdbs[1].Set("MobilePhone", *m.MobilePhone, 60*time.Second)

	// return m, nil
	return nil, nil
}
