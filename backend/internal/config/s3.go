package config

type S3 struct {
	// Real AWS (used when UseLocalStack=false)
	Bucket    string `env:"AWS_S3_BUCKET"`
	Region    string `env:"AWS_REGION"`
	AccessKey string `env:"AWS_ACCESS_KEY_ID"`
	SecretKey string `env:"AWS_SECRET_ACCESS_KEY"`

	UseLocalStack bool `env:"USE_LOCALSTACK, default=false"`

	// LocalStack (used when UseLocalStack=true)
	LocalStackEndpoint  string `env:"LOCALSTACK_ENDPOINT, default=http://localstack:4566"`
	LocalStackBucket    string `env:"LOCALSTACK_S3_BUCKET"`
	LocalStackRegion    string `env:"LOCALSTACK_REGION, default=us-east-1"`
	LocalStackAccessKey string `env:"LOCALSTACK_ACCESS_KEY_ID, default=test"`
	LocalStackSecretKey string `env:"LOCALSTACK_SECRET_ACCESS_KEY, default=test"`
}
