// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/alandtsang/go-leaderelection/datasource/mysql/dal/model"
)

func newLeaderInfo(db *gorm.DB) leaderInfo {
	_leaderInfo := leaderInfo{}

	_leaderInfo.leaderInfoDo.UseDB(db)
	_leaderInfo.leaderInfoDo.UseModel(&model.LeaderInfo{})

	tableName := _leaderInfo.leaderInfoDo.TableName()
	_leaderInfo.ALL = field.NewField(tableName, "*")
	_leaderInfo.ID = field.NewUint64(tableName, "id")
	_leaderInfo.LeaderIdentity = field.NewString(tableName, "leader_identity")
	_leaderInfo.Transitions = field.NewUint64(tableName, "transitions")
	_leaderInfo.LeaseDurationSeconds = field.NewUint64(tableName, "lease_duration_seconds")
	_leaderInfo.CreateAt = field.NewTime(tableName, "created_at")
	_leaderInfo.LastAcquireAt = field.NewTime(tableName, "last_acquire_at")
	_leaderInfo.LastRenewAt = field.NewTime(tableName, "last_renew_at")

	_leaderInfo.fillFieldMap()

	return _leaderInfo
}

type leaderInfo struct {
	leaderInfoDo leaderInfoDo

	ALL                  field.Field
	ID                   field.Uint64
	LeaderIdentity       field.String
	Transitions          field.Uint64
	LeaseDurationSeconds field.Uint64
	CreateAt             field.Time
	LastAcquireAt        field.Time
	LastRenewAt          field.Time

	fieldMap map[string]field.Expr
}

func (l leaderInfo) Table(newTableName string) *leaderInfo {
	l.leaderInfoDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l leaderInfo) As(alias string) *leaderInfo {
	l.leaderInfoDo.DO = *(l.leaderInfoDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *leaderInfo) updateTableName(table string) *leaderInfo {
	l.ALL = field.NewField(table, "*")
	l.ID = field.NewUint64(table, "id")
	l.LeaderIdentity = field.NewString(table, "leader_identity")
	l.Transitions = field.NewUint64(table, "transitions")
	l.LeaseDurationSeconds = field.NewUint64(table, "lease_duration_seconds")
	l.CreateAt = field.NewTime(table, "created_at")
	l.LastAcquireAt = field.NewTime(table, "last_acquire_at")
	l.LastRenewAt = field.NewTime(table, "last_renew_at")

	l.fillFieldMap()

	return l
}

func (l *leaderInfo) WithContext(ctx context.Context) *leaderInfoDo {
	return l.leaderInfoDo.WithContext(ctx)
}

func (l leaderInfo) TableName() string { return l.leaderInfoDo.TableName() }

func (l leaderInfo) Alias() string { return l.leaderInfoDo.Alias() }

func (l *leaderInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *leaderInfo) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 7)
	l.fieldMap["id"] = l.ID
	l.fieldMap["leader_identity"] = l.LeaderIdentity
	l.fieldMap["transitions"] = l.Transitions
	l.fieldMap["lease_duration_seconds"] = l.LeaseDurationSeconds
	l.fieldMap["created_at"] = l.CreateAt
	l.fieldMap["last_acquire_at"] = l.LastAcquireAt
	l.fieldMap["last_renew_at"] = l.LastRenewAt
}

func (l leaderInfo) clone(db *gorm.DB) leaderInfo {
	l.leaderInfoDo.ReplaceDB(db)
	return l
}

type leaderInfoDo struct{ gen.DO }

func (l leaderInfoDo) Debug() *leaderInfoDo {
	return l.withDO(l.DO.Debug())
}

func (l leaderInfoDo) WithContext(ctx context.Context) *leaderInfoDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l leaderInfoDo) ReadDB() *leaderInfoDo {
	return l.Clauses(dbresolver.Read)
}

func (l leaderInfoDo) WriteDB() *leaderInfoDo {
	return l.Clauses(dbresolver.Write)
}

func (l leaderInfoDo) Clauses(conds ...clause.Expression) *leaderInfoDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l leaderInfoDo) Returning(value interface{}, columns ...string) *leaderInfoDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l leaderInfoDo) Not(conds ...gen.Condition) *leaderInfoDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l leaderInfoDo) Or(conds ...gen.Condition) *leaderInfoDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l leaderInfoDo) Select(conds ...field.Expr) *leaderInfoDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l leaderInfoDo) Where(conds ...gen.Condition) *leaderInfoDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l leaderInfoDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *leaderInfoDo {
	return l.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (l leaderInfoDo) Order(conds ...field.Expr) *leaderInfoDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l leaderInfoDo) Distinct(cols ...field.Expr) *leaderInfoDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l leaderInfoDo) Omit(cols ...field.Expr) *leaderInfoDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l leaderInfoDo) Join(table schema.Tabler, on ...field.Expr) *leaderInfoDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l leaderInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) *leaderInfoDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l leaderInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) *leaderInfoDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l leaderInfoDo) Group(cols ...field.Expr) *leaderInfoDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l leaderInfoDo) Having(conds ...gen.Condition) *leaderInfoDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l leaderInfoDo) Limit(limit int) *leaderInfoDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l leaderInfoDo) Offset(offset int) *leaderInfoDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l leaderInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *leaderInfoDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l leaderInfoDo) Unscoped() *leaderInfoDo {
	return l.withDO(l.DO.Unscoped())
}

func (l leaderInfoDo) Create(values ...*model.LeaderInfo) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l leaderInfoDo) CreateInBatches(values []*model.LeaderInfo, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l leaderInfoDo) Save(values ...*model.LeaderInfo) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l leaderInfoDo) First() (*model.LeaderInfo, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.LeaderInfo), nil
	}
}

func (l leaderInfoDo) Take() (*model.LeaderInfo, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.LeaderInfo), nil
	}
}

func (l leaderInfoDo) Last() (*model.LeaderInfo, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.LeaderInfo), nil
	}
}

func (l leaderInfoDo) Find() ([]*model.LeaderInfo, error) {
	result, err := l.DO.Find()
	return result.([]*model.LeaderInfo), err
}

func (l leaderInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.LeaderInfo, err error) {
	buf := make([]*model.LeaderInfo, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l leaderInfoDo) FindInBatches(result *[]*model.LeaderInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l leaderInfoDo) Attrs(attrs ...field.AssignExpr) *leaderInfoDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l leaderInfoDo) Assign(attrs ...field.AssignExpr) *leaderInfoDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l leaderInfoDo) Joins(fields ...field.RelationField) *leaderInfoDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l leaderInfoDo) Preload(fields ...field.RelationField) *leaderInfoDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l leaderInfoDo) FirstOrInit() (*model.LeaderInfo, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.LeaderInfo), nil
	}
}

func (l leaderInfoDo) FirstOrCreate() (*model.LeaderInfo, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.LeaderInfo), nil
	}
}

func (l leaderInfoDo) FindByPage(offset int, limit int) (result []*model.LeaderInfo, count int64, err error) {
	result, err = l.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = l.Offset(-1).Limit(-1).Count()
	return
}

func (l leaderInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l leaderInfoDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l leaderInfoDo) Delete(models ...*model.LeaderInfo) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *leaderInfoDo) withDO(do gen.Dao) *leaderInfoDo {
	l.DO = *do.(*gen.DO)
	return l
}
