# S3 bucket (import existing bucket into this module)
module "dev_bucket" {
  source = "../modules/s3-bucket"
  bucket_name = "skillspark-dev"
  tags = {
    Project     = var.project_name
    Environment = "dev"
  }
}

import {
  to = module.dev_bucket.aws_s3_bucket.this
  id = "skillspark-dev"
}

module "prod_bucket" {
  source = "../modules/s3-bucket"
  bucket_name = "skillspark-prod"
  tags = {
    Project     = var.project_name
    Environment = "prod"
  }
}

import {
  to = module.prod_bucket.aws_s3_bucket.this
  id = "skillspark-prod"
}