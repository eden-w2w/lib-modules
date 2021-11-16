package admins

import (
	"crypto"
	"crypto/sha256"
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
	isInit          bool
	db              sqlx.DBExecutor
	expiredDuration time.Duration
	salt            string
}

func (c *Controller) Init(db sqlx.DBExecutor, salt string, d time.Duration) {
	c.db = db
	c.salt = salt
	c.expiredDuration = d
	c.isInit = true
}

func (c Controller) LoginCheck(params LoginParams) (*databases.Administrators, error) {
	model := &databases.Administrators{UserName: params.UserName}
	err := model.FetchByUserName(c.db)
	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			return nil, general_errors.AdminNotFound
		}
		logrus.Errorf("[LoginCheck] model.FetchByUserName err: %v, params: %+v", err, params)
		return nil, general_errors.InternalError
	}

	if c.password(params.Password) != model.Password {
		return nil, general_errors.InvalidUserNamePassword
	}
	return model, nil
}

func (c Controller) password(password string) string {
	hasher := sha256.New()
	hash := hasher.Sum([]byte(password + c.salt))
	return hex.EncodeToString(hash)
}

func (c Controller) ResetPassword(adminID uint64, params ResetPasswordParams) error {
	model, err := c.GetAdminByID(adminID)
	if err != nil {
		return err
	}

	model.Password = c.password(params.Password)
	err = model.UpdateByAdministratorsIDWithStruct(c.db)
	if err != nil {
		logrus.Errorf("[ResetPassword] model.UpdateByAdministratorsIDWithStruct err: %v, adminID: %d", err, adminID)
		return general_errors.InternalError
	}

	_, err = c.RefreshToken(adminID)
	return err
}

func (c Controller) GetAdminByToken(token string) (*databases.Administrators, error) {
	model := &databases.Administrators{
		Token: token,
	}
	err := model.FetchByToken(c.db)
	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			return nil, general_errors.AdminTokenExpired
		}
		logrus.Errorf("[GetAdminByToken] model.FetchByToken err: %v, token: %s", err, token)
		return nil, general_errors.InternalError
	}
	return model, nil
}

func (c Controller) GetAdminByID(id uint64) (*databases.Administrators, error) {
	model := &databases.Administrators{
		AdministratorsID: id,
	}
	err := model.FetchByAdministratorsID(c.db)
	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			return nil, general_errors.AdminNotFound
		}
		logrus.Errorf("[GetAdminByID] model.FetchByAdministratorsID err: %v, id: %d", err, id)
		return nil, general_errors.InternalError
	}
	return model, nil
}

func (c Controller) RefreshToken(id uint64) (*databases.Administrators, error) {
	token := c.generateToken(id)
	model := &databases.Administrators{
		AdministratorsID: id,
		Token:            token,
		ExpiredAt:        datatypes.MySQLTimestamp(time.Now().Add(c.expiredDuration)),
	}
	err := model.UpdateByAdministratorsIDWithStruct(c.db)
	if err != nil {
		logrus.Errorf("[RefreshToken] model.UpdateByAdministratorsIDWithStruct(c.db) err: %v, adminID: %d", err, id)
		return nil, general_errors.InternalError
	}

	err = model.FetchByAdministratorsID(c.db)
	if err != nil {
		logrus.Errorf("[RefreshToken] model.FetchByAdministratorsID(c.db) err: %v, adminID: %d", err, id)
		return nil, general_errors.InternalError
	}

	return model, nil
}

func (c Controller) CreateAdmin(params LoginParams) (*databases.Administrators, error) {
	id, _ := id_generator.GetGenerator().GenerateUniqueID()
	model := &databases.Administrators{
		AdministratorsID: id,
		UserName:         params.UserName,
		Password:         c.password(params.Password),
	}
	err := model.Create(c.db)
	if err != nil {
		logrus.Errorf("[CreateAdmin] model.Create err: %v, params: %+v", err, params)
		return nil, general_errors.InternalError
	}

	model, err = c.RefreshToken(id)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (c Controller) generateToken(userID uint64) string {
	id := strconv.FormatUint(userID, 10)
	t := strconv.FormatInt(time.Now().UnixNano(), 10)
	h := crypto.SHA256.New()
	h.Write([]byte(id + t))
	hash := h.Sum(nil)
	return hex.EncodeToString(hash)
}

func (c Controller) GetAdmins() ([]databases.Administrators, error) {
	model := &databases.Administrators{}
	admins, err := model.List(c.db, nil)
	if err != nil {
		logrus.Errorf("[GetAdmins] model.List err: %v", err)
		return nil, general_errors.InternalError
	}
	return admins, nil
}
