// @Author xiaozhaofu 2023/2/23 14:44:00
package repository

import (
	"strconv"
	"time"

	"ydsd_gin/internal/model"
)

func (r *reposit) Ping() (*model.AssistantMember, error) {
	var m *model.AssistantMember
	err := r.mdb.Where("mobile_phone = ?", "15605388820").First(&m).Error
	if err != nil {
		return nil, err
	}

	// redis 只使用一个库的时候
	// r.rdb.Set("member_name", m.NickName, 60*time.Second)
	// r.rdb.Set("member_id", strconv.Itoa(int(m.ID)), 60*time.Second)

	// 用到多个redis 库的情况
	r.rdbs[0].Set("member_id", strconv.Itoa(int(m.ID)), 0)
	r.rdbs[1].Set("MobilePhone", *m.MobilePhone, 60*time.Second)

	return m, nil
}
