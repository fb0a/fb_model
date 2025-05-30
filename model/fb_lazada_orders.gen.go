// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameFbLazadaOrder = "fb_lazada_orders"

// FbLazadaOrder mapped from table <fb_lazada_orders>
type FbLazadaOrder struct {
	OrderID              uint64  `gorm:"column:order_id;primaryKey" json:"order_id"`
	OrderStatus          string  `gorm:"column:order_status" json:"order_status"`
	ExchangeStatus       string  `gorm:"column:exchange_status" json:"exchange_status"`
	SmsStatus            string  `gorm:"column:sms_status" json:"sms_status"`
	OrderIn              string  `gorm:"column:order_in" json:"order_in"`
	OrderItems           string  `gorm:"column:order_items" json:"order_items"`
	ExchangeCode         string  `gorm:"column:exchange_code" json:"exchange_code"`
	Price                float64 `gorm:"column:price" json:"price"`
	DigitalDeliveryInfo  string  `gorm:"column:digital_delivery_info" json:"digital_delivery_info"`
	BuyerPhoneNumber     string  `gorm:"column:buyer_phone_number" json:"buyer_phone_number"`
	ExchangerPhoneNumber string  `gorm:"column:exchanger_phone_number" json:"exchanger_phone_number"`
	ExchangedAt          int64   `gorm:"column:exchanged_at" json:"exchanged_at"`
	Product              string  `gorm:"column:product" json:"product"`
	CreatedAt            uint64  `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt            uint64  `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName FbLazadaOrder's table name
func (*FbLazadaOrder) TableName() string {
	return TableNameFbLazadaOrder
}
