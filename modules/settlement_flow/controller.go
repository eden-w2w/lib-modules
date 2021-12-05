package settlement_flow

import (
	"fmt"
	"github.com/eden-framework/sqlx"
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/eden-w2w/lib-modules/constants/general_errors"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules"
	"github.com/eden-w2w/lib-modules/modules/id_generator"
	"github.com/eden-w2w/lib-modules/modules/promotion_flow"
	"github.com/eden-w2w/lib-modules/modules/task_flow"
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
	config SettlementConfig
}

type settlementPromotionMapping struct {
	params  *CreateSettlementParams
	flowIDs *[]uint64
}

func (c *Controller) Init(db sqlx.DBExecutor, config SettlementConfig) {
	c.db = db
	c.config = config
	c.isInit = true
}

func (c *Controller) TaskSettlement() {
	_, week := time.Now().ISOWeek()
	_ = c.RunTaskSettlement(fmt.Sprintf("第%d周", week))
}

func (c *Controller) RunTaskSettlement(settlementName string) (err error) {
	if !c.isInit {
		logrus.Panicf("[SettlementFlowController] not Init")
	}

	logrus.Infof("[TaskSettlement] start settlement for %s", settlementName)
	task, _ := task_flow.GetController().CreateTaskFlow(
		task_flow.CreateTaskFlowParams{
			Name: settlementName,
			Type: enums.TASK_TYPE__SETTLEMENT,
		}, nil,
	)
	defer func() {
		if task != nil {
			if err != nil {
				_ = task_flow.GetController().UpdateTaskFlow(
					task.FlowID, task_flow.UpdateTaskParams{
						EndedAt: datatypes.MySQLTimestamp(time.Now()),
						Status:  enums.TASK_PROCESS_STATUS__FAIL,
						Message: err.Error(),
					},
				)
			} else {
				_ = task_flow.GetController().UpdateTaskFlow(
					task.FlowID, task_flow.UpdateTaskParams{
						EndedAt: datatypes.MySQLTimestamp(time.Now()),
						Status:  enums.TASK_PROCESS_STATUS__COMPLETE,
					},
				)
			}
		}

		logrus.Infof("[TaskSettlement] complete settlement for %s", settlementName)
	}()

	list, _, err := promotion_flow.GetController().GetPromotionFlows(
		promotion_flow.GetPromotionFlowParams{
			IsNotSettlement: datatypes.BOOL_TRUE,
			CreateLt:        datatypes.MySQLTimestamp(time.Now().Add(-c.config.SettlementDuration)),
			Pagination: modules.Pagination{
				Size: -1,
			},
		}, false,
	)

	if err != nil {
		logrus.Errorf("[TaskSettlement] promotion_flow.GetController().GetPromotionFlows err: %v", err)
		return
	}

	var settlements = make(map[uint64]settlementPromotionMapping)
	for _, flow := range list {
		if _, ok := settlements[flow.UserID]; !ok {
			settlements[flow.UserID] = settlementPromotionMapping{
				params: &CreateSettlementParams{
					UserID:     flow.UserID,
					NickName:   flow.UserNickName,
					OpenID:     flow.UserOpenID,
					Name:       settlementName,
					TotalSales: 0,
					Proportion: 0,
					Amount:     0,
				},
				flowIDs: &[]uint64{},
			}
		}
		settlements[flow.UserID].params.TotalSales += flow.Amount
		settlements[flow.UserID].params.Proportion = c.config.GetProportion(settlements[flow.UserID].params.TotalSales)
		settlements[flow.UserID].params.Amount = uint64(
			float64(settlements[flow.UserID].params.TotalSales) * settlements[flow.UserID].params.Proportion,
		)
		*settlements[flow.UserID].flowIDs = append(*settlements[flow.UserID].flowIDs, flow.FlowID)
	}

	tx := sqlx.NewTasks(c.db)
	tx = tx.With(
		func(db sqlx.DBExecutor) error {
			for _, m := range settlements {
				s, err := c.CreateSettlement(*m.params, db)
				if err != nil {
					return err
				}

				for _, flowID := range *m.flowIDs {
					err = promotion_flow.GetController().UpdatePromotionSettlements(flowID, s.SettlementID, db)
					if err != nil {
						return err
					}
				}
			}
			return nil
		},
	)

	err = tx.Do()
	return
}

