variable "bucket_name" {
  description = "Name of the S3 bucket (use the existing bucket name when importing)"
  type        = string
}

variable "tags" {
  description = "Tags to apply to the bucket"
  type        = map(string)
  default     = {}
}
