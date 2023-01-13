variable "aws_region" {
  description = "AWS region for all resources."

  type    = string
  default = "us-west-2"
}

variable "environment" {
  description = "Environment the infrasructure is for."

  type    = string
  default = "prd"
}