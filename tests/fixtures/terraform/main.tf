terraform {
  required_version = ">= 1.9.8"
}

variable "environment" {
  type    = string
  default = "test"
}

output "environment" {
  value = var.environment
}
