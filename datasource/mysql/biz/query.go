package biz

import (
	"context"
	"fmt"

	"github.com/alandtsang/go-leaderelection/datasource/mysql/dal"
	"github.com/alandtsang/go-leaderelection/datasource/mysql/dal/model"
	"github.com/alandtsang/go-leaderelection/datasource/mysql/dal/query"
)

var q = query.Use(dal.DB.Debug())

func Fetch(ctx context.Context, leaderID string) (*model.LeaderInfo, error) {
	c, cd := q.LeaderInfo, q.LeaderInfo.WithContext(ctx)
	person, err := cd.Where(c.LeaderIdentity.Eq(leaderID)).First()
	if err != nil {
		fmt.Printf("Fetch LeaderInfo %s failed, %v", leaderID, err)
		return nil, err
	}
	return person, nil
}
