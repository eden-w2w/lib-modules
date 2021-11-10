package task_flow

import (
	"github.com/eden-framework/sqlx"
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/eden-w2w/lib-modules/constants/general_errors"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules/id_generator"
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
}

func (c *Controller) Init(db sqlx.DBExecutor) {
	c.db = db
	c.isInit = true
}

func (c Controller) CreateTaskFlow(params CreateTaskFlowParams, db sqlx.DBExecutor) (*databases.TaskFlow, error) {
	if !c.isInit {
		logrus.Panicf("[TaskFlowController] not Init")
	}

	if db == nil {
		db = c.db
	}

	id, _ := id_generator.GetGenerator().GenerateUniqueID()
	model := &databases.TaskFlow{
		FlowID:    id,
		Name:      params.Name,
		StartedAt: datatypes.MySQLTimestamp(time.Now()),
		Status:    enums.TASK_PROCESS_STATUS__CREATED,
		Type:      params.Type,
	}
	err := model.Create(db)
	if err != nil {
		logrus.Errorf("[CreateTaskFlow] model.Create err: %v, params: %+v", err, params)
	}
	return model, err
}

func (c Controller) UpdateTaskFlow(id uint64, params UpdateTaskParams) error {
	model := &databases.TaskFlow{
		FlowID: id,
	}
	err := model.FetchByFlowID(c.db)
	if err != nil {
		logrus.Errorf("[UpdateTaskFlow] model.FetchByFlowID err: %v, id: %d", err, id)
		return general_errors.InternalError
	}

	if params.StartedAt != datatypes.TimestampZero {
		model.StartedAt = params.StartedAt
	}
	if params.EndedAt != datatypes.TimestampZero {
		model.EndedAt = params.EndedAt
	}
	if params.Status != enums.TASK_PROCESS_STATUS_UNKNOWN {
		model.Status = params.Status
	}
	if params.Message != "" {
		model.Message = params.Message
	}
	err = model.UpdateByFlowIDWithStruct(c.db)
	if err != nil {
		logrus.Errorf("[UpdateTaskFlow] model.UpdateByFlowIDWithStruct err: %v, params: %+v", err, params)
		return general_errors.InternalError
	}
	return err
}

func (c Controller) GetTaskFlows(params GetTaskParams, withCount bool) (
	data []databases.TaskFlow,
	total int,
	err error,
) {
	model := &databases.TaskFlow{}
	data, err = model.List(c.db, params.Conditions(), params.Additions()...)
	if err != nil {
		logrus.Errorf("[GetTaskFlows] model.List err: %v, params: %+v", err, params)
		return nil, 0, general_errors.InternalError
	}

	if withCount {
		total, err = model.Count(c.db, params.Conditions())
		if err != nil {
			logrus.Errorf("[GetTaskFlows] model.Count err: %v, params: %+v", err, params)
			return nil, 0, general_errors.InternalError
		}
	}

	return
}
