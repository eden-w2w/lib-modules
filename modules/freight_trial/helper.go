package freight_trial

import (
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/eden-w2w/lib-modules/constants/types"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules/freight_template"
	"github.com/eden-w2w/lib-modules/pkg/search"
	"math"
)

type FreightTrialParams struct {
	databases.Goods
	Amount        uint32
	IsBooking     datatypes.Bool
	BookingFlowID uint64
}

type GoodsFreight struct {
	FreightCal  enums.FreightCal
	Amount      uint32
	TotalWeight uint32
	UseDefault  bool
	Template    *databases.FreightTemplate
	Rule        *databases.FreightTemplateRules
}

func FreightTrial(
	goodsList []FreightTrialParams,
	shipping *databases.ShippingAddress,
) (freightAmount uint64, freightName string, err error) {
	var goodsFreight = make([]GoodsFreight, 0)
	for _, goods := range goodsList {
		template, err := freight_template.GetController().GetTemplateByID(goods.FreightTemplateID, nil, false)
		if err != nil {
			return 0, "", err
		}
		rules, err := freight_template.GetController().GetTemplateRules(
			template.TemplateID,
			freight_template.GetTemplateRuleParams{},
		)
		if err != nil {
			return 0, "", err
		}

		// 检测是否命中规则
		var hitRule *databases.FreightTemplateRules
		for _, rule := range rules {
			ok, _, _ := search.In(
				rule.Area, shipping, func(current interface{}, needle interface{}) bool {
					var area = current.(types.FreightRuleArea)
					if area.Level == enums.AREA_LEVEL__PROVINCE && shipping.ProvinceCode == area.ADCode {
						return true
					}
					if area.Level == enums.AREA_LEVEL__CITY && shipping.CityCode == area.ADCode {
						return true
					}
					if area.Level == enums.AREA_LEVEL__DISTRICT && shipping.DistrictCode == area.ADCode {
						return true
					}
					return false
				},
			)
			if ok {
				hitRule = &rule
				break
			}
		}
		if hitRule == nil {
			// 未命中规则则使用模板默认配置
			goodsFreight = append(
				goodsFreight, GoodsFreight{
					Template:    template,
					FreightCal:  template.Cal,
					Amount:      goods.Amount,
					TotalWeight: goods.Amount * goods.UnitNetWeight,
					UseDefault:  true,
				},
			)
		} else {
			goodsFreight = append(
				goodsFreight, GoodsFreight{
					Template:    template,
					FreightCal:  template.Cal,
					Amount:      goods.Amount,
					TotalWeight: goods.Amount * goods.UnitNetWeight,
					UseDefault:  false,
					Rule:        hitRule,
				},
			)
		}
	}

	// 计算运费，按重量和按数量的运费分开计算，所有规则按照总量进行计算
	// 1. 例如有两件商品A-10件、B-20件，两件商品同时应用运费模板【首件10件，10元，每增加1件增加2元】，那么总运费计算公式为
	// 		首件价格10元10件 + ((A数量10件 + B数量20件) - (模板首件范围10件)) * 2元 = 50元
	// 2. 假如AB两件商品应用不同的模板则首件价格及续件价格都取其单价最大值的项，例如A应用模板【首件1件，10元，每增加1件增加1元】、
	// B应用模板【首件10件，20元，每增加1件增加2元】，那么总运费计算公式为
	// 		首件价格10元1件 + ((A数量10件 + B数量20件) - (模板首件范围1件)) * 2元 = 68元
	var totalAmount, totalWeight uint32
	var countFirstRange, countFirstPrice uint32
	var countContinueRange, countContinuePrice uint32
	var weightFirstRange, weightFirstPrice uint32
	var weightContinueRange, weightContinuePrice uint32
	for _, freight := range goodsFreight {
		var firstRange, firstPrice, continueRange, continuePrice uint32
		if freight.UseDefault {
			freightName = "默认运费"
			firstRange = freight.Template.FirstRange
			firstPrice = freight.Template.FirstPrice
			continueRange = freight.Template.ContinueRange
			continuePrice = freight.Template.ContinuePrice
		} else {
			freightName = freight.Rule.Description
			firstRange = freight.Rule.FirstRange
			firstPrice = freight.Rule.FirstPrice
			continueRange = freight.Rule.ContinueRange
			continuePrice = freight.Rule.ContinuePrice
		}
		if freight.FreightCal == enums.FREIGHT_CAL__COUNT {
			if !freight.UseDefault && freight.Rule.IsFreeFreight.False() {
				totalAmount += freight.Amount
			}
			if (countFirstRange == 0 && countFirstPrice == 0) || (float64(firstPrice)/float64(firstRange)) > (float64(countFirstPrice)/float64(countFirstRange)) {
				countFirstRange = firstRange
				countFirstPrice = firstPrice
			}
			if (countContinueRange == 0 && countContinuePrice == 0) || (float64(continuePrice)/float64(continueRange)) > (float64(countContinuePrice)/float64(countContinueRange)) {
				countContinueRange = continueRange
				countContinuePrice = continuePrice
			}
		} else {
			if !freight.UseDefault && freight.Rule.IsFreeFreight.False() {
				totalWeight += freight.TotalWeight
			}
			if (weightFirstRange == 0 && weightFirstPrice == 0) || (float64(firstPrice)/float64(firstRange)) > (float64(weightFirstPrice)/float64(weightFirstRange)) {
				weightFirstRange = firstRange
				weightFirstPrice = firstPrice
			}
			if (weightContinueRange == 0 && weightContinuePrice == 0) || (float64(continuePrice)/float64(continueRange)) > (float64(weightContinuePrice)/float64(weightContinueRange)) {
				weightContinueRange = continueRange
				weightContinuePrice = continuePrice
			}
		}
	}

	if totalAmount > 0 {
		freightAmount += uint64(countFirstPrice)
		if totalAmount > countFirstRange {
			freightAmount += uint64(math.Ceil(float64(totalAmount-countFirstRange)/float64(countContinueRange))) * uint64(countContinuePrice)
		}
	}
	if totalWeight > 0 {
		freightAmount += uint64(weightFirstPrice)
		if totalAmount > weightFirstRange {
			freightAmount += uint64(math.Ceil(float64(totalAmount-weightFirstRange)/float64(weightContinueRange))) * uint64(weightContinuePrice)
		}
	}

	return
}
