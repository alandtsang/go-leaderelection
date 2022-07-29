package leaderelection

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/alandtsang/go-leaderelection/datasource/mysql/biz/leaderinfo"
	"github.com/alandtsang/go-leaderelection/datasource/mysql/dal/model"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
)

var (
	lock    *mysqlResourceLock = nil
	once    sync.Once
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

type mysqlResourceLock struct {
	identity string
	record   *model.LeaderInfo
}

func GetResourceLock() resourcelock.Interface {
	once.Do(func() {
		lock = &mysqlResourceLock{
			identity: randomName(),
			record:   model.NewLeaderInfo(),
		}
	})
	return lock
}

// Get returns the LeaderElectionRecord.
func (l *mysqlResourceLock) Get(ctx context.Context) (*resourcelock.LeaderElectionRecord, []byte, error) {
	var record *resourcelock.LeaderElectionRecord

	li, err := leaderinfo.Fetch(ctx, l.identity)
	if err != nil {
		log.Printf("mysqlResourceLock Get failed, %v", err)
		return nil, nil, err
	}

	record = &resourcelock.LeaderElectionRecord{
		HolderIdentity:       li.LeaderIdentity,
		LeaseDurationSeconds: int(li.LeaseDurationSeconds),
		AcquireTime:          v1.NewTime(*l.record.LastAcquireAt),
		RenewTime:            v1.NewTime(*l.record.LastRenewAt),
		LeaderTransitions:    int(l.record.Transitions),
	}

	recordByte, err := json.Marshal(*record)
	if err != nil {
		return nil, nil, err
	}

	return record, recordByte, nil
}

// Create attempts to create a LeaderElectionRecord.
func (l *mysqlResourceLock) Create(ctx context.Context, ler resourcelock.LeaderElectionRecord) error {
	params := &leaderinfo.CreateLeaderInfoParams{
		LeaderIdentity:       ler.HolderIdentity,
		Transitions:          uint64(ler.LeaderTransitions),
		LeaseDurationSeconds: uint64(ler.LeaseDurationSeconds),
		LastAcquireAt:        ler.AcquireTime.Time,
		LastRenewAt:          ler.RenewTime.Time,
	}
	if err := leaderinfo.Create(ctx, params); err != nil {
		log.Printf("mysqlResourceLock Create params %+v failed, %v", params, err)
		return err
	}

	return nil
}

// Update will update and existing LeaderElectionRecord.
func (l *mysqlResourceLock) Update(ctx context.Context, ler resourcelock.LeaderElectionRecord) error {
	params := &leaderinfo.UpdateLeaderInfoParams{
		ID:                   1,
		LeaderIdentity:       ler.HolderIdentity,
		Transitions:          uint64(ler.LeaderTransitions),
		LeaseDurationSeconds: uint64(ler.LeaseDurationSeconds),
		LastAcquireAt:        ler.AcquireTime.Time,
		LastRenewAt:          ler.RenewTime.Time,
	}

	if err := leaderinfo.Update(ctx, params); err != nil {
		log.Printf("mysqlResourceLock Update params %+v failed, %v", params, err)
		return err
	}

	return nil
}

// RecordEvent is used to record events.
func (l *mysqlResourceLock) RecordEvent(event string) {
	log.Printf("Leader Election %s, %s", l.identity, event)
}

// Identity will return the locks Identity.
func (l *mysqlResourceLock) Identity() string {
	return l.identity
}

// Describe is used to convert details on current resource lock into a string.
func (l *mysqlResourceLock) Describe() string {
	return fmt.Sprintf("lock/%s", l.identity)
}

func randomName() string {
	rand.Seed(time.Now().Unix())
	size := 6
	name := make([]rune, size)
	for i := 0; i < size; i++ {
		name[i] = letters[rand.Intn(len(letters))]
	}
	return string(name)
}
