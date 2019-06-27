
output "project" {
  value = "homeproject-231004"
  description = "The Google Project ID that you will be installing your cluster into"
}

output "region" {
  value = "us-central1"
  description = "The Region that will be used to install the cluster"
}

output "credentials" {
  value = "../../.././.creds/terraform/terraform-service-account.json"
  description = "The relative path location of the json file you will use to create this environment.   See exporting Service Account key.json from GCP if you need help."
}

output "clusterName" {
  value = "grpc-demo--backup-cluster"
  description = "the name of the cluster."
}

output "network" {
  value = "default"
  description = "the network that will be used."
}

output "vpcLocation" {
  value = "us-central1-b"
  description = "the specific zone that the VPC Will reside within"
}
