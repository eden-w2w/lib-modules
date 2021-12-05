package refund_flow

import (
	"github.com/eden-framework/sqlx"
	"github.com/eden-framework/sqlx/builder"
	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/eden-w2w/lib-modules/constants/general_errors"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules/id_generator"
	"github.com/sirupsen/logrus"
)

var controller *Controller

type Controller struct {
	isInit bool
	db     sqlx.DBExecutor
}

func (c *Controller) Init(db sqlx.DBExecutor) {
	c.db = db
	c.isInit = true
}

func GetController() *Controller {
	if controller == nil {
		controller = &Controller{isInit: false}
	}
	return controller
}

func (c Controller) CreateRefundFlow(params CreateRefundFlowRequest, db sqlx.DBExecutor) (
	*databases.RefundFlow,
	error,
) {
	if !c.isInit {
		logrus.Panicf("[RefundFlowController] not Init")
	}
	if db == nil {
		db = c.db
	}

	id := id_generator.GetGenerator().GenerateUniqueID()
	model := &databases.RefundFlow{
		FlowID:              id,
		PaymentFlowID:       params.PaymentFlowID,
		RemotePaymentFlowID: params.RemotePaymentFlowID,
		Channel:             enums.REFUND_CHANNEL_UNKNOWN,
		Account:             "",
		Status:              enums.REFUND_STATUS__PROCESSING,
		TotalAmount:         params.TotalAmount,
		RefundAmount:        params.RefundAmount,
	}
	err := model.Create(db)
	if err != nil {
		logrus.Errorf("[CreateRefundFlow] model.Create err: %v, params: %+v", err, params)
		return nil, general_errors.InternalError
	}
	return model, nil
}

func (c Controller) UpdateRefundFlowRemoteID(flowID uint64, remoteID string, db sqlx.DBExecutor) error {
	if !c.isInit {
		logrus.Panicf("[PaymentFlowController] not Init")
	}
	if db == nil {
		db = c.db
	}
	model := &databases.RefundFlow{FlowID: flowID}
	fields := builder.FieldValues{
		model.FieldKeyRemoteFlowID(): remoteID,
	}
	err := model.UpdateByFlowIDWithMap(db, fields)
	if err != nil {
		logrus.Errorf(
			"[UpdateRefundFlowRemoteID] model.UpdateByFlowIDWithMap err: %v, flowID: %d, remoteID: %s",
			err,
			flowID,
			remoteID,
		)
		return general_errors.InternalError
	}
	return nil
}

func (c Controller) GetRefundFlows(params GetRefundFlowsRequest, withCount bool) (data []databases.RefundFlow, total int, err error) {
	if !c.isInit {
		logrus.Panicf("[PaymentFlowController] not Init")
	}
	model := databases.RefundFlow{}
	data, err = model.List(c.db, params.Conditions(), params.Additions()...)
	if err != nil {
		logrus.Errorf("[GetRefundFlows] model.List err: %v, params: %+v", err, params)
		err = general_errors.InternalError
		return
	}

	if withCount {
		total, err = model.Count(c.db, params.Conditions())
		if err != nil {
			logrus.Errorf("[GetRefundFlows] model.Count err: %v, params: %+v", err, params)
			err = general_errors.InternalError
			return
		}
	}
	return
}

func (c Controller) UpdateRefundFlow(flowID uint64, params UpdateRefundFlowRequest, db sqlx.DBExecutor) error {
	if !c.isInit {
		logrus.Panicf("[PaymentFlowController] not Init")
	}
	if db == nil {
		db = c.db
	}
	model := &databases.RefundFlow{FlowID: flowID}
	fields := builder.FieldValues{}
	if params.Status != enums.REFUND_STATUS_UNKNOWN {
		fields[model.FieldKeyStatus()] = params.Status
	}
	if params.Account != "" {
		fields[model.FieldKeyAccount()] = params.Account
	}
	if params.RefundAmount != 0 {
		fields[model.FieldKeyRefundAmount()] = params.RefundAmount
	}
	if params.TotalAmount != 0 {
		fields[model.FieldKeyTotalAmount()] = params.TotalAmount
	}
	if params.Channel != enums.REFUND_CHANNEL_UNKNOWN {
		fields[model.FieldKeyChannel()] = params.Channel
	}
	err := model.UpdateByFlowIDWithMap(db, fields)
	if err != nil {
		logrus.Errorf(
			"[UpdateRefundFlowRemoteID] model.UpdateByFlowIDWithMap err: %v, flowID: %d, params: %+v",
			err,
			flowID,
			params,
		)
		return general_errors.InternalError
	}
	return nil
}

func (c Controller) GetRefundFlowByFlowID(flowID uint64, db sqlx.DBExecutor, forUpdate bool) (
	model *databases.RefundFlow,
	err error,
) {
	if !c.isInit {
		logrus.Panicf("[PaymentFlowController] not Init")
	}
	if db == nil {
		db = c.db
	}
	model = &databases.RefundFlow{FlowID: flowID}
	if forUpdate {
		err = model.FetchByFlowIDForUpdate(db)
	} else {
		err = model.FetchByFlowID(db)
	}
	if err != nil {
		logrus.Errorf("[GetRefundFlowByFlowID] model.FetchByFlowID err: %v, flowID: %d", err, flowID)
		return nil, general_errors.InternalError
	}
	return model, nil
}
