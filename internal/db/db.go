package db

import (
	"time"

	"github.com/cham/elk/internal/db/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

func InitDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 添加 OpenTelemetry 插件用于 tracing
	if err := db.Use(tracing.NewPlugin()); err != nil {
		return nil, err
	}

	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 自动迁移模型
	err = db.AutoMigrate(&model.User{}, &model.Schema{}, &model.Code{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// 创建一个中间件来添加 metrics
func MetricsMiddleware(db *gorm.DB) {
	db.Callback().Create().Before("gorm:create").Register("metrics:create", func(db *gorm.DB) {
		// 在这里添加创建操作的 metrics
	})

	db.Callback().Query().Before("gorm:query").Register("metrics:query", func(db *gorm.DB) {
		// 在这里添加查询操作的 metrics
	})

	// 可以为其他操作添加类似的 metrics
}
