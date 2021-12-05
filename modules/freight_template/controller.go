package freight_template

import (
	"github.com/eden-framework/sqlx"
	"github.com/eden-w2w/lib-modules/constants/general_errors"
	"github.com/eden-w2w/lib-modules/databases"
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
}

func (c *Controller) Init(db sqlx.DBExecutor) {
	c.db = db
	c.isInit = true
}

func (c Controller) GetTemplates(params GetTemplatesParams, withCount bool) (
	list []databases.FreightTemplate,
	count int,
	err error,
) {
	if !c.isInit {
		logrus.Panicf("[FreightTemplateController] not Init")
	}
	model := &databases.FreightTemplate{}
	list, err = model.List(c.db, params.Conditions(), params.Additions()...)
	if err != nil {
		logrus.Errorf("[GetTemplates] model.List err: %v, params: %+v", err, params)
		return nil, 0, general_errors.InternalError
	}
	if withCount {
		count, err = model.Count(c.db, params.Conditions())
		if err != nil {
			logrus.Errorf("[GetTemplates] model.Count err: %v, params: %+v", err, params)
			return nil, 0, general_errors.InternalError
		}
	}
	return
}

func (c Controller) GetTemplateByID(
	templateID uint64,
	db sqlx.DBExecutor,
	forUpdate bool,
) (model *databases.FreightTemplate, err error) {
	if !c.isInit {
		logrus.Panicf("[FreightTemplateController] not Init")
	}
	if db == nil {
		db = c.db
	}
	model = &databases.FreightTemplate{TemplateID: templateID}
	if forUpdate {
		err = model.FetchByTemplateIDForUpdate(db)
	} else {
		err = model.FetchByTemplateID(db)
	}
	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			return nil, general_errors.FreightTemplateNotFound
		}
		logrus.Errorf("[GetTemplateByID] model.FetchByTemplateID(db) err: %v, id: %d", err, templateID)
		return nil, general_errors.InternalError
	}
	return
}

func (c Controller) GetTemplateRuleByID(
	ruleID uint64,
	db sqlx.DBExecutor,
	forUpdate bool,
) (model *databases.FreightTemplateRules, err error) {
	if !c.isInit {
		logrus.Panicf("[FreightTemplateController] not Init")
	}
	if db == nil {
		db = c.db
	}
	model = &databases.FreightTemplateRules{RuleID: ruleID}
	if forUpdate {
		err = model.FetchByRuleIDForUpdate(db)
	} else {
		err = model.FetchByRuleID(db)
	}
	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			return nil, general_errors.FreightTemplateNotFound
		}
		logrus.Errorf("[GetTemplateRuleByID] model.FetchByRuleID(db) err: %v, id: %d", err, ruleID)
		return nil, general_errors.InternalError
	}
	return
}

func (c Controller) GetTemplateRules(
	templateID uint64,
	params GetTemplateRuleParams,
) (list []databases.FreightTemplateRules, err error) {
	model := &databases.FreightTemplateRules{}
	var conditions = model.FieldTemplateID().Eq(templateID)
	conditions = conditions.And(params.Conditions())
	list, err = model.List(c.db, conditions)
	if err != nil {
		logrus.Errorf(
			"[GetRules] model.BatchFetchByTemplateIDList err: %v, params: %+v",
			err,
			params,
		)
		return nil, general_errors.InternalError
	}
	return
}

func (c Controller) CreateTemplate(
	params CreateTemplateParams,
	db sqlx.DBExecutor,
	tx *sqlx.Tasks,
) (txr *sqlx.Tasks, template *databases.FreightTemplate, err error) {
	if !c.isInit {
		logrus.Panicf("[FreightTemplateController] not Init")
	}
	if db == nil {
		db = c.db
	}

	template, err = params.Model()
	if err != nil {
		return
	}

	var run = false
	if tx == nil {
		run = true
		tx = sqlx.NewTasks(db)
	}
	tx = tx.With(
		func(db sqlx.DBExecutor) error {
			err := template.Create(db)
			if err != nil {
				logrus.Errorf("[CreateTemplate] template.Create(db) err: %v, params: %+v", err, params)
			}
			return err
		},
	)
	if run {
		err = tx.Do()
	} else {
		txr = tx
	}
	return
}

func (c Controller) CreateTemplateRule(
	templateID uint64,
	params CreateTemplateRuleParams,
	db sqlx.DBExecutor,
	tx *sqlx.Tasks,
) (txr *sqlx.Tasks, rule *databases.FreightTemplateRules, err error) {
	if !c.isInit {
		logrus.Panicf("[FreightTemplateController] not Init")
	}
	if db == nil {
		db = c.db
	}

	rule, err = params.Model()
	if err != nil {
		return
	}
	rule.TemplateID = templateID

	var run = false
	if tx == nil {
		run = true
		tx = sqlx.NewTasks(db)
	}
	tx = tx.With(
		func(db sqlx.DBExecutor) error {
			rules, err := c.GetTemplateRules(templateID, GetTemplateRuleParams{})
			if err != nil {
				return err
			}
			for _, templateRules := range rules {
				if templateRules.Area.ContainsOne(rule.Area) {
					return general_errors.RegionsConflict
				}
			}

			err = rule.Create(db)
			if err != nil {
				logrus.Errorf("[CreateTemplateRule] rule.Create(db) err: %v, params: %+v", err, params)
			}
			return err
		},
	)
	if run {
		err = tx.Do()
	} else {
		txr = tx
	}
	return
}

