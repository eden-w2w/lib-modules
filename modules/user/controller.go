package user

import (
	"crypto"
	"encoding/hex"
	"github.com/eden-framework/sqlx"
	"github.com/eden-framework/sqlx/builder"
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/lib-modules/constants/general_errors"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules/id_generator"
	"github.com/sirupsen/logrus"
	"strconv"
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

func (c Controller) GetUsers(params GetUsersParams, withCount bool) (result []databases.User, count int, err error) {
	model := &databases.User{}
	result, err = model.List(c.db, params.Conditions(), params.Additions()...)
	if err != nil {
		logrus.Errorf("[GetUsers] model.List err: %v, params: %+v", err, params)
		return nil, 0, general_errors.InternalError
	}
	if withCount {
		count, err = model.Count(c.db, params.Conditions())
		if err != nil {
			logrus.Errorf("[GetUsers] model.Count err: %v, params: %+v", err, params)
			return nil, 0, general_errors.InternalError
		}
	}
	return
}

func (c Controller) CreateUserByWechatSession(params CreateUserByWechatSessionParams) (*databases.User, error) {
	if !c.isInit {
		logrus.Panicf("[UserController] not Init")
	}
	id, _ := id_generator.GetGenerator().GenerateUniqueID()
	model := &databases.User{
		UserID:      id,
		Token:       c.generateToken(id),
		OpenID:      params.OpenID,
		UnionID:     params.UnionID,
		SessionKey:  params.SessionKey,
		OperateTime: datatypes.OperateTime{},
	}
	err := model.Create(c.db)
	if err != nil {
		logrus.Errorf("[CreateUserByWechatSession] model.Create(c.db) err: %v, params: %+v", err, params)
		return nil, general_errors.InternalError
	}

	return model, nil
}

func (c Controller) RefreshToken(userID uint64) (*databases.User, error) {
	if !c.isInit {
		logrus.Panicf("[UserController] not Init")
	}
	token := c.generateToken(userID)
	model := &databases.User{
		UserID: userID,
		Token:  token,
	}
	err := model.UpdateByUserIDWithStruct(c.db)
	if err != nil {
		logrus.Errorf("[RefreshToken] model.UpdateByUserIDWithStruct(c.db) err: %v, userID: %d", err, userID)
		return nil, general_errors.InternalError
	}

	err = model.FetchByUserID(c.db)
	if err != nil {
		logrus.Errorf("[RefreshToken] model.UpdateByUserIDWithStruct(c.db) err: %v, userID: %d", err, userID)
		return nil, general_errors.InternalError
	}

	return model, nil
}

func (c Controller) UpdateUserInfo(userID uint64, params UpdateUserInfoParams) error {
	if !c.isInit {
		logrus.Panicf("[UserController] not Init")
	}
	model := databases.User{
		UserID: userID,
	}
	err := model.FetchByUserID(c.db)
	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			return general_errors.UserNotFound
		}
		logrus.Errorf("[UpdateUserInfo] model.FetchByUserID(c.db) err: %v, userID: %d, params: %+v", err, userID, params)
		return err
	}

	if params.Diff(&model) {
		err = model.UpdateByUserIDWithStruct(c.db)
		if err != nil {
			logrus.Errorf("[UpdateUserInfo] model.UpdateByUserIDWithStruct(c.db) err: %v, userID: %d, params: %+v", err, userID, params)
			return general_errors.InternalError
		}
	}

	return nil
}

func (c Controller) generateToken(userID uint64) string {
	id := strconv.FormatUint(userID, 10)
	t := strconv.FormatInt(time.Now().UnixNano(), 10)
	sha256 := crypto.SHA256.New()
	sha256.Write([]byte(id + t))
	hash := sha256.Sum(nil)
	return hex.EncodeToString(hash)
}

func (c Controller) GetUserByUserID(userID uint64, db sqlx.DBExecutor, forUpdate bool) (model *databases.User, err error) {
	if !c.isInit {
		logrus.Panicf("[UserController] not Init")
	}
	if db == nil {
		db = c.db
	}
	model = &databases.User{
		UserID: userID,
	}
	if forUpdate {
		err = model.FetchByUserIDForUpdate(db)
	} else {
		err = model.FetchByUserID(db)
	}
	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			return nil, general_errors.UserNotFound
		}
		logrus.Errorf("[GetUserByUserID] model.FetchByUserID err: %v, userID: %d", err, userID)
		return nil, general_errors.InternalError
	}
	return model, nil
}

func (c Controller) GetUserByOpenID(openID string) (*databases.User, error) {
	if !c.isInit {
		logrus.Panicf("[UserController] not Init")
	}
	model := &databases.User{
		OpenID: openID,
	}
	err := model.FetchByOpenID(c.db)
	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			return nil, general_errors.UserNotFound
		}
		logrus.Errorf("[GetUserByOpenID] model.FetchByOpenID err: %v, openID: %s", err, openID)
		return nil, general_errors.InternalError
	}
	return model, nil
}

func (c Controller) GetUserByToken(token string) (*databases.User, error) {
	if !c.isInit {
		logrus.Panicf("[UserController] not Init")
	}
	model := &databases.User{
		Token: token,
	}
	err := model.FetchByToken(c.db)
	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			return nil, general_errors.UserNotFound
		}
		logrus.Errorf("[GetUserByToken] model.FetchByToken err: %v, token: %s", err, token)
		return nil, general_errors.InternalError
	}
	return model, nil
}

