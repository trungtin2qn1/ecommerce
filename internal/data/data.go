package data

import (
	"database/sql"
	"ecommerce/internal/conf"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	_ "github.com/lib/pq"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewBuyerRepo, NewSellerRepo)

type Data struct {
	db *sql.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.Database.Host, c.Database.Port, c.Database.User, c.Database.Password, c.Database.DbName)
	db, err := sql.Open(c.Database.Driver, psqlInfo)
	if err != nil {
		panic(err)
	}
	if err != nil {
		return nil, nil, err
	}
	d := &Data{
		db: db,
	}
	cleanup := func() {
		log.Info("message", "closing the data resources")
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
	}
	return d, cleanup, nil
}
