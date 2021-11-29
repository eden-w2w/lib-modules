package databases

import (
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/eden-w2w/lib-modules/constants/types"
)

//go:generate eden generate model MarketingDiscount --database Config.DB --with-comments
//go:generate eden generate tag MarketingDiscount --defaults=true
// @def primary ID
// @def unique_index U_discount_id DiscountID
// @def I_default Status Type Cal
type MarketingDiscount struct {
	datatypes.PrimaryID
	// 业务ID
	DiscountID uint64 `json:"discountID,string" db:"f_discount_id"`
	// 营销名称
	Name string `json:"name" db:"f_name"`
	// 营销类型
	Type enums.DiscountType `json:"type" db:"f_type"`
	// 营销状态
	Status enums.DiscountStatus `json:"status" db:"f_status"`
	// 计算方式
	Cal enums.DiscountCal `json:"cal" db:"f_cal"`
	// 有效期开始
	ValidityStart datatypes.MySQLTimestamp `json:"validityStart" db:"f_validity_start"`
	// 有效期结束
	ValidityEnd datatypes.MySQLTimestamp `json:"validityEnd" db:"f_validity_end"`
	// 单用户优惠次数上限
	UserLimit uint32 `json:"userLimit" db:"f_user_limit,default=0"`
	// 总数限制
	Limit uint32 `json:"limit" db:"f_limit,default=0"`
	// 已优惠次数
	Times uint64 `json:"times" db:"f_times,default=0"`
	// 优惠上限
	DiscountLimit uint64 `json:"discountLimit" db:"f_discount_limit,default=0"`

	// 总价满额单价优惠阈值
	MinTotalPrice uint64 `json:"minTotalPrice" db:"f_min_total_price,default=0"`
	// 单价折扣比例
	DiscountRate float64 `json:"discountRate" db:"f_discount_rate,null"`
	// 阶梯式折扣比例
	MultiStepRate types.MultiStepRateConfig `json:"multiStepRate" db:"f_multi_step_rate,null"`

	// 单价立减金额
	DiscountAmount uint64 `json:"discountAmount" db:"f_discount_amount,null"`
	// 阶梯式立减金额
	MultiStepReduction types.MultiStepReductionConfig `json:"multiStepReduction" db:"f_multi_step_reduction,null"`

	datatypes.OperateTime
}
