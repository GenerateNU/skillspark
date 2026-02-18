package jobs

import (
	"log"
	"skillspark/internal/storage"
	"skillspark/internal/stripeClient"

	"github.com/robfig/cron/v3"
)

type JobScheduler struct {
	cron         *cron.Cron
	repo         *storage.Repository
	stripeClient stripeClient.StripeClientInterface
}

func NewJobScheduler(repo *storage.Repository, sc stripeClient.StripeClientInterface) *JobScheduler {
	return &JobScheduler{
		cron:         cron.New(),
		repo:         repo,
		stripeClient: sc,
	}
}

func (j *JobScheduler) Start() {
	
	j.cron.AddFunc("0 * * * *", func() {
		log.Println("Running payment capture job...")
		j.CapturePaymentsJob()
	})

	j.cron.Start()
	log.Println("Cron jobs started")
}

func (j *JobScheduler) Stop() {
	j.cron.Stop()
}