func (c Controller) UpdateTemplate(
	templateID uint64,
	params UpdateTemplateParams,
	db sqlx.DBExecutor,
	tx *sqlx.Tasks,
) (txr *sqlx.Tasks, err error) {
	if !c.isInit {
		logrus.Panicf("[FreightTemplateController] not Init")
	}
	if db == nil {
		db = c.db
	}

	var run = false
	if tx == nil {
		run = true
		tx = sqlx.NewTasks(db)
	}
	tx = tx.With(
		func(db sqlx.DBExecutor) error {
			model, err := c.GetTemplateByID(templateID, db, true)
			if err != nil {
				return err
			}

			params.Fill(model)
			err = model.UpdateByIDWithStruct(db)
			if err != nil {
				logrus.Errorf(
					"[UpdateTemplate] model.UpdateByIDWithStruct(db) err: %v, templateID: %d, params: %+v",
					err,
					templateID,
					params,
				)
				return general_errors.InternalError
			}
			return nil
		},
	)
	if run {
		err = tx.Do()
	} else {
		txr = tx
	}
	return
}

func (c Controller) UpdateTemplateRule(
	ruleID uint64, params UpdateTemplateRuleParams,
	db sqlx.DBExecutor,
	tx *sqlx.Tasks,
) (txr *sqlx.Tasks, err error) {
	if !c.isInit {
		logrus.Panicf("[FreightTemplateController] not Init")
	}
	if db == nil {
		db = c.db
	}

	var run = false
	if tx == nil {
		run = true
		tx = sqlx.NewTasks(db)
	}
	tx = tx.With(
		func(db sqlx.DBExecutor) error {
			model, err := c.GetTemplateRuleByID(ruleID, db, true)
			if err != nil {
				return err
			}

			params.Fill(model)
			err = model.UpdateByRuleIDWithStruct(db)
			if err != nil {
				logrus.Errorf(
					"[UpdateTemplateRule] model.UpdateByRuleIDWithStruct(db) err: %v, ruleID: %d, params: %+v",
					err,
					ruleID,
					params,
				)
				return general_errors.InternalError
			}
			return nil
		},
	)
	if run {
		err = tx.Do()
	} else {
		txr = tx
	}
	return
}

func (c Controller) DeleteTemplate(templateID uint64, db sqlx.DBExecutor, tx *sqlx.Tasks) (txr *sqlx.Tasks, err error) {
	if !c.isInit {
		logrus.Panicf("[FreightTemplateController] not Init")
	}
	if db == nil {
		db = c.db
	}

	var run = false
	if tx == nil {
		run = true
		tx = sqlx.NewTasks(db)
	}
	tx = tx.With(
		func(db sqlx.DBExecutor) error {
			model, err := c.GetTemplateByID(templateID, db, true)
			if err != nil {
				return err
			}

			err = model.SoftDeleteByTemplateID(db)
			if err != nil {
				logrus.Errorf("[DeleteTemplate] model.DeleteByTemplateID(db) err: %v, templateID: %d", err, templateID)
				return general_errors.InternalError
			}
			return nil
		},
	)

	tx, _ = c.DeleteTemplateRulesByTemplateID(templateID, nil, tx)

	if run {
		err = tx.Do()
	} else {
		txr = tx
	}
	return
}

func (c Controller) DeleteTemplateRuleByID(ruleID uint64, db sqlx.DBExecutor, tx *sqlx.Tasks) (
	txr *sqlx.Tasks,
	err error,
) {
	if !c.isInit {
		logrus.Panicf("[FreightTemplateController] not Init")
	}
	if db == nil {
		db = c.db
	}

	var run = false
	if tx == nil {
		run = true
		tx = sqlx.NewTasks(db)
	}

	tx = tx.With(
		func(db sqlx.DBExecutor) error {
			model := &databases.FreightTemplateRules{RuleID: ruleID}
			err := model.SoftDeleteByRuleID(db)
			if err != nil {
				logrus.Errorf("[DeleteTemplateRuleByID] model.SoftDeleteByRuleID(db) err: %v, ruleID: %d", err, ruleID)
				return general_errors.InternalError
			}
			return nil
		},
	)

	if run {
		err = tx.Do()
	} else {
		txr = tx
	}
	return
}

func (c Controller) DeleteTemplateRulesByTemplateID(
	templateID uint64,
	db sqlx.DBExecutor,
	tx *sqlx.Tasks,
) (txr *sqlx.Tasks, err error) {
	if !c.isInit {
		logrus.Panicf("[FreightTemplateController] not Init")
	}
	if db == nil {
		db = c.db
	}

	var run = false
	if tx == nil {
		run = true
		tx = sqlx.NewTasks(db)
	}

	tx = tx.With(
		func(db sqlx.DBExecutor) error {
			model := &databases.FreightTemplateRules{TemplateID: templateID}
			err := model.SoftDeleteByTemplateID(db)
			if err != nil {
				logrus.Errorf(
					"[DeleteTemplateRulesByTemplateID] model.SoftDeleteByTemplateID(db) err: %v, templateID: %d",
					err,
					templateID,
				)
				return general_errors.InternalError
			}
			return nil
		},
	)

	if run {
		err = tx.Do()
	} else {
		txr = tx
	}
	return
}
