package model

import "time"

type LeaderInfo struct {
	ID                   uint64     `gorm:"column:id;type:bigint unsigned AUTO_INCREMENT;PRIMARY_KEY"`
	LeaderIdentity       string     `gorm:"column:leader_identity;type:varchar(128); NOT NULL DEFAULT ''"`
	Transitions          uint64     `gorm:"column:transitions;type:bigint unsigned;NOT NULL DEFAULT 0"`
	LeaseDurationSeconds uint64     `gorm:"column:lease_duration_seconds;type:bigint unsigned;NOT NULL DEFAULT 0"`
	CreateAt             *time.Time `gorm:"column:created_at;type:timestamp;NOT NULL DEFAULT now()"`
	LastAcquireAt        *time.Time `gorm:"column:last_acquire_at;type:timestamp;NOT NULL DEFAULT now()"`
	LastRenewAt          *time.Time `gorm:"column:last_renew_at;type:timestamp;NOT NULL DEFAULT now()"`
}

func (LeaderInfo) TableName() string {
	return "leaderinfo"
}

func NewLeaderInfo() *LeaderInfo {
	return &LeaderInfo{}
}
