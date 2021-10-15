package promotion_flow

import (
	"github.com/eden-framework/sqlx"
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

func (c Controller) CreatePromotionFlow(params CreatePromotionFlowParams, db sqlx.DBExecutor) (*databases.PromotionFlow, error) {
	if !c.isInit {
		logrus.Panicf("[PromotionFlowController] not Init")
	}
	if db == nil {
		db = c.db
	}
	id, _ := id_generator.GetGenerator().GenerateUniqueID()
	model := &databases.PromotionFlow{
		FlowID:          id,
		UserID:          params.UserID,
		UserNickName:    params.UserNickName,
		UserOpenID:      params.UserOpenID,
		RefererID:       params.RefererID,
		RefererNickName: params.RefererNickName,
		RefererOpenID:   params.RefererOpenID,
		Amount:          params.Amount,
		PaymentFlowID:   params.PaymentFlowID,
	}
	err := model.Create(db)
	if err != nil {
		logrus.Errorf("[CreatePromotionFlow] model.Create err: %v, params: %+v", err, params)
		return nil, general_errors.InternalError
	}
	return model, nil
}

func (c Controller) GetPromotionFlows(params GetPromotionFlowParams, withCount bool) (list []databases.PromotionFlow, count int, err error) {
	if !c.isInit {
		logrus.Panicf("[PromotionFlowController] not Init")
	}
	model := &databases.PromotionFlow{}
	list, err = model.List(c.db, params.Conditions(), params.Additions()...)
	if err != nil {
		logrus.Errorf("[GetPromotionFlows] model.List err: %v, params: %+v", err, params)
		return
	}

	if withCount {
		count, err = model.Count(c.db, params.Conditions())
		if err != nil {
			logrus.Errorf("[GetPromotionFlows] model.Count err: %v, params: %+v", err, params)
			return
		}
	}
	return
}

func (c Controller) UpdatePromotionSettlements(flowID, settlementID uint64, db sqlx.DBExecutor) error {
	if db == nil {
		db = c.db
	}
	model := &databases.PromotionFlow{
		FlowID:       flowID,
		SettlementID: settlementID,
	}
	err := model.UpdateByFlowIDWithStruct(db)
	if err != nil {
		logrus.Errorf("[PromotionFlowController] model.UpdateByFlowIDWithStruct err: %v, flowID: %d, settlementID: %d", err, flowID, settlementID)
		return err
	}
	return nil
}
