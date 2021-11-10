package refund_flow

import (
	"github.com/eden-framework/sqlx"
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

	id, _ := id_generator.GetGenerator().GenerateUniqueID()
	model := &databases.RefundFlow{
		FlowID:              id,
		RemoteFlowID:        params.RemoteFlowID,
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
