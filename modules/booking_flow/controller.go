package booking_flow

import (
	"github.com/eden-framework/sqlx"
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/eden-w2w/lib-modules/constants/general_errors"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules/id_generator"
	"github.com/sirupsen/logrus"
	"time"
)

var controller *Controller

func GetController() *Controller {
	if controller == nil {
		controller = &Controller{isInit: false}
	}
	return controller
}

type Controller struct {
	isInit bool
	db     sqlx.DBExecutor
}

func (c *Controller) Init(db sqlx.DBExecutor) {
	c.db = db
	c.isInit = true
}

func (c Controller) CreateBookingFlow(params CreateBookingFlowParams, db sqlx.DBExecutor) (
	*databases.BookingFlow,
	error,
) {
	if !c.isInit {
		logrus.Panicf("[BookingFlowController] not Init")
	}
	if db == nil {
		db = c.db
	}
	model := params.Model()
	id := id_generator.GetGenerator().GenerateUniqueID()
	model.FlowID = id

	err := model.Create(db)
	if err != nil {
		logrus.Errorf("[CreateBookingFlow] model.Create(db) err: %v, params: %+v", err, params)
		return nil, general_errors.InternalError
	}
	return model, nil
}

func (c Controller) UpdateBookingFlow(
	model *databases.BookingFlow,
	params UpdateBookingFlowParams,
	db sqlx.DBExecutor,
) error {
	if !c.isInit {
		logrus.Panicf("[BookingFlowController] not Init")
	}
	if db == nil {
		db = c.db
	}
	if model == nil {
		logrus.Errorf("[UpdateBookingFlow] model == nil")
		return general_errors.InternalError
	}
	zeroFields := params.Fill(model)
	err := model.UpdateByFlowIDWithStruct(db, zeroFields...)
	if err != nil {
		logrus.Errorf("[UpdateBookingFlow] model.UpdateByFlowIDWithStruct(db) err: %v, params: %+v", err, params)
		return general_errors.InternalError
	}
	return nil
}

func (c Controller) GetBookingFlows(params GetBookingFlowParams, withCount bool) (
	data []databases.BookingFlow,
	count int,
	err error,
) {
	if !c.isInit {
		logrus.Panicf("[BookingFlowController] not Init")
	}
	model := &databases.BookingFlow{}
	data, err = model.List(c.db, params.Conditions(), params.Additions()...)
	if err != nil {
		logrus.Errorf("[GetBookingFlows] model.List err: %v, params: %+v", err, params)
		return nil, 0, general_errors.InternalError
	}
	if withCount {
		count, err = model.Count(c.db, params.Conditions())
		if err != nil {
			logrus.Errorf("[GetBookingFlows] model.Count err: %v, params: %+v", err, params)
			return nil, 0, general_errors.InternalError
		}
	}

	return
}

func (c Controller) GetBookingFlowByID(flowID uint64, db sqlx.DBExecutor, forUpdate bool) (
	*databases.BookingFlow,
	error,
) {
	if !c.isInit {
		logrus.Panicf("[BookingFlowController] not Init")
	}
	if db == nil {
		db = c.db
	}
	var err error
	model := &databases.BookingFlow{FlowID: flowID}
	if forUpdate {
		err = model.FetchByFlowIDForUpdate(db)
	} else {
		err = model.FetchByFlowID(db)
	}
	if err != nil {
		logrus.Errorf("[GetBookingFlowByID] model.FetchByFlowID(db) err: %v, flowID: %d", err, flowID)
		return nil, general_errors.InternalError
	}
	return model, nil
}

func (c Controller) GetBookingFlowByGoodsID(goodsID uint64) (
	[]databases.BookingFlow,
	error,
) {
	if !c.isInit {
		logrus.Panicf("[BookingFlowController] not Init")
	}
	currentTime := time.Now()
	list, _, err := c.GetBookingFlows(
		GetBookingFlowParams{
			GoodsID:        goodsID,
			Status:         enums.BOOKING_STATUS__PROCESS,
			StartTimeBegin: datatypes.MySQLTimestamp(currentTime),
		}, false,
	)
	if err != nil {
		logrus.Errorf(
			"[GetBookingFlowByGoodsID] c.GetBookingFlows err: %v, goodsID: %d",
			err,
			goodsID,
		)
		return nil, general_errors.InternalError
	}
	for i := 0; i < len(list); i++ {
		if !list[i].EndTime.IsZero() {
			if currentTime.After(time.Time(list[i].EndTime)) {
				list = append(list[:i], list[i+1:]...)
				i--
			}
		}
	}
	return list, nil
}
