package s3_client

import (
	"context"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/joho/godotenv"
)

// TestAWSCredentials tests to see if the credentials being used are valid + work
func init() {
	// Load .env from backend root
	// Current file: backend/internal/s3_client/aws_creds_test.go
	// Target file:  backend/.env
	// Go up 2 levels: ../../.env
	_ = godotenv.Load("../../.env")
}

func TestAWSCredentials(t *testing.T) {
	t.Log("Testing AWS credentials...")
	t.Log("================================")

	// Verify credentials are loaded
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_REGION")

	if accessKey == "" || secretKey == "" || region == "" {
		t.Fatal("❌ AWS credentials not found. Make sure backend/.env exists with:\n" +
			"   AWS_ACCESS_KEY_ID=...\n" +
			"   AWS_SECRET_ACCESS_KEY=...\n" +
			"   AWS_REGION=...")
	}

	t.Logf("✅ Loaded credentials from .env")
	t.Logf("   AWS_ACCESS_KEY_ID: %s...", accessKey[:4])
	t.Logf("   AWS_REGION: %s", region)

	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		t.Fatalf("❌ Failed to load AWS config: %v", err)
	}

	// Test 1: Get caller identity
	t.Log("\n🔐 Testing credentials...")
	stsClient := sts.NewFromConfig(cfg)
	identity, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		t.Fatalf("❌ Failed to get caller identity: %v", err)
	}

	t.Log("✅ Credentials are valid!")
	t.Logf("   Account: %s", *identity.Account)
	t.Logf("   User ARN: %s", *identity.Arn)
	t.Logf("   User ID: %s", *identity.UserId)

	// Test 2: List S3 buckets
	t.Log("\n🪣 Testing S3 access...")
	s3Client := s3.NewFromConfig(cfg)
	result, err := s3Client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		t.Logf("⚠️  Failed to list buckets: %v", err)
		t.Log("   (Your credentials might not have S3 permissions)")
	} else {
		t.Logf("✅ S3 access works! Found %d bucket(s):", len(result.Buckets))
		for _, bucket := range result.Buckets {
			t.Logf("   - %s", *bucket.Name)
		}
	}
}
