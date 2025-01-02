package scheduler

import (
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/sirupsen/logrus"
)

var s gocron.Scheduler

func Init() {
    s, err := gocron.NewScheduler(gocron.WithLocation(time.UTC))
    if err != nil {
		logrus.Errorf("Failed to run Scheduler: ", err)
    }else{
		logrus.Infof("Scheduler initialized")
    }
    s.Start()
}

func AddJob(second int, cmd func()) {
    j, err := s.NewJob(
        gocron.DurationJob(
            time.Duration(second) * time.Second,
        ),
        gocron.NewTask(cmd),
    )
	if err != nil {
		logrus.Errorf("Failed to add job: ", err)
	}else{
        logrus.Infof("Job added: %v", j)
    }
}

func Stop() {
    err := s.Shutdown()
	if err != nil {
		logrus.Errorf("Failed to shutdown Scheduler: ", err)
	}else{
        logrus.Infof("Scheduler stopped")
    }
}