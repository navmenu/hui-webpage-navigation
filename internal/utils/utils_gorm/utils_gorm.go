package utils_gorm

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type GormDB gorm.DB

func NewGormDB(db *gorm.DB) (gormDB *GormDB) {
	return (*GormDB)(db)
}

func (T *GormDB) DB() *gorm.DB {
	return (*gorm.DB)(T)
}

func (T *GormDB) Save(value interface{}) *gorm.DB {
	return T.DB().Save(value)
}

func (T *GormDB) Transaction(run func(db *gorm.DB) error) error {
	return T.DB().Transaction(func(db *gorm.DB) error {
		return run(db)
	})
}

func (T *GormDB) TransactionExec(run func(*GormDB) error) error {
	return T.DB().Transaction(func(db *gorm.DB) error {
		return run(NewGormDB(db))
	})
}

func (T *GormDB) TransactionSafe(run func(*GormDB)) error {
	return T.DB().Transaction(func(db *gorm.DB) error {
		run(NewGormDB(db))
		return nil
	})
}

// SELECT * FROM users FOR UPDATE
func (T *GormDB) LockingUpdate() *GormDB {
	return NewGormDB(T.DB().Clauses(clause.Locking{Strength: "UPDATE"}))
}

// SELECT * FROM users FOR SHARE OF users
func (T *GormDB) LockingShareTable() *GormDB {
	return NewGormDB(T.DB().Clauses(clause.Locking{Strength: "SHARE", Table: clause.Table{Name: clause.CurrentTable}}))
}

// SELECT * FROM users FOR UPDATE NOWAIT
func (T *GormDB) LockingUpdateNowait() *GormDB {
	return NewGormDB(T.DB().Clauses(clause.Locking{Strength: "UPDATE", Options: "NOWAIT"}))
}

func (T *GormDB) Close() error {
	return Close(T.DB())
}

func Close(db *gorm.DB) error {
	rawDB, err := db.DB()
	if err != nil {
		return err
	}
	return rawDB.Close()
}

// SilentDB 这里不需要关闭因为会关闭最底层的连接
func SilentDB(db *gorm.DB) *gorm.DB {
	var sdb = db.Debug()
	sdb.Logger = logger.Default.LogMode(logger.Silent)
	return sdb
}

func Ping(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Exec("SELECT 200").Error
}
