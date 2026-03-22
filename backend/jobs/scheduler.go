package jobs

import (
	"log"
	"skillspark/internal/notification"
	"skillspark/internal/storage"
	"skillspark/internal/stripeClient"

	"github.com/robfig/cron/v3"
)

type JobScheduler struct {
	cron         *cron.Cron
	repo         *storage.Repository
	stripeClient stripeClient.StripeClientInterface
	notifService notification.Service
}

func NewJobScheduler(repo *storage.Repository, sc stripeClient.StripeClientInterface, notif notification.Service) *JobScheduler {
	return &JobScheduler{
		cron:         cron.New(),
		repo:         repo,
		stripeClient: sc,
		notifService: notif,
	}
}

func (j *JobScheduler) Start() {
	_, err := j.cron.AddFunc("0 * * * *", func() {
		log.Println("Running payment capture job...")
		j.CapturePaymentsJob()
	})
	if err != nil {
		log.Fatalf("Failed to schedule payment capture job: %v", err)
	}

	_, err = j.cron.AddFunc("*/5 * * * *", func() {
		log.Println("Running scheduled notification job...")
		j.SendScheduledNotificationsJob()
	})
	if err != nil {
		log.Fatalf("Failed to schedule notification job: %v", err)
	}

	j.cron.Start()
	log.Println("Cron jobs started")

	go j.CapturePaymentsJob()
	go j.SendScheduledNotificationsJob()
}

func (j *JobScheduler) Stop() {
	j.cron.Stop()
}
