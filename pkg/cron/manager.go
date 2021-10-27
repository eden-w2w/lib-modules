package cron

import (
	"context"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

var manager *Manager

type Manager struct {
	t *cron.Cron
}

func GetManager() *Manager {
	if manager == nil {
		parser := cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
		t := cron.New(cron.WithParser(parser))
		manager = &Manager{t: t}
	}
	return manager
}

func (m *Manager) AddFunc(spec string, cmd func()) (cron.EntryID, error) {
	logrus.Infof("[CronManager] add new task %s", spec)
	return m.t.AddFunc(spec, cmd)
}

func (m Manager) Start() {
	m.t.Start()
}

func (m Manager) Stop() context.Context {
	return m.t.Stop()
}
