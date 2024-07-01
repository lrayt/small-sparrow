package model

import (
	"gorm.io/gorm"
	"time"
)

type OrderStatus uint8

const (
	OrderStatusUnpaid        OrderStatus = iota + 1 // 未付款,
	OrderStatusPaid                                 // 已付款
	OrderStatusShipped                              // 已发货
	OrderStatusReceived                             // 已签收
	OrderStatusReturnRequest                        // 退货申请
	OrderStatusReturnDuring                         // 退货中
	OrderStatusReturnGoods                          // 已退货
	OrderStatusCancel                               // 取消交易
)

type PayChannel uint8

const (
	PayChannelAlipay = iota + 1
	PayChannelWXPay
)

// 订单信息
type OrderInfo struct {
	gorm.Model
	OrderNO      string      `gorm:"column:order_no;index;type:varchar(32);not null;comment:订单单号"`
	ProductId    string      `gorm:"column:product_id;index;type:varchar(32);not null;comment:商品ID"`
	Status       OrderStatus `gorm:"column:status;index;not null;comment:订单状态"`
	ProductCount uint32      `gorm:"column:product_count;not null;comment:商品数量"`
	AmountTotal  float64     `gorm:"column:amount_total;not null;comment:商品总价"`
	LogisticsFee float64     `gorm:"column:logistics_fee;not null;comment:运费金额"`
	AddressId    string      `gorm:"column:address_id;index;type:varchar(32);not null;comment:收货地址编号"`
	PayDate      *time.Time  `gorm:"column:pay_date;index;type:varchar(32);not null;comment:支付时间"`
	UserId       string      `gorm:"column:user_id;index;type:varchar(32);not null;comment:用户ID"`
	PayWay       PayChannel  `gorm:"column:pay_way;not null;comment:订单支付渠道"`
	OutTradeNO   string      `gorm:"column:out_trade_no;index;type:varchar(32);not null;comment:第三方支付流水号"`
}
