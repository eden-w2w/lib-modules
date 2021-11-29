package order

import (
	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/eden-w2w/lib-modules/databases"
)

func ToDiscountAmount(
	model *databases.MarketingDiscount,
	goods []CreateOrderGoodsModelParams,
) (preGoodsList []PreCreateOrderGoodsParams, totalAmount, discountAmount uint64) {
	for _, item := range goods {
		totalAmount += item.Price * uint64(item.Amount)
	}
	for _, item := range goods {
		isBooking := item.IsBooking.True()
		discountPrice := uint64(0)
		if model.Cal == enums.DISCOUNT_CAL__UNIT {
			if model.Type == enums.DISCOUNT_TYPE__ALL {
				if item.Price > model.DiscountAmount && (model.MinTotalPrice == 0 || totalAmount >= model.MinTotalPrice) {
					if item.Price > model.DiscountAmount {
						discountPrice = model.DiscountAmount
						discountAmount += discountPrice * uint64(item.Amount)
					}
				}
			} else if model.Type == enums.DISCOUNT_TYPE__ALL_PERCENT {
				if model.MinTotalPrice == 0 || totalAmount >= model.MinTotalPrice {
					discountPrice = uint64((1.0 - model.DiscountRate) * float64(item.Price))
					discountAmount += discountPrice * uint64(item.Amount)
				}
			}
		}
		preGoodsList = append(
			preGoodsList, PreCreateOrderGoodsParams{
				GoodsID:       item.GoodsID,
				Amount:        item.Amount,
				IsBooking:     &isBooking,
				Price:         item.Price,
				DiscountPrice: discountPrice,
			},
		)
	}
	if model.Cal == enums.DISCOUNT_CAL__MULTISTEP {
		if model.Type == enums.DISCOUNT_TYPE__ALL {
			discountAmount += model.MultiStepReduction.DiscountAmount(totalAmount)
		} else if model.Type == enums.DISCOUNT_TYPE__ALL_PERCENT {
			discountAmount += model.MultiStepRate.DiscountAmount(totalAmount)
		}
	}
	return
}
