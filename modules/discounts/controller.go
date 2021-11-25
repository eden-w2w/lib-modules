package discounts

import (
	"github.com/eden-framework/sqlx"
	"github.com/eden-w2w/lib-modules/constants/general_errors"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules"
	"github.com/sirupsen/logrus"
	"sync"
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

	mutex map[uint64]*sync.RWMutex
}

func (c *Controller) Init(db sqlx.DBExecutor) {
	c.db = db
	c.isInit = true
	c.mutex = make(map[uint64]*sync.RWMutex)

	list, _, err := c.GetDiscounts(GetDiscountsParams{
		Pagination: modules.Pagination{
			Size:   -1,
			Offset: 0,
		},
	}, false)
	if err != nil {
		return
	}

	for _, discount := range list {
		c.mutex[discount.DiscountID] = &sync.RWMutex{}
	}
}

func (c Controller) Lock(id uint64) error {
	if _, ok := c.mutex[id]; !ok {
		return general_errors.InternalError
	}
	c.mutex[id].Lock()
	return nil
}

func (c Controller) Unlock(id uint64) {
	if m, ok := c.mutex[id]; ok {
		m.Unlock()
	}
}

func (c Controller) RLock(id uint64) error {
	if _, ok := c.mutex[id]; !ok {
		return general_errors.InternalError
	}
	c.mutex[id].RLock()
	return nil
}

func (c Controller) RUnlock(id uint64) {
	if m, ok := c.mutex[id]; ok {
		m.RUnlock()
	}
}

func (c Controller) GetDiscounts(params GetDiscountsParams, withCount bool) (
	list []databases.MarketingDiscount,
	total int,
	err error,
) {
	if !c.isInit {
		logrus.Panicf("[DiscountController] not Init")
	}
	model := &databases.MarketingDiscount{}
	list, err = model.List(c.db, params.Conditions(), params.Additions()...)
	if err != nil {
		logrus.Errorf("[GetDiscounts] model.List err: %v, params: %+v", err, params)
		return nil, 0, general_errors.InternalError
	}
	if withCount {
		total, err = model.Count(c.db, params.Conditions())
		if err != nil {
			logrus.Errorf("[GetDiscounts] model.Count err: %v, params: %+v", err, params)
			return nil, 0, general_errors.InternalError
		}
	}
	return
}

func (c Controller) CreateDiscount(params CreateDiscountParams, db sqlx.DBExecutor) (
	*databases.MarketingDiscount,
	error,
) {
	if !c.isInit {
		logrus.Panicf("[DiscountController] not Init")
	}
	if db == nil {
		db = c.db
	}

	model, err := params.Model()
	if err != nil {
		return nil, err
	}

	err = model.Create(db)
	if err != nil {
		logrus.Errorf("[CreateDiscount] model.Create err: %v, params: %+v", err, params)
		return nil, general_errors.InternalError
	}
	c.mutex[model.DiscountID] = &sync.RWMutex{}
	return model, nil
}

func (c Controller) GetDiscountByID(id uint64, db sqlx.DBExecutor, forUpdate bool) (
	model *databases.MarketingDiscount,
	err error,
) {
	if !c.isInit {
		logrus.Panicf("[DiscountController] not Init")
	}
	if db == nil {
		db = c.db
	}
	model = &databases.MarketingDiscount{DiscountID: id}
	if forUpdate {
		err = model.FetchByDiscountIDForUpdate(db)
	} else {
		err = model.FetchByDiscountID(db)
	}
	if err != nil {
		logrus.Errorf("[GetDiscountByID] model.FetchByDiscountID err: %v, id: %d", err, id)
		return nil, general_errors.InternalError
	}
	return
}

func (c Controller) UpdateDiscount(
	model *databases.MarketingDiscount,
	params UpdateDiscountParams,
	db sqlx.DBExecutor,
) error {
	if !c.isInit {
		logrus.Panicf("[DiscountController] not Init")
	}
	if db == nil {
		db = c.db
	}
	if model == nil {
		logrus.Errorf("[UpdateDiscount] model is nil")
		return general_errors.InternalError
	}
	params.Fill(model)
	err := model.UpdateByDiscountIDWithStruct(db)
	if err != nil {
		logrus.Errorf("[UpdateDiscount] model.UpdateByDiscountIDWithStruct err: %v, params: %+v", err, params)
		return general_errors.InternalError
	}
	return nil
}

func (c Controller) DeleteDiscount(id uint64, db sqlx.DBExecutor) error {
	if !c.isInit {
		logrus.Panicf("[DiscountController] not Init")
	}
	if db == nil {
		db = c.db
	}
	model := &databases.MarketingDiscount{
		DiscountID: id,
	}
	err := model.DeleteByDiscountID(db)
	if err != nil {
		logrus.Errorf("[DeleteDiscount] model.DeleteByDiscountID err: %v, id: %d", err, id)
		return general_errors.InternalError
	}
	return nil
}
