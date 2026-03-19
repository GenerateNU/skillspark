output "dev_bucket_id" {
  description = "The name (id) of the bucket"
  value       = module.dev_bucket.bucket_id
}

output "dev_bucket_arn" {
  description = "The ARN of the bucket"
  value       = module.dev_bucket.bucket_arn
}

output "dev_bucket_domain_name" {
  description = "The bucket domain name"
  value       = module.dev_bucket.bucket_domain_name
}


output "prod_bucket_id" {
  description = "The name (id) of the bucket"
  value       = module.prod_bucket.bucket_id
}

output "prod_bucket_arn" {
  description = "The ARN of the bucket"
  value       = module.prod_bucket.bucket_arn
}

output "prod_bucket_domain_name" {
  description = "The bucket domain name"
  value       = module.prod_bucket.bucket_domain_name
}