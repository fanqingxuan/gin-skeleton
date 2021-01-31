package core

import (
	"context"
	"fmt"
	"gin-demo/config"
	"gin-demo/core/logger"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/gorm/utils"
)

var db *gorm.DB

// ConnectDB connect to database
func ConnectDB() {
	databaseURL := config.DB.DATABASE_URL
	var err error

	mysqlConfig := gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.DB.PREFIX, // table name prefix, table for `User` would be `t_users`
			SingularTable: true,             // use singular table name, table for `User` would be `user` with this option enabled
		},
		Logger: newLogger(),
	}
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       databaseURL, // data source name
		DefaultStringSize:         255,         // default size for string fields
		DisableDatetimePrecision:  true,        // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,        // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,        // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,       // auto configure based on currently MySQL version
	}), &mysqlConfig)
	if err != nil {
		panic(err)
	}
	mysqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	//设置与数据库建立连接的最大数目
	mysqlDB.SetMaxOpenConns(config.DB.MaxOpenConns)
	//设置连接池中的最大闲置连接数
	mysqlDB.SetMaxIdleConns(config.DB.MaxIdleConns)
	mysqlDB.SetMaxOpenConns(config.DB.ConnMaxLifeTime)

}

// DB inject with ctx for log
func DB(ctx *gin.Context) *gorm.DB {
	return db.WithContext(ctx)
}

//DisconnectDB disconnect database
func DisconnectDB() {
	mysqlDB, _ := db.DB()
	if err := mysqlDB.Close(); err != nil {
		panic(err)
	}
}

func newLogger() gLogger.Interface {
	var (
		traceStr     = "[%.3fms] [rows:%v] %s"
		traceWarnStr = "%s[%.3fms] [rows:%v] %s"
		traceErrStr  = "%s[%.3fms] [rows:%v] %s"
	)
	return &gormLogger{
		SlowThreshold: config.DB.SlowThreshold,
		traceStr:      traceStr,
		traceWarnStr:  traceWarnStr,
		traceErrStr:   traceErrStr,
	}
}

type gormLogger struct {
	SlowThreshold                       time.Duration
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
	gLogger.Interface
}

// LogMode log mode
func (l *gormLogger) LogMode(level gLogger.LogLevel) gLogger.Interface {
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
	case err != nil:
		sql, rows := fc()
		message := ""
		if rows == -1 {
			message = fmt.Sprintf(l.traceErrStr, err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			message = fmt.Sprintf(l.traceErrStr, err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
		logData := getFileWithLineNum(utils.FileWithLineNum(), message)
		panic(logData)
	case elapsed > time.Duration(l.SlowThreshold)*time.Second && l.SlowThreshold != 0:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", time.Duration(l.SlowThreshold*time.Second).Seconds())
		message := ""
		if rows == -1 {
			message = fmt.Sprintf(l.traceWarnStr, slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			message = fmt.Sprintf(l.traceWarnStr, slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
		logData := getFileWithLineNum(utils.FileWithLineNum(), message)
		logger.Warn(ctx, "gorm", logData)
	default:
		sql, rows := fc()
		message := ""
		if rows == -1 {
			message = fmt.Sprintf(l.traceStr, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			message = fmt.Sprintf(l.traceStr, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
		logData := getFileWithLineNum(utils.FileWithLineNum(), message)
		logger.Debug(ctx, "gorm", logData)
	}
}

func getFileWithLineNum(str string, message string) logger.CustomLog {
	FileAndLineNum := strings.Split(str, ":")
	len := len(FileAndLineNum)
	file := strings.Join(FileAndLineNum[0:len-1], ":")
	fmt.Print(file)
	line, _ := strconv.Atoi(FileAndLineNum[len-1])
	logData := logger.CustomLog{
		File:    file,
		Line:    line,
		Message: message,
	}
	return logData
}
