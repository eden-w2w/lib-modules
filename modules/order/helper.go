package order

import (
	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/eden-w2w/lib-modules/databases"
)

func ToDiscountAmount(
	model *databases.MarketingDiscount,
	goods []CreateOrderGoodsModelParams,
) (totalAmount, discountAmount uint64) {
	for _, item := range goods {
		totalAmount += item.Price * uint64(item.Amount)
		if model.Cal == enums.DISCOUNT_CAL__UNIT {
			if model.Type == enums.DISCOUNT_TYPE__ALL {
				discountAmount += model.DiscountAmount * uint64(item.Amount)
			} else if model.Type == enums.DISCOUNT_TYPE__ALL_PERCENT {
				discountAmount += uint64(model.DiscountRate * float64(item.Price) * float64(item.Amount))
			}
		}
	}
	if model.Cal == enums.DISCOUNT_CAL__MULTISTEP {
		if model.Type == enums.DISCOUNT_TYPE__ALL {
			discountAmount = model.MultiStepReduction.DiscountAmount(totalAmount)
		} else if model.Type == enums.DISCOUNT_TYPE__ALL_PERCENT {
			discountAmount = model.MultiStepRate.DiscountAmount(totalAmount)
		}
	}
	return
}
