variable "project" {
  description = "The Google Project ID that you will be installing your cluster into"
  default = "homeproject-231004"
}

variable "region" {
  description = "The Region that will be used to install the cluster"
  default = "us-central1"
}

variable "credentials" {
  description = "The relative path location of the json file you will use to create this environment.   See exporting Service Account key.json from GCP if you need help."
  default = "../../.././.creds/terraform/terraform-service-account.json"
}

variable "clusterName" {
  default = "grpc-demo-cluster"
  description = "the name of the cluster"
}

variable "vpcLocation" {
  default = "us-central1-b"
  description = "the specific zone that the VPC Will reside within"
}
