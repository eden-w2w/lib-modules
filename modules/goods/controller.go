package goods

import (
	"github.com/eden-framework/sqlx"
	"github.com/eden-framework/sqlx/builder"
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/sirupsen/logrus"
	"sync"

	"github.com/eden-w2w/lib-modules/constants/general_errors"
	"github.com/eden-w2w/lib-modules/databases"
)

var controller *Controller

func GetController() *Controller {
	if controller == nil {
		controller = &Controller{isInit: false}
	}
	return controller
}

// Controller 商品控制器，兼顾库存管理的能力
type Controller struct {
	isInit   bool
	db       sqlx.DBExecutor
	managers map[uint64]sync.Mutex
}

func (c *Controller) Init(db sqlx.DBExecutor) {
	goods := databases.Goods{}
	goodsList, err := goods.List(db, nil)
	if err != nil {
		logrus.Panicf("goods.newController goods.List(db, nil) err: %v", err)
	}
	managers := make(map[uint64]sync.Mutex)
	for _, g := range goodsList {
		managers[g.GoodsID] = sync.Mutex{}
	}
	c.db = db
	c.managers = managers
	c.isInit = true
}

func (c Controller) GetGoods(p GetGoodsParams) ([]databases.Goods, error) {
	if !c.isInit {
		logrus.Panicf("[GoodsController] not Init")
	}

	m := databases.Goods{}
	goods, err := m.List(c.db, p.Conditions(c.db), p.Additions()...)
	if err != nil {
		logrus.Errorf("[GetGoods] m.List err: %v, params: %+v", err, p)
		return nil, general_errors.InternalError
	}
	return goods, nil
}

func (c Controller) GetGoodsByID(goodsID uint64, db sqlx.DBExecutor, forUpdate bool) (*databases.Goods, error) {
	if !c.isInit {
		logrus.Panicf("[GoodsController] not Init")
	}
	if db == nil {
		db = c.db
	}

	var err error
	m := &databases.Goods{GoodsID: goodsID}
	if forUpdate {
		err = m.FetchByGoodsIDForUpdate(db)
	} else {
		err = m.FetchByGoodsID(c.db)
	}
	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			return nil, general_errors.GoodsNotFound
		}
		logrus.Errorf("[GetGood] m.FetchByGoodsID err: %v, goodsID: %d", err, goodsID)
		return nil, general_errors.InternalError
	}
	return m, nil
}

func (c Controller) LockInventory(db sqlx.DBExecutor, goodsID uint64, amount uint32) error {
	if !c.isInit {
		logrus.Panicf("[GoodsController] not Init")
	}

	if locker, ok := c.managers[goodsID]; ok {
		locker.Lock()
		defer locker.Unlock()

		goods, err := c.GetGoodsByID(goodsID, db, false)
		if err != nil {
			return err
		}

		inventory := goods.Inventory - uint64(amount)
		if inventory < 0 {
			return general_errors.GoodsInventoryShortage
		}
		f := builder.FieldValues{
			"Inventory": inventory,
		}
		err = goods.UpdateByGoodsIDWithMap(db, f)
		if err != nil {
			logrus.Errorf(
				"[LockInventory] goods.UpdateByGoodsIDWithStruct(db) err: %v, goodsID: %d, fields: %+v",
				err,
				goodsID,
				f,
			)
			return general_errors.InternalError
		}

		return nil
	}

	logrus.Errorf("[LockInventory] goodsID not found, goodsID: %d", goodsID)
	return general_errors.NotFound
}

