module "global_variables" {
  source = "../global/variables"
}

//Setup your provider to make it aware of your google setup
provider "google" {
  //This needs to be updated to wherever you put your credentials
  credentials = "${file("${module.global_variables.credentials}")}"
  project     = "${module.global_variables.project}"
  region      = "${module.global_variables.region}"
}
//Setup your gke cluster
resource "google_container_cluster" "gke-cluster" {
  name               = "${module.global_variables.clusterName}"
  network            = "${module.global_variables.network}"
  location           = "${module.global_variables.vpcLocation}"

  initial_node_count = 5
  //Get your credentials for the newly created cluster so that microservices can be deployed
    provisioner "local-exec" {
    command = "gcloud config set project ${module.global_variables.project}"
  }
  provisioner "local-exec" {
    command = "gcloud container clusters get-credentials ${module.global_variables.clusterName} --zone ${module.global_variables.vpcLocation}"
  }
}


