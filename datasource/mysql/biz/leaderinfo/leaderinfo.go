package leaderinfo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/alandtsang/go-leaderelection/datasource/mysql/dal"
	"github.com/alandtsang/go-leaderelection/datasource/mysql/dal/model"
	"github.com/alandtsang/go-leaderelection/datasource/mysql/dal/query"
)

var q = query.Use(dal.DB.Debug())

type CreateLeaderInfoParams struct {
	LeaderIdentity       string
	Transitions          uint64
	LeaseDurationSeconds uint64
	LastAcquireAt        time.Time
	LastRenewAt          time.Time
}

func Create(ctx context.Context, p *CreateLeaderInfoParams) error {
	li := q.LeaderInfo.WithContext(ctx)
	now := time.Now()

	x := model.LeaderInfo{
		LeaderIdentity:       p.LeaderIdentity,
		Transitions:          p.Transitions,
		LeaseDurationSeconds: p.LeaseDurationSeconds,
		CreateAt:             &now,
		LastAcquireAt:        &p.LastAcquireAt,
		LastRenewAt:          &p.LastRenewAt,
	}
	return li.Create(&x)
}

type UpdateLeaderInfoParams struct {
	ID                   uint64
	LeaderIdentity       string
	Transitions          uint64
	LeaseDurationSeconds uint64
	LastAcquireAt        time.Time
	LastRenewAt          time.Time
}

func Update(ctx context.Context, p *UpdateLeaderInfoParams) error {
	li := q.LeaderInfo.WithContext(ctx)
	to := make(map[string]interface{})

	if p.ID != 0 {
		to["id"] = p.ID
	}
	if p.Transitions != 0 {
		to["transitions"] = p.Transitions
	}
	if p.LeaseDurationSeconds != 0 {
		to["lease_duration_seconds"] = p.LeaseDurationSeconds
	}
	if p.LastAcquireAt.Unix() != 0 {
		to["last_acquire_at"] = p.LastAcquireAt
	}
	if p.LastRenewAt.Unix() != 0 {
		to["last_renew_at"] = p.LastRenewAt
	}

	if _, err := li.Where(q.LeaderInfo.LeaderIdentity.Eq(p.LeaderIdentity)).Updates(to); err != nil {
		log.Printf("Update LeaderInfo failed, %v", err)
		return err
	}

	return nil
}

func Fetch(ctx context.Context, leaderID string) (*model.LeaderInfo, error) {
	li := q.LeaderInfo.WithContext(ctx)
	leaderInfo, err := li.Where(q.LeaderInfo.LeaderIdentity.Eq(leaderID)).First()
	if err != nil {
		fmt.Printf("Fetch LeaderInfo %s failed, %v", leaderID, err)
		return nil, err
	}
	return leaderInfo, nil
}
