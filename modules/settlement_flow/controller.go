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
	"github.com/robfig/cron/v3"
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
	config *SettlementConfig
	task   *cron.Cron
}

type settlementPromotionMapping struct {
	params  *CreateSettlementParams
	flowIDs *[]uint64
}

func (c *Controller) Init(db sqlx.DBExecutor, config *SettlementConfig) {
	c.db = db

	if config != nil {
		parser := cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
		t := cron.New(cron.WithParser(parser))
		c.config = config
		c.task = t
		_, err := t.AddFunc(config.ToSettlementCronRule(), controller.TaskSettlement)
		if err != nil {
			logrus.Panicf("[settlement_flow.newController] t.AddFunc err: %v, rules: %s", err, config.ToSettlementCronRule())
		}
	}
	c.isInit = true
}

func (c Controller) TaskSettlement() {
	_, week := time.Now().ISOWeek()
	_ = c.RunTaskSettlement(fmt.Sprintf("%d周", week))
}

func (c Controller) RunTaskSettlement(settlementName string) error {
	logrus.Infof("[TaskSettlement] start settlement for %s", settlementName)
	defer logrus.Infof("[TaskSettlement] complete settlement for %s", settlementName)

	list, err := promotion_flow.GetController().GetPromotionFlows(promotion_flow.GetPromotionFlowParams{
		IsNotSettlement: true,
		CreateLt:        datatypes.MySQLTimestamp(time.Now().Add(-c.config.SettlementDuration)),
		Pagination: modules.Pagination{
			Size: -1,
		},
	})

	if err != nil {
		logrus.Errorf("[TaskSettlement] promotion_flow.GetController().GetPromotionFlows err: %v", err)
		return err
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
			float64(settlements[flow.UserID].params.TotalSales) * settlements[flow.UserID].params.Proportion)
		*settlements[flow.UserID].flowIDs = append(*settlements[flow.UserID].flowIDs, flow.FlowID)
	}

	tx := sqlx.NewTasks(c.db)
	tx = tx.With(func(db sqlx.DBExecutor) error {
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
	})

	return tx.Do()
}

func (c Controller) StartTask() {
	if !c.isInit {
		logrus.Panicf("[SettlementFlowController] not Init")
	}
	if c.task != nil {
		c.task.Start()
	}
}

func (c Controller) StopTask() {
	if !c.isInit {
		logrus.Panicf("[SettlementFlowController] not Init")
	}
	if c.task != nil {
		c.task.Stop()
	}
}

func (c Controller) GetSettlementByUserIDAndName(userID uint64, name string, db sqlx.DBExecutor, forUpdate bool) (model *databases.SettlementFlow, err error) {
	if !c.isInit {
		logrus.Panicf("[SettlementFlowController] not Init")
	}
	if db == nil {
		db = c.db
	}
	model = &databases.SettlementFlow{
		UserID: userID,
		Name:   name,
	}

	if forUpdate {
		err = model.FetchByUserIDAndNameForUpdate(db)
	} else {
		err = model.FetchByUserIDAndName(db)
	}
	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			return nil, general_errors.NotFound
		}
		logrus.Errorf("[GetSettlementByUserIDAndName] err: %v, userID: %d, name: %s", err, userID, name)
		return nil, err
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

func (c Controller) CreateSettlement(params CreateSettlementParams, db sqlx.DBExecutor) (*databases.SettlementFlow, error) {
	if !c.isInit {
		logrus.Panicf("[SettlementFlowController] not Init")
	}
	if db == nil {
		db = c.db
	}
	id, _ := id_generator.GetGenerator().GenerateUniqueID()
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
