provider "mongodbatlas" {
  public_key  = var.mongodb_atlas_api_pub_key
  private_key = var.mongodb_atlas_api_pri_key
}

variable "mongodb_atlas_api_pub_key" {}
variable "mongodb_atlas_api_pri_key" {}
variable "mongodb_atlas_org_id" {}
variable "mongodb_atlas_database_username" {}
variable "mongodb_atlas_database_user_password" {}
variable "mongodb_atlas_whitelistip_dev" {}
variable "mongodb_atlas_whitelistip_prod" {}

resource "mongodbatlas_project" "project" {
  name   = "test-project"
  org_id = var.mongodb_atlas_org_id
}

resource "mongodbatlas_cluster" "cluster" {
  project_id                  = mongodbatlas_project.project.id
  provider_name               = "GCP"
  name                        = "test-cluster"
  provider_instance_size_name = "M2"
  mongo_db_major_version      = "4.2"
  provider_region_name        = "WESTERN_EUROPE"
}

resource "mongodbatlas_database_user" "user" {
  auth_database_name = "admin"
  project_id         = mongodbatlas_project.project.id
  roles {
    role_name     = "atlasAdmin"
    database_name = "admin"
  }
  username = var.mongodb_atlas_database_username
  password = var.mongodb_atlas_database_user_password
}

resource "mongodbatlas_project_ip_whitelist" "dev" {
  project_id = mongodbatlas_project.project.id
  ip_address = var.mongodb_atlas_whitelistip_dev
  comment    = "dev"
}

resource "mongodbatlas_project_ip_whitelist" "prod" {
  project_id = mongodbatlas_project.project.id
  ip_address = var.mongodb_atlas_whitelistip_prod
  comment    = "prod"
}

output "connection_strings" {
  value = mongodbatlas_cluster.cluster.connection_strings
}
