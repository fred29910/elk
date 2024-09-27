package system

import "gorm.io/gorm"

type System struct {
	gorm.Model
	Config string `gorm:"type:text;not null" json:"config"`
	Type   string `gorm:"type:varchar(20);not null" json:"type"`    // 系统类型，如：elk
	Status uint8  `gorm:"type:tinyint(3);not null" json:"status"`   // 系统状态，如：1表示运行中，0表示停止
	Remark string `gorm:"type:varchar(255);not null" json:"remark"` // 备注
}
