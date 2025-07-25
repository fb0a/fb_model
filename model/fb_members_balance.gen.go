// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameFbMembersBalance = "fb_members_balance"

// FbMembersBalance 会员余额表
type FbMembersBalance struct {
	UID          int64   `gorm:"column:uid;primaryKey" json:"uid"`
	Balance      float64 `gorm:"column:balance;not null;comment:余额" json:"balance"`                                 // 余额
	Available    float64 `gorm:"column:available;not null;default:0.00000000;comment:可用余额" json:"available"`        // 可用余额
	Withdrawable float64 `gorm:"column:withdrawable;not null;default:0.00000000;comment:可提现余额" json:"withdrawable"` // 可提现余额
	CreatedAt    int64   `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`                         // 创建时间
	UpdatedAt    int64   `gorm:"column:updated_at;not null;comment:创建时间" json:"updated_at"`                         // 创建时间
}

// TableName FbMembersBalance's table name
func (*FbMembersBalance) TableName() string {
	return TableNameFbMembersBalance
}
