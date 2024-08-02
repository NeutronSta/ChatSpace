package modules

type FriendShip struct {
	ID     uint `gorm:"primarykey"`
	User   uint `gorm:"column:user1_id;type:bigint unsigned;index" json:"user1" form:"user1"`
	Friend uint `gorm:"column:user2_id;type:bigint unsigned;index" json:"user2" form:"user2"`
	State  bool `gorm:"column:state;type:tinyint(1)" json:"state" form:"state"`
}
