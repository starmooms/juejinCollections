package dal

import (
	"juejinCollections/logger"

	"github.com/sirupsen/logrus"
	"xorm.io/xorm/log"
)

type DalLog struct {
	logCtx   *logrus.Logger
	showSql  bool
	logLevel log.LogLevel
}

func DalLogNew() log.Logger {
	daLog := &DalLog{
		logCtx:   logger.Logger,
		logLevel: log.LOG_WARNING,
	}
	daLog.CheckLevel()
	return daLog
}

func (l *DalLog) Debug(v ...interface{}) {
	l.logCtx.Debug(v...)
}

func (l *DalLog) Debugf(format string, v ...interface{}) {
	l.logCtx.Debugf(format, v...)
}

func (l *DalLog) Error(v ...interface{}) {
	l.logCtx.Error(v...)
}

func (l *DalLog) Errorf(format string, v ...interface{}) {
	l.logCtx.Errorf(format, v...)
}

func (l *DalLog) Info(v ...interface{}) {
	l.logCtx.Info(v...)
}

func (l *DalLog) Infof(format string, v ...interface{}) {
	l.logCtx.Infof(format, v...)
}

func (l *DalLog) Warn(v ...interface{}) {
	l.logCtx.Warn(v...)
}

func (l *DalLog) Warnf(format string, v ...interface{}) {
	l.logCtx.Warnf(format, v...)
}

func (l *DalLog) IsShowSQL() bool {
	return l.showSql
}

func (l *DalLog) ShowSQL(show ...bool) {
	if len(show) == 0 {
		l.showSql = true
	} else {
		l.showSql = show[0]
	}
}

func (l *DalLog) Level() log.LogLevel {
	return l.logLevel
}

func (l *DalLog) SetLevel(level log.LogLevel) {
	l.logCtx.Warn("%s daLog SetLevel invalid is auto set from logger", level)
}

func (l *DalLog) CheckLevel() {
	daLevel := l.logLevel
	level := l.logCtx.GetLevel()
	switch level {
	case logrus.DebugLevel, logrus.TraceLevel:
		daLevel = log.LOG_DEBUG
	case logrus.InfoLevel:
		daLevel = log.LOG_INFO
	case logrus.WarnLevel:
		daLevel = log.LOG_WARNING
	case logrus.ErrorLevel:
		daLevel = log.LOG_ERR
	}
	l.logLevel = daLevel
}
