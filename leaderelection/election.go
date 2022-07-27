package leaderelection

import (
	"context"
	"time"

	"k8s.io/client-go/tools/leaderelection"
)

func New(
	onStartLeading func(context.Context),
	onStopLeading func(),
	onNewLeader func(string)) *leaderelection.LeaderElector {
	// pass callbacks
	conf := leaderelection.LeaderElectionConfig{
		Lock:            GetResourceLock(),
		LeaseDuration:   time.Second * 30,
		RenewDeadline:   time.Second * 20,
		RetryPeriod:     time.Second * 10,
		ReleaseOnCancel: true,
		Name:            GetResourceLock().Identity(),
		Callbacks: leaderelection.LeaderCallbacks{
			OnStartedLeading: onStartLeading,
			OnStoppedLeading: onStopLeading,
			OnNewLeader:      onNewLeader,
		},
	}
	elector, err := leaderelection.NewLeaderElector(conf)
	if err != nil {
		panic(err)
	}

	return elector
}
