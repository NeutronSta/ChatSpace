package modules

import "time"

type Message struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	User      uint   `gorm:"column:user1_id;type:bigint unsigned;index" json:"user1" form:"user1"`
	Friend    uint   `gorm:"column:user2_id;type:bigint unsigned;index" json:"user2" form:"user2"`
	Message   string `gorm:"column:message;type:text;not null" json:"message" form:"message"`
	State     bool   `gorm:"column:state;type:tinyint(1)" json:"state" form:"state"`
	Type      string `gorm:"column:type;type:VARCHAR(20)" json:"type" form:"type"`
}
