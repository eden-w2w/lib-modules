package order

import (
	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules/freight_trial"
	"github.com/shopspring/decimal"
)

func ToDiscountAmount(
	model *databases.MarketingDiscount,
	goods []freight_trial.FreightTrialParams,
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
					discountPrice = uint64(decimal.NewFromInt(1).
						Sub(decimal.NewFromFloat(model.DiscountRate)).
						Mul(decimal.NewFromInt(int64(item.Price))).IntPart())
					discountAmount += discountPrice * uint64(item.Amount)
				}
			}
		} else if model.Cal == enums.DISCOUNT_CAL__MULTISTEP_UNIT {
			if model.Type == enums.DISCOUNT_TYPE__ALL {
				discountPrice = model.MultiStepReduction.DiscountAmount(totalAmount)
				discountAmount += discountPrice * uint64(item.Amount)
			} else if model.Type == enums.DISCOUNT_TYPE__ALL_PERCENT {
				discountPrice = model.MultiStepRate.DiscountAmountUnit(totalAmount, item.Price)
				discountAmount += discountPrice * uint64(item.Amount)
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
