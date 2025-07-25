// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameFbPromoGameList = "fb_promo_game_list"

// FbPromoGameList 活动统计的游戏
type FbPromoGameList struct {
	Pid       int64  `gorm:"column:pid;primaryKey;comment:活动ID" json:"pid"`                      // 活动ID
	GameID    string `gorm:"column:game_id;primaryKey;comment:游戏ID" json:"game_id"`              // 游戏ID
	EnName    string `gorm:"column:en_name;not null;comment:英文名称" json:"en_name"`                // 英文名称
	State     int32  `gorm:"column:state;not null;default:1;comment:状态 1:未启用 2:启用" json:"state"` // 状态 1:未启用 2:启用
	UpdatedAt int64  `gorm:"column:updated_at;not null;comment:更新时间" json:"updated_at"`          // 更新时间
	CreatedAt int64  `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`          // 创建时间
}

// TableName FbPromoGameList's table name
func (*FbPromoGameList) TableName() string {
	return TableNameFbPromoGameList
}
