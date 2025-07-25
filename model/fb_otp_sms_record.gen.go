// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameFbOtpSmsRecord = "fb_otp_sms_record"

// FbOtpSmsRecord otp_sms表
type FbOtpSmsRecord struct {
	ID          int64  `gorm:"column:id;primaryKey" json:"id"`
	CountryCode string `gorm:"column:country_code;not null;comment:国家区号" json:"country_code"` // 国家区号
	Phone       string `gorm:"column:phone;not null;comment:手机号" json:"phone"`                // 手机号
	Code        string `gorm:"column:code;not null;comment:code" json:"code"`                 // code
	Reason      string `gorm:"column:reason;not null;comment:reason" json:"reason"`           // reason
	State       int32  `gorm:"column:state;not null;comment:0=正常" json:"state"`               // 0=正常
	CreatedAt   int64  `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`     // 创建时间
	UpdatedAt   int64  `gorm:"column:updated_at;not null" json:"updated_at"`
	ExpireAt    int64  `gorm:"column:expire_at;not null" json:"expire_at"`
	ChannelName string `gorm:"column:channel_name;not null" json:"channel_name"`
	ConsumedAt  int64  `gorm:"column:consumed_at;not null" json:"consumed_at"`
}

// TableName FbOtpSmsRecord's table name
func (*FbOtpSmsRecord) TableName() string {
	return TableNameFbOtpSmsRecord
}
