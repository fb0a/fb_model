// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameFbShop = "fb_shops"

// FbShop mapped from table <fb_shops>
type FbShop struct {
	ID           int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name         string `gorm:"column:name;not null;comment:门店名" json:"name"`                   // 门店名
	Address      string `gorm:"column:address;not null;comment:门店地址" json:"address"`            // 门店地址
	Domains      string `gorm:"column:domains;comment:门店域名" json:"domains"`                     // 门店域名
	State        int32  `gorm:"column:state;not null;default:1;comment:1:开启 2:关闭" json:"state"` // 1:开启 2:关闭
	CreatedAt    int64  `gorm:"column:created_at;not null" json:"created_at"`
	Note         string `gorm:"column:note;not null;comment:备注" json:"note"`                     // 备注
	ReviewerTime int64  `gorm:"column:reviewer_time;not null;comment:操作时间" json:"reviewer_time"` // 操作时间
	Reviewer     string `gorm:"column:reviewer;not null;comment:操作人" json:"reviewer"`            // 操作人
}

// TableName FbShop's table name
func (*FbShop) TableName() string {
	return TableNameFbShop
}
