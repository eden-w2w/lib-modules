package settlement_flow

import (
	"github.com/eden-framework/sqlx"
	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/eden-w2w/lib-modules/constants/general_errors"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules/id_generator"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
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
	task   *cron.Cron
}

func (c *Controller) Init(db sqlx.DBExecutor, config SettlementConfig) {
	parser := cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
	t := cron.New(cron.WithParser(parser))
	c.db = db
	c.config = config
	c.task = t
	_, err := t.AddFunc(config.ToSettlementCronRule(), controller.TaskSettlement)
	if err != nil {
		logrus.Panicf("[settlement_flow.newController] t.AddFunc err: %v, rules: %s", err, config.ToSettlementCronRule())
	}
	c.isInit = true
}

func (c Controller) TaskSettlement() {

}

func (c Controller) StartTask() {
	if !c.isInit {
		logrus.Panicf("[SettlementFlowController] not Init")
	}
	c.task.Start()
}

func (c Controller) StopTask() {
	if !c.isInit {
		logrus.Panicf("[SettlementFlowController] not Init")
	}
	c.task.Stop()
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