func (c Controller) GetSettlementFlows(
	params GetSettlementFlowsParams,
	withCount bool,
) (data []databases.SettlementFlow, count int, err error) {
	if !c.isInit {
		logrus.Panicf("[SettlementFlowController] not Init")
	}
	model := &databases.SettlementFlow{}
	data, err = model.List(c.db, params.Conditions(), params.Additions()...)
	if err != nil {
		logrus.Errorf("[GetSettlementFlows] model.List err: %v, params: %+v", err, params)
		return nil, 0, general_errors.InternalError
	}
	if withCount {
		count, err = model.Count(c.db, params.Conditions())
		if err != nil {
			logrus.Errorf("[GetSettlementFlows] model.Count err: %v, params: %+v", err, params)
			return nil, 0, general_errors.InternalError
		}
	}
	return
}

func (c Controller) GetSettlementByID(
	settlementID uint64,
	db sqlx.DBExecutor,
	forUpdate bool,
) (model *databases.SettlementFlow, err error) {
	if !c.isInit {
		logrus.Panicf("[SettlementFlowController] not Init")
	}
	if db == nil {
		db = c.db
	}
	model = &databases.SettlementFlow{SettlementID: settlementID}
	if forUpdate {
		err = model.FetchBySettlementIDForUpdate(db)
	} else {
		err = model.FetchBySettlementID(db)
	}
	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			return nil, general_errors.SettlementFlowNotFound
		}
		logrus.Errorf(
			"[GetSettlementByID] model.FetchBySettlementID err: %v, settlementID: %d, forUpdate: %v",
			err,
			settlementID,
			forUpdate,
		)
		return nil, general_errors.InternalError
	}
	return
}

func (c Controller) GetPromotionSettlementAmount(flows []databases.PromotionFlow) (totalSales, expectedAmount uint64) {
	if !c.isInit {
		logrus.Panicf("[SettlementFlowController] not Init")
	}
	for _, flow := range flows {
		totalSales += flow.Amount
	}
	for _, r := range c.config.SettlementRules {
		if totalSales >= r.MinSales && totalSales < r.MaxSales {
			expectedAmount = uint64(float64(totalSales) * r.Proportion)
			return
		}
	}
	return
}

func (c Controller) CreateSettlement(params CreateSettlementParams, db sqlx.DBExecutor) (
	*databases.SettlementFlow,
	error,
) {
	if !c.isInit {
		logrus.Panicf("[SettlementFlowController] not Init")
	}
	if db == nil {
		db = c.db
	}
	id := id_generator.GetGenerator().GenerateUniqueID()
	model := &databases.SettlementFlow{
		SettlementID: id,
		UserID:       params.UserID,
		NickName:     params.NickName,
		OpenID:       params.OpenID,
		Name:         params.Name,
		TotalSales:   params.TotalSales,
		Proportion:   params.Proportion,
		Amount:       params.Amount,
		Status:       enums.SETTLEMENT_STATUS__CREATED,
	}
	err := model.Create(db)
	if err != nil {
		logrus.Errorf("[CreateSettlement] model.Create err: %v, params: %+v", err, params)
		return nil, err
	}
	return model, nil
}

func (c Controller) UpdateSettlement(
	model *databases.SettlementFlow,
	params UpdateSettlementParams,
	db sqlx.DBExecutor,
) error {
	if !c.isInit {
		logrus.Panicf("[SettlementFlowController] not Init")
	}
	if db == nil {
		db = c.db
	}

	if params.TotalSales > 0 && model.TotalSales != params.TotalSales {
		model.TotalSales = params.TotalSales
	}
	if params.Proportion > 0 && model.Proportion != params.Proportion {
		model.Proportion = params.Proportion
	}
	if params.Amount > 0 && model.Amount != params.Amount {
		model.Amount = params.Amount
	}
	if params.Status != enums.SETTLEMENT_STATUS_UNKNOWN && model.Status != params.Status {
		model.Status = params.Status
	}
	err := model.UpdateBySettlementIDWithStruct(db)
	if err != nil {
		logrus.Errorf("[UpdateSettlement] model.UpdateBySettlementIDWithStruct err: %v, params: %+v", err, params)
		return general_errors.InternalError
	}
	return nil
}
