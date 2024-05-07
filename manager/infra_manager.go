package manager

import (
	"eniqloStore/config"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Infra interface {
	TokenConfig() config.TokenConfig

	SqlDb() *gorm.DB
}

type infra struct {
	tokenConfig config.TokenConfig

	dbResource *gorm.DB
}

func (i *infra) TokenConfig() config.TokenConfig {
	return i.tokenConfig
}

func (i *infra) SqlDb() *gorm.DB {
	return i.dbResource
}

func NewInfra(config config.Config) Infra {

	resource, err := initDbResource(config.DataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Print("Database Connected!")

	return &infra{dbResource: resource, tokenConfig: config.TokenConfig} // tokenConfig: config.TokenConfig

}

// FIXME: ubah ini ke go-migration

func initDbResource(dataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})

	env := os.Getenv("ENV")
	dbReturn := db
	if env == "migration" {
		dbReturn = db.Debug()

		// nambain model tabel disini

		db.AutoMigrate(
		// add db disini buat auto migrate

		// &model.User{},
		)

		//masukin table untuk dimigrate
	} else if env == "dev" {
		dbReturn = db.Debug()
	}
	if err != nil {
		return nil, err
	}
	return dbReturn, nil
}
