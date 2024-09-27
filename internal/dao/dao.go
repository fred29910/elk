package dao

import (
	"context"
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"

	// "gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/cham/elk/internal/config"
)

var db *gorm.DB

func Init() error {
	dbConfig := config.GetDbConfig()

	var err error
	if dbConfig.Model == "sqlite" {
		db, err = gorm.Open(sqlite.Open(dbConfig.Sqlite.Path), &gorm.Config{})
	} else {
		dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.Mysql.User, dbConfig.Mysql.Password, dbConfig.Mysql.Host, dbConfig.Mysql.Port, dbConfig.Mysql.Database)
		db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	}
	return err
}

type contextTxKey struct{}

func DB(ctx context.Context) *gorm.DB {

	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return db
}

func InTransaction(ctx context.Context, fn func(context.Context) error) error {
	return db.Transaction(func(tx *gorm.DB) error {
		psCtx := context.WithValue(ctx, contextTxKey{}, tx)
		err := fn(psCtx)
		if err != nil {
			// 回滚事务
			tx.Rollback()
			return err
		}
		// 提交事务
		return tx.Commit().Error
	})
}

// page
func Pageble(tx *gorm.DB, page int, pageSize int) *gorm.DB {
	p := 1
	if page > 0 {
		p = page
	}
	ps := 10
	if pageSize > 0 && pageSize < 100 {
		ps = pageSize
	}
	if pageSize > 100 {
		ps = 100
	}
	offset := (p - 1) * ps
	return tx.Offset(offset).Limit(ps)
}
