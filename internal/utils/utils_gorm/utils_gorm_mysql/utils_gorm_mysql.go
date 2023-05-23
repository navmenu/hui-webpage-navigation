package utils_gorm_mysql

import (
	"hui-webpage-navigation/internal/utils/utils_gorm"
	"hui-webpage-navigation/internal/utils/utils_log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGormDB(dsn string, echo bool) (db *gorm.DB, err error) {
	return NewMysqlGormDB(&Config{Dsn: dsn, Echo: echo})
}

type Config = utils_gorm.Config

func NewMysqlGormDB(cfg *Config) (db *gorm.DB, err error) {
	xLog := logger.Default.LogMode(logger.Silent)
	if cfg.Echo {
		xLog = logger.New(
			utils_log.NewShortDebug(os.Stdout, "postgres"), // IO WRITER
			logger.Config{
				SlowThreshold:             time.Second, // 慢 SQL 阈值
				LogLevel:                  logger.Info, // 日志级别
				IgnoreRecordNotFoundError: true,        // 禁用 ErrRecordNotFound（记录未找到）错误
				Colorful:                  true,        // 是否彩色打印
			},
		)
	}
	opts := &gorm.Config{
		AllowGlobalUpdate: true, // ALLOW GLOBAL UPDATE
		Logger:            xLog,
	}
	db, err = gorm.Open(mysql.Open(cfg.Dsn), opts)
	if err != nil {
		return nil, err
	}
	rawDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	if cfg.MaxOpenConns > 0 {
		rawDB.SetMaxOpenConns(cfg.MaxOpenConns)
	}
	return db, nil
}
