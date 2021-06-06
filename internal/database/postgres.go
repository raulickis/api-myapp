package database

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/raulickis/api-myapp/config"
	db "github.com/raulickis/api-myapp/tools"
	"go.elastic.co/apm/module/apmgorm"
	_ "go.elastic.co/apm/module/apmgorm/dialects/postgres"
	"sync"
)

type Repository struct {
	dbPostgres *gorm.DB
	once       sync.Once
}

//StartPostgres start the DB
func (r *Repository) Start() error {
	err := db.LoadGormPostGres(
		config.DB_USER,
		config.DB_PASS,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
		false)
	return err
}

//StopPostgres stop the DB
func (r *Repository) Stop() {
	defer r.dbPostgres.Close()
}

// GetInstance returns a unique instance of gorm.DB
func (r *Repository) GetInstance() *gorm.DB {
	r.once.Do(func() {
		var err error
		r.dbPostgres, err = db.GetGormDb()
		if err != nil {
			panic(err.Error())
		}
		r.dbPostgres.SingularTable(true)
		r.dbPostgres.LogMode(true)
	})
	return r.dbPostgres
}

// GetInstance returns a unique instance of gorm.DB
func (r *Repository) GetInstanceCtx(ctx context.Context) *gorm.DB {
	r.once.Do(func() {
		var err error
		r.dbPostgres, err = db.GetGormDb()
		if err != nil {
			panic(err.Error())
		}
		r.dbPostgres.SingularTable(true)
		r.dbPostgres.LogMode(true)
	})
	var database = apmgorm.WithContext(ctx, r.dbPostgres)
	return database
}