func (c Controller) GetUserByNameOrOpenID(params GetUserByNameOrOpenIDParams) ([]databases.User, error) {
	if !c.isInit {
		logrus.Panicf("[UserController] not Init")
	}
	model := databases.User{}
	list, err := model.List(c.db, params.Conditions(), builder.Limit(10))
	if err != nil {
		logrus.Errorf("[GetUserByNameOrOpenID] model.List err: %v, keywords: %s", err, params.Keywords)
		return nil, general_errors.InternalError
	}
	return list, nil
}

func (c Controller) GetShippingAddressByUserID(userID uint64) ([]databases.ShippingAddress, error) {
	if !c.isInit {
		logrus.Panicf("[UserController] not Init")
	}
	model := &databases.ShippingAddress{}
	data, err := model.BatchFetchByUserIDList(c.db, []uint64{userID})
	if err != nil {
		logrus.Errorf("[GetShippingAddressByUserID] model.BatchFetchByUserIDList err: %v, userID: %d", err, userID)
		return nil, general_errors.InternalError
	}
	return data, nil
}

func (c Controller) CreateShippingAddress(params CreateShippingAddressParams, db sqlx.DBExecutor) (*databases.ShippingAddress, error) {
	if !c.isInit {
		logrus.Panicf("[UserController] not Init")
	}
	if db == nil {
		db = c.db
	}
	id, _ := id_generator.GetGenerator().GenerateUniqueID()
	model := &databases.ShippingAddress{
		ShippingID: id,
		UserID:     params.UserID,
		Recipients: params.Recipients,
		District:   params.District,
		Address:    params.Address,
		Mobile:     params.Mobile,
	}
	err := model.Create(db)
	if err != nil {
		logrus.Errorf("[CreateShippingAddress] model.Create err: %v, params: %+v", err, params)
		return nil, general_errors.InternalError
	}
	return model, nil
}

func (c Controller) UpdateShippingAddress(params UpdateShippingAddressParams, userID uint64, db sqlx.DBExecutor) error {
	if !c.isInit {
		logrus.Panicf("[UserController] not Init")
	}
	if db == nil {
		db = c.db
	}
	model := &databases.ShippingAddress{ShippingID: params.ShippingID}
	err := model.FetchByShippingID(db)
	if err != nil {
		logrus.Errorf("[UpdateShippingAddress] model.FetchByShippingID err: %v, params: %+v", err, params)
		return general_errors.InternalError
	}
	if userID != 0 && userID != model.UserID {
		logrus.Errorf("[UpdateShippingAddress] userID != 0 && userID != model.UserID, params: %+v", params)
		return general_errors.Forbidden
	}

	model.Recipients = params.Recipients
	model.Mobile = params.Mobile
	model.District = params.District
	model.Address = params.Address
	err = model.UpdateByShippingIDWithStruct(db)
	if err != nil {
		logrus.Errorf("[UpdateShippingAddress] model.UpdateByShippingIDWithStruct err: %v, userID: %d, params: %+v", err, userID, params)
		return general_errors.InternalError
	}
	return nil
}

func (c Controller) DeleteShippingAddress(shippingID uint64, userID uint64, db sqlx.DBExecutor) error {
	if !c.isInit {
		logrus.Panicf("[UserController] not Init")
	}
	if db == nil {
		db = c.db
	}
	model := &databases.ShippingAddress{ShippingID: shippingID}
	err := model.FetchByShippingID(db)
	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			return general_errors.NotFound.StatusError().WithMsg("未找到收货地址信息")
		}
		logrus.Errorf("[UpdateShippingAddress] model.FetchByShippingID err: %v, id: %d", err, shippingID)
		return general_errors.InternalError
	}
	if userID != 0 && userID != model.UserID {
		logrus.Errorf("[UpdateShippingAddress] userID != 0 && userID != model.UserID, shippingID: %d, userID: %d", shippingID, userID)
		return general_errors.Forbidden
	}

	err = model.SoftDeleteByShippingID(db)
	if err != nil {
		logrus.Errorf("[UpdateShippingAddress] model.SoftDeleteByShippingID err: %v, id: %d", err, shippingID)
		return general_errors.InternalError
	}
	return nil
}

func (c Controller) SetDefaultShippingAddress(userID, shippingID uint64, db sqlx.DBExecutor) error {
	if !c.isInit {
		logrus.Panicf("[UserController] not Init")
	}
	if db == nil {
		db = c.db
	}
	model := &databases.ShippingAddress{ShippingID: shippingID}
	err := model.FetchByShippingID(db)
	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			return general_errors.NotFound.StatusError().WithMsg("未找到收货地址信息")
		}
		logrus.Errorf("[UpdateShippingAddress] model.FetchByShippingID err: %v, id: %d", err, shippingID)
		return general_errors.InternalError
	}
	if userID != model.UserID {
		logrus.Errorf("[UpdateShippingAddress] userID != 0 && userID != model.UserID, shippingID: %d, userID: %d", shippingID, userID)
		return general_errors.Forbidden
	}

	err = model.ResetAllDefault(db)
	if err != nil {
		logrus.Errorf("[UpdateShippingAddress] model.ResetAllDefault err: %v, id: %d", err, shippingID)
		return general_errors.InternalError
	}

	model.Default = datatypes.BOOL_TRUE
	err = model.UpdateByShippingIDWithStruct(db)
	if err != nil {
		logrus.Errorf("[UpdateShippingAddress] model.UpdateByShippingIDWithStruct err: %v, id: %d", err, shippingID)
		return general_errors.InternalError
	}
	return nil
}
