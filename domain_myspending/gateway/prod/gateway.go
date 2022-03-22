package prod

import (
	"context"
	"your/path/project/domain_myspending/model/entity"
	"your/path/project/shared/driver"
	"your/path/project/shared/infrastructure/config"
	"your/path/project/shared/infrastructure/logger"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type gateway struct {
	log     logger.Logger
	appData driver.ApplicationData
	config  *config.Config
	db      *gorm.DB
}

// NewGateway ...
func NewGateway(log logger.Logger, appData driver.ApplicationData, config *config.Config) *gateway {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&entity.Expense{})
	if err != nil {
		panic("cannot create schema")
	}

	return &gateway{
		log:     log,
		appData: appData,
		config:  config,
		db:      db,
	}
}

func (r *gateway) SaveExpense(ctx context.Context, obj *entity.Expense) error {
	r.log.Info(ctx, "called")

	err := r.db.Create(obj).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gateway) FindAllExpense(ctx context.Context, someID string) ([]*entity.Expense, error) {
	r.log.Info(ctx, "called")

	var result []*entity.Expense

	err := r.db.Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
