package user

import (
	"crypto"
	"encoding/hex"
	"github.com/eden-framework/sqlx"
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
