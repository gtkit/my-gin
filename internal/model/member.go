// @Author xiaozhaofu 2023/6/26 20:37:00
package model

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const TimeFormat = time.DateTime

// MarshalJSON 为AssistantMember类型实现自定义的MarshalJSON方法
func (m *AssistantMember) MarshalJSON() ([]byte, error) {
	type TempMember AssistantMember // 定义与AssistantMember字段一致的新类型
	return json.Marshal(struct {
		CreatedAt   string `json:"created_at"`
		UpdatedAt   string `json:"updated_at"`
		VipEndAt    string `json:"vip_end_at"`
		*TempMember        // 避免直接嵌套AssistantMember进入死循环
	}{
		CreatedAt:  m.CreatedAt.Format(TimeFormat),
		UpdatedAt:  m.UpdatedAt.Format(TimeFormat),
		VipEndAt:   m.VipEndAt.Format(TimeFormat),
		TempMember: (*TempMember)(m),
	})
}

// UnmarshalJSON 为AssistantMember类型实现自定义的UnmarshalJSON方法
func (m *AssistantMember) UnmarshalJSON(data []byte) error {
	type TempMember AssistantMember // 定义与AssistantMember字段一致的新类型
	ot := struct {
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		VipEndAt  string `json:"vip_end_at"`

		*TempMember // 避免直接嵌套AssistantMember进入死循环
	}{
		TempMember: (*TempMember)(m),
	}
	if err := json.Unmarshal(data, &ot); err != nil {
		return err
	}

	if len(ot.CreatedAt) > 0 {
		createdat, err := time.Parse(TimeFormat, ot.CreatedAt)
		if err != nil {
			return err
		}
		m.CreatedAt = &createdat
	}
	if len(ot.UpdatedAt) > 0 {
		updatedat, err := time.Parse(TimeFormat, ot.UpdatedAt)
		if err != nil {
			return err
		}
		m.UpdatedAt = &updatedat
	}
	if len(ot.VipEndAt) > 0 {
		vipendtat, err := time.Parse(TimeFormat, ot.VipEndAt)
		if err != nil {
			return err
		}
		m.VipEndAt = &vipendtat
	}

	return nil
}

// CreateOrUpdate AssistantMember 会员信息创建或更新
func (m *AssistantMember) CreateOrUpdate(db *gorm.DB, primarykey string, updatekey []string) error {
	return db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: primarykey}},
		DoUpdates: clause.AssignmentColumns(updatekey),
	}).Create(&m).Error
}
