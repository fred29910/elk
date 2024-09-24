package dao

import (
	"context"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(dns string) {
	var err error
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
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
