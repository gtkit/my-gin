// @Author xiaozhaofu 2023/2/23 14:44:00
package repository

import (
	"strconv"
	"time"

	"github.com/gtkit/goerr"

	"ydsd_gin/internal/model"
)

func (r *reposit) Ping() (*model.AssistantMember, goerr.Error) {
	var m *model.AssistantMember
	err := r.mdb.Where("mobile_phone = ?", "15605388820").First(&m).Error
	if err != nil {
		return nil, goerr.New(err, goerr.ErrMysqlServer, "mysql 查询失败")
	}
	err = m.CreateOrUpdate(r.mdb, "id", []string{"nick_name", "mobile_phone"})
	if err != nil {
		return nil, goerr.New(err, goerr.ErrMysqlServer, "mysql 更新失败")
	}

	// 用到多个redis 库的情况
	r.rdbs[0].Set("member_id", strconv.Itoa(int(m.ID)), 0)
	r.rdbs[1].Set("MobilePhone", *m.MobilePhone, 60*time.Second)

	return m, nil
}
