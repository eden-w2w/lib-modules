package booking_flow

import (
	"github.com/eden-framework/sqlx/builder"
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules"
	"time"
)

type CreateBookingFlowParams struct {
	// 商品ID
	GoodsID uint64 `json:"goodsID,string"`
	// 预售限量
	Limit uint32 `json:"limit" default:""`
	// 预售模式
	Type enums.BookingType `json:"type"`
	// 预售开始时间
	StartTime datatypes.MySQLTimestamp `json:"startTime"`
	// 预售结束时间
	EndTime datatypes.MySQLTimestamp `json:"endTime" default:""`
	// 预计到货时间
	EstimatedTimeArrival datatypes.MySQLTimestamp `json:"eta" default:""`
}

func (c CreateBookingFlowParams) Model() *databases.BookingFlow {
	model := &databases.BookingFlow{
		GoodsID:              c.GoodsID,
		Limit:                c.Limit,
		Type:                 c.Type,
		Status:               enums.BOOKING_STATUS__READY,
		StartTime:            c.StartTime,
		EndTime:              c.EndTime,
		EstimatedTimeArrival: c.EstimatedTimeArrival,
	}
	if time.Now().Before(time.Time(c.StartTime)) {
		model.Status = enums.BOOKING_STATUS__PROCESS
	}
	return model
}

type UpdateBookingFlowParams struct {
	// 预售销量
	Sales *uint32 `json:"sales" default:""`
	// 预售限量
	Limit *uint32 `json:"limit" default:""`
	// 预售模式
	Type enums.BookingType `json:"type" default:""`
	// 预售状态
	Status enums.BookingStatus `json:"status" default:""`
	// 预售开始时间
	StartTime datatypes.MySQLTimestamp `json:"startTime" default:""`
	// 预售结束时间
	EndTime datatypes.MySQLTimestamp `json:"endTime" default:""`
	// 预计到货时间
	EstimatedTimeArrival datatypes.MySQLTimestamp `json:"eta" default:""`
}

func (p *UpdateBookingFlowParams) Fill(model *databases.BookingFlow) (zeroFields []string) {
	zeroFields = make([]string, 0)
	if model == nil {
		return
	}
	if p.Sales != nil {
		model.Sales = *p.Sales
		zeroFields = append(zeroFields, model.FieldKeySales())
	}
	if p.Limit != nil {
		model.Limit = *p.Limit
		zeroFields = append(zeroFields, model.FieldKeyLimit())
	}
	if p.Type != enums.BOOKING_TYPE_UNKNOWN {
		model.Type = p.Type
	}
	if p.Status != enums.BOOKING_STATUS_UNKNOWN {
		model.Status = p.Status
	}
	if p.StartTime != datatypes.TimestampZero {
		model.StartTime = p.StartTime
	}
	if p.EndTime != datatypes.TimestampZero {
		model.EndTime = p.EndTime
	}
	if p.EstimatedTimeArrival != datatypes.TimestampZero {
		model.EstimatedTimeArrival = p.EstimatedTimeArrival
	}
	return
}

type GetBookingFlowParams struct {
	// 商品ID
	GoodsID uint64 `in:"query" name:"goodsID,string" default:""`
	// 预售模式
	Type enums.BookingType `in:"query" name:"type" default:""`
	// 预售状态
	Status enums.BookingStatus `in:"query" name:"status" default:""`
	// 预售开始时间大于等于
	StartTimeBegin datatypes.MySQLTimestamp `in:"query" name:"startTimeBegin" default:""`
	// 预售开始时间小于
	StartTimeEnd datatypes.MySQLTimestamp `in:"query" name:"startTimeEnd" default:""`
	// 预售结束时间大于等于
	EndTimeBegin datatypes.MySQLTimestamp `in:"query" name:"endTimeBegin" default:""`
	// 预售结束时间小于
	EndTimeEnd datatypes.MySQLTimestamp `in:"query" name:"endTimeEnd" default:""`
	modules.Pagination
}

func (p GetBookingFlowParams) Conditions() builder.SqlCondition {
	var condition builder.SqlCondition
	var model = databases.BookingFlow{}
	if p.GoodsID != 0 {
		condition = builder.And(condition, model.FieldGoodsID().Eq(p.GoodsID))
	}
	if p.Type != enums.BOOKING_TYPE_UNKNOWN {
		condition = builder.And(condition, model.FieldType().Eq(p.Type))
	}
	if p.Status != enums.BOOKING_STATUS_UNKNOWN {
		condition = builder.And(condition, model.FieldStatus().Eq(p.Status))
	}
	if p.StartTimeBegin != datatypes.TimestampZero && p.StartTimeEnd != datatypes.TimestampZero {
		condition = builder.And(condition, model.FieldStartTime().Gte(p.StartTimeBegin))
		condition = builder.And(condition, model.FieldStartTime().Lt(p.StartTimeEnd))
	}
	if p.EndTimeBegin != datatypes.TimestampZero && p.EndTimeEnd != datatypes.TimestampZero {
		condition = builder.And(condition, model.FieldEndTime().Gte(p.EndTimeBegin))
		condition = builder.And(condition, model.FieldEndTime().Lt(p.EndTimeEnd))
	}
	return condition
}

func (p GetBookingFlowParams) Additions() []builder.Addition {
	var additions = make([]builder.Addition, 0)

	if p.Size != 0 {
		limit := builder.Limit(int64(p.Size))
		if p.Offset != 0 {
			limit = limit.Offset(int64(p.Offset))
		}
		additions = append(additions, limit)
	}

	additions = append(additions, builder.OrderBy(builder.DescOrder((&databases.Order{}).FieldCreatedAt())))

	return additions
}
