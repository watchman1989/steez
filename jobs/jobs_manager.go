package jobs

import (
	"sync"
	"time"

	"github.com/watchman1989/steez/comm"

	"github.com/robfig/cron/v3"
	"github.com/watchman1989/utils/ft"
)

func InitJobs() (jm *JobManager) {
	jm = NewJobManager()
	jm.addJob("sdnList", "@every 10s")
	return
}

type JobManager struct {
	jobs map[string]*cron.Cron
	mu sync.Mutex
	done chan struct{}
}

func NewJobManager() *JobManager {
	jm := &JobManager{
		jobs: make(map[string]*cron.Cron),
		done: make(chan struct{}),
	}
	go jm.run()
	return jm
}

func (jm *JobManager) run() {
	ft, err := ft.NewFloorTicker("@every 10s")
	if err != nil {
		comm.GContext.Logger.Errorf("new floor ticker error: %s", err.Error())
		return
	}
	for {
		select {
		case <-ft.C:
			jm.mu.Lock()
			for name := range jm.jobs {
				comm.GContext.Logger.Infof("job %s is running", name)
			}
			jm.mu.Unlock()
		case <-jm.done:
			jm.stopJobs()
			comm.GContext.Logger.Infof("job manager exit")
			return
		}
	}
	
}

func (jm *JobManager) stopJobs() {
	jm.mu.Lock()
	defer jm.mu.Unlock()
	for name, job := range jm.jobs {
		job.Stop()
		comm.GContext.Logger.Infof("stop job %s", name)
	}
}


func (jm *JobManager) Stop(){
	close(jm.done)
}


func (jm *JobManager) addJob(name string, spec string) (err error) {
	jm.mu.Lock()
	defer jm.mu.Unlock()

	job := cron.New(cron.WithSeconds())
	_, err = job.AddJob(spec,
		cron.NewChain(cron.SkipIfStillRunning(cron.DefaultLogger)).Then(&SdnListJob{}))
	if err != nil {
		comm.GContext.Logger.Errorf("add job %s error: %s", name, err.Error())
		return
	}
	jm.jobs[name] = job
	jm.jobs[name].Start()
	comm.GContext.Logger.Infof("add job %s success", name)
	return
}


type SdnListJob struct {}
func (s *SdnListJob) Run() {
	comm.GContext.Logger.Infof("sdn list job run start at %s", time.Now().Format("2006-01-02 15:04:05"))
	defer comm.GContext.Logger.Infof("sdn list job run end at %s", time.Now().Format("2006-01-02 15:04:05"))
}