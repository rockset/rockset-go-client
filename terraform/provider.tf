terraform {
  backend "s3" {
    bucket = "terraform-state-rockset-test"
    key    = "go-client"
    region = "us-west-2"
  }
  required_version = "0.13.5"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.40"
    }
  }
}

provider "aws" {
  region = "us-west-2"
}
