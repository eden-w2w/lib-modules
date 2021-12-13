package settings

import (
	"github.com/eden-framework/sqlx"
	"github.com/eden-w2w/lib-modules/constants/general_errors"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules/id_generator"
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
	model  *databases.Settings
}

func (c *Controller) Init(db sqlx.DBExecutor) {
	model := &databases.Settings{}
	models, err := model.List(db, nil)
	if err != nil {
		logrus.Errorf("[SettingController] Init model.List err: %v", err)
		return
	}

	if len(models) == 0 {
		id := id_generator.GetGenerator().GenerateUniqueID()
		model.SettingsID = id
		err = model.Create(db)
		if err != nil {
			logrus.Errorf("[SettingController] Init model.Create err: %v", err)
			return
		}
	} else {
		model = &models[0]
	}
	c.model = model
	c.db = db
	c.isInit = true
}

func (c Controller) GetSetting() *databases.Settings {
	err := c.model.FetchBySettingsID(c.db)
	if err != nil {
		return nil
	}
	return c.model
}

func (c Controller) UpdateSetting(params UpdateSettingParams, db sqlx.DBExecutor) error {
	if !c.isInit || c.model == nil {
		logrus.Panicf("[SettingController] not Init")
	}

	if db == nil {
		db = c.db
	}

	params.Fill(c.model)

	err := c.model.UpdateBySettingsIDWithStruct(db)
	if err != nil {
		logrus.Errorf("[UpdateSetting] c.model.UpdateBySettingsIDWithStruct err: %v, params: %+v", err, params)
		return general_errors.InternalError
	}

	return nil
}
