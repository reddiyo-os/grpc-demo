//Setup the global variables
module "global_variables" {
  source="../global/variables"
}

//Setup your provider to make it aware of your google setup - this probably just happened if you just ran the infrastructure
provider "google" {
    //This needs to be updated to wherever you put your credentials
  credentials = "${file("${module.global_variables.credentials}")}"
  project     = "${module.global_variables.project}"
  region      = "${module.global_variables.region}"
  provisioner "local-exec" {
    command = "gcloud container clusters get-credentials ${module.global_variables.clusterName} --zone ${module.global_variables.vpcLocation}"
  }
}

module "microservices" {
  source = "./modules"
}


