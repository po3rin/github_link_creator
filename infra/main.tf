provider "aws" {
  region = "ap-northeast-1"
}

terraform {
  required_version = ">= 0.12"

  backend "s3" {
    bucket = "po3rin-tfstate"
    region = "ap-northeast-1"
    key    = "ghcard/terraform.tfstate"
  }
}

module "iam" {
  source   = "./modules/iam"
  app_name = var.app_name
}

module "s3" {
  source   = "./modules/s3"
}
