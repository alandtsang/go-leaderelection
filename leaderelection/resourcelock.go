package leaderelection

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/alandtsang/go-leaderelection/datasource/mysql/dal/model"
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

	recordByte, err := json.Marshal(*record)
	if err != nil {
		return nil, nil, err
	}

	return record, recordByte, nil
}

// Create attempts to create a LeaderElectionRecord.
func (l *mysqlResourceLock) Create(ctx context.Context, ler resourcelock.LeaderElectionRecord) error {
	return nil
}

// Update will update and existing LeaderElectionRecord.
func (l *mysqlResourceLock) Update(ctx context.Context, ler resourcelock.LeaderElectionRecord) error {
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
