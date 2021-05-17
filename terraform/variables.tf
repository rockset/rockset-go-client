variable "bucket" {
  description = "Name of the S3 bucket to create with integration with Rockset"
  type        = string
  default     = "rockset-go-tests"
}

variable "rockset-external-id" {
  description = "External id provided by the Rockset Console"
  type        = string
  // TODO: once the rockset provider is available, use the data source to get this
  default     = "4ef12b664c96efcd836edfef7b3e085e908f46052b4dcaec3faac57a5b08048e"
}

variable "rockset-account-id" {
  description = "Account id provided by the Rockset Console"
  type        = string
  // TODO: once the rockset provider is available, use the data source to get this
  default     = "318212636800"
}

variable "region" {
  default = "us-west-2"
  type    = string
}

variable "rockset-role-name" {
  default = "rockset-s3-integration"
  type    = string
}