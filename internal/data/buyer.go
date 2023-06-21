package data

import (
	"context"
	"database/sql"
	biz "ecommerce/internal/biz/v1"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type buyerRepo struct {
	data *Data
	log  *log.Helper
}

func NewBuyerRepo(data *Data, logger log.Logger) biz.BuyerRepo {
	return &buyerRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *buyerRepo) New(ctx context.Context, email, password string) (int64, error) {
	tx, err := r.data.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return 0, err
	}
	defer func() {
		if err != nil {
			r.log.WithContext(ctx).Warnf("Fail new buyer: %s by %v", email, err)
			if err = tx.Rollback(); err != nil {
				r.log.Errorf("can not rollback tx for new buyer action with buyer %s", email)
			}
		}
		r.log.WithContext(ctx).Warnf("Finish new buyer: %s", email)
	}()
	var id int64
	err = tx.QueryRowContext(ctx, `INSERT INTO "buyers" ("email", "password", "created_at") VALUES ($1, $2, $3) RETURNING id `, email, password, time.Now().UTC()).Scan(&id)
	if err != nil {
		return 0, err
	}
	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *buyerRepo) GetByEmail(ctx context.Context, email string) (int64, string, error) {
	row := r.data.db.QueryRowContext(ctx, "SELECT id, password FROM buyers WHERE email = $1", email)
	if row.Err() != nil {
		return 0, "", row.Err()
	}
	var id int64
	var password string
	row.Scan(&id, &password)

	return id, password, nil
}
