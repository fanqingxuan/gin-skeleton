package svc

import (
	"context"
	"fmt"
	"gin-skeleton/config"
	"time"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/gorm/utils"
)

func NewDB(c config.Config, log *Log) *gorm.DB {
	DSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", c.DB.Username, c.DB.Password, c.DB.Host, c.DB.Port, c.DB.DbName, c.DB.Charset)
	DB, err := gorm.Open(mysql.New(mysql.Config{
		DSN: DSN,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.DB.TablePrefix,
			SingularTable: c.DB.SingularTable,
		},
		Logger: newDBLogger(c, log),
	})
	if err != nil {
		panic(err)
	}
	return DB
}

func newDBLogger(c config.Config, log *Log) logger.Interface {
	var (
		traceStr     = "[%.3fms] [rows:%v] %s"
		traceWarnStr = "%s[%.3fms] [rows:%v] %s"
	)
	return &gormLogger{
		SlowThreshold: time.Duration(c.DB.SlowThreshold * int64(time.Second)),
		traceStr:      traceStr,
		traceWarnStr:  traceWarnStr,
		logger:        log,
	}
}

type gormLogger struct {
	SlowThreshold            time.Duration
	infoStr, warnStr, errStr string
	traceStr, traceWarnStr   string
	logger                   *Log
	logger.Interface
}

// LogMode log mode
func (l *gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newlogger := *l
	return &newlogger
}

// Info print info
func (l *gormLogger) Info(ctx context.Context, msg string, data ...interface{}) {

}

// Warn print warn messages
func (l *gormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {

}

// Error print error messages
func (l *gormLogger) Error(ctx context.Context, msg string, data ...interface{}) {

}

// Trace print sql message
func (l *gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	switch {
	case err != nil && !errors.Is(err, gorm.ErrRecordNotFound):
		sql, _ := fc()
		message := fmt.Sprintf("database error:%s in file %s, Executed SQL:%s ", err, trimmedPath(utils.FileWithLineNum()), sql)
		panic(message)
	case elapsed > time.Duration(l.SlowThreshold)*time.Second && l.SlowThreshold != 0:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", time.Duration(l.SlowThreshold*time.Second).Seconds())
		message := ""
		if rows == -1 {
			message = fmt.Sprintf(l.traceWarnStr, slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			message = fmt.Sprintf(l.traceWarnStr, slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
		l.logger.WithContext(ctx).WithCaller(5).Warn("slow sql", message)
	default:
		sql, rows := fc()
		message := ""
		if rows == -1 {
			message = fmt.Sprintf(l.traceStr, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			message = fmt.Sprintf(l.traceStr, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
		l.logger.WithContext(ctx).WithCaller(5).Debug("sql", message)
	}
}
