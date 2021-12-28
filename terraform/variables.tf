variable "region" {
  description = "Bucket region"
  type        = string
}

variable "bucket_name" {
  description = "Bucket name"
  type        = string
}

variable "enable_versioning" {
  description = "Bucket versioning"
  type        = bool
  default     = true
}

variable "tags" {
  description = "Additional resource tags"
  type        = map(string)
  default     = {}
}
