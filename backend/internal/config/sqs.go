package config

type SQS struct {
	QueueURL string `env:"AWS_SQS_QUEUE_URL, required"`
	Region   string `env:"AWS_REGION, required"`
	// Can reuse from S3 config, but making explicit for clarity
	AccessKey string `env:"AWS_ACCESS_KEY_ID, required"`
	SecretKey string `env:"AWS_SECRET_ACCESS_KEY, required"`
}

