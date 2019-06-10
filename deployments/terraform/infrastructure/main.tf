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
    # We can't create a cluster with no node pool defined, but we want to only use
  # separately managed node pools. So we create the smallest possible default
  # node pool and immediately delete it.
  remove_default_node_pool = true
  initial_node_count = 1
  //Get your credentials for the newly created cluster so that microservices can be deployed
  provisioner "local-exec" {
    command = "gcloud container clusters get-credentials ${module.global_variables.clusterName} --zone ${module.global_variables.vpcLocation}"
  }
}

resource "google_container_node_pool" "primary_preemptible_nodes" {
  name       = "grpc-demo-node-pool"
  location   = "${module.global_variables.region}"
  cluster    = "${module.global_variables.clusterName}"
    
  autoscaling {
    # Minimum number of nodes in the NodePool. Must be >=0 and <= max_node_count.
    min_node_count = "3"

    # Maximum number of nodes in the NodePool. Must be >= min_node_count.
    max_node_count = "10"
  }

  node_config {
    preemptible  = true
    machine_type = "n1-standard-1"

    oauth_scopes = [
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/monitoring",
    ]
  }
}