func (c Controller) UnlockInventory(db sqlx.DBExecutor, goodsID uint64, amount uint32) error {
	if !c.isInit {
		logrus.Panicf("[GoodsController] not Init")
	}

	if locker, ok := c.managers[goodsID]; ok {
		locker.Lock()
		defer locker.Unlock()

		goods := databases.Goods{GoodsID: goodsID}
		err := goods.FetchByGoodsID(db)
		if err != nil {
			logrus.Errorf("[UnlockInventory] goods.FetchByGoodsID(db) err: %v, goodsID: %d", err, goodsID)
			return general_errors.InternalError
		}

		inventory := goods.Inventory + uint64(amount)
		f := builder.FieldValues{
			"Inventory": inventory,
		}
		err = goods.UpdateByGoodsIDWithMap(db, f)
		if err != nil {
			logrus.Errorf(
				"[LockInventory] goods.UpdateByGoodsIDWithStruct(db) err: %v, goodsID: %d, fields: %+v",
				err,
				goodsID,
				f,
			)
			return general_errors.InternalError
		}

		return nil
	}

	logrus.Errorf("[UnlockInventory] goodsID not found, goodsID: %d", goodsID)
	return general_errors.NotFound
}

func (c Controller) UpdateGoods(goodsID uint64, params UpdateGoodsParams, db sqlx.DBExecutor) error {
	if !c.isInit {
		logrus.Panicf("[GoodsController] not Init")
	}
	if db == nil {
		db = c.db
	}

	var zeroFields = make([]string, 0)
	model := &databases.Goods{GoodsID: goodsID}
	if params.Name != "" {
		model.Name = params.Name
	}
	if params.Comment != "" {
		model.Comment = params.Comment
	}
	if params.DispatchAddr != "" {
		model.DispatchAddr = params.DispatchAddr
	}
	if params.Sales != 0 {
		model.Sales = params.Sales
	}
	if params.MainPicture != "" {
		model.MainPicture = params.MainPicture
	}
	if len(params.Pictures) > 0 {
		model.Pictures = params.Pictures
	}
	if len(params.Specifications) > 0 {
		model.Specifications = params.Specifications
	}
	if len(params.Activities) > 0 {
		model.Activities = params.Activities
	}
	if params.LogisticPolicy != "" {
		model.LogisticPolicy = params.LogisticPolicy
	}
	if params.Price != 0 {
		model.Price = params.Price
	}
	if params.Inventory != nil {
		model.Inventory = *params.Inventory
		zeroFields = append(zeroFields, model.FieldKeyInventory())
	}
	if params.Detail != "" {
		model.Detail = params.Detail
	}
	if params.IsAllowBooking != datatypes.BOOL_UNKNOWN {
		model.IsAllowBooking = params.IsAllowBooking
	}
	if params.EstimatedTimeArrival != datatypes.TimestampZero {
		model.EstimatedTimeArrival = params.EstimatedTimeArrival
	}
	err := model.UpdateByGoodsIDWithStruct(db, zeroFields...)
	if err != nil {
		logrus.Errorf("[UpdateGoods] model.UpdateByGoodsIDWithStruct err: %v, params: %+v", err, params)
		return general_errors.InternalError
	}
	return nil
}

func (c Controller) CreateGoods(params CreateGoodsParams) (*databases.Goods, error) {
	if !c.isInit {
		logrus.Panicf("[GoodsController] not Init")
	}

	model := &databases.Goods{
		Name:                 params.Name,
		Comment:              params.Comment,
		DispatchAddr:         params.DispatchAddr,
		Sales:                params.Sales,
		MainPicture:          params.MainPicture,
		Pictures:             params.Pictures,
		Specifications:       params.Specifications,
		Activities:           params.Activities,
		LogisticPolicy:       params.LogisticPolicy,
		Price:                params.Price,
		Detail:               params.Detail,
		IsAllowBooking:       params.IsAllowBooking,
		EstimatedTimeArrival: params.EstimatedTimeArrival,
	}
	if params.Inventory != nil {
		model.Inventory = *params.Inventory
	}
	id, err := model.MaxGoodsID(c.db, nil)
	if err != nil {
		logrus.Errorf("[CreateGoods] model.MaxGoodsID err: %v, params: %+v", err, params)
		return nil, general_errors.InternalError
	}
	if id == 0 {
		id = 10001
	} else {
		id++
	}
	model.GoodsID = id

	err = model.Create(c.db)
	if err != nil {
		logrus.Errorf("[CreateGoods] model.Create err: %v, params: %+v", err, params)
		return nil, general_errors.InternalError
	}
	return model, nil
}
