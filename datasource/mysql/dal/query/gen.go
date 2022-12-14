// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
)

var (
	Q          = new(Query)
	LeaderInfo *leaderInfo
)

func SetDefault(db *gorm.DB) {
	*Q = *Use(db)
	LeaderInfo = &Q.LeaderInfo
}

func Use(db *gorm.DB) *Query {
	return &Query{
		db:         db,
		LeaderInfo: newLeaderInfo(db),
	}
}

type Query struct {
	db *gorm.DB

	LeaderInfo leaderInfo
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:         db,
		LeaderInfo: q.LeaderInfo.clone(db),
	}
}

type queryCtx struct {
	LeaderInfo *leaderInfoDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		LeaderInfo: q.LeaderInfo.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	return &QueryTx{q.clone(q.db.Begin(opts...))}
}

type QueryTx struct{ *Query }

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
