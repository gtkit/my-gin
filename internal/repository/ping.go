// @Author xiaozhaofu 2023/2/23 14:44:00
package repository

import (
	"ydsd_gin/internal/model"
)

func (r *reposit) Ping() (*model.AssistantMember, error) {
	var m *model.AssistantMember
	err := r.mdb.Where("mobile_phone = ?", "15605388820").First(&m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}